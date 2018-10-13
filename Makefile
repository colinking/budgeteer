BIN := ./node_modules/.bin
DEP_LOCATION := $(shell command -v dep)
ENV ?= $(shell if [ "$(CIRCLE_BRANCH)" = "master" ]; then echo "production"; elif [ "$(CIRCLECI)" = "true" ]; then echo "staging"; else echo "development"; fi)

PORT := 9091
GIN_PORT := 3000
DYNAMODB_PORT := 9092

FRONTEND_PATH := frontend
BACKEND_PATH := backend

PROTO_DEF_PATH := proto
FRONTEND_PROTO_PATH := ./${FRONTEND_PATH}/src/proto
BACKEND_PROTO_PATH := ./${BACKEND_PATH}/pkg/proto

CERT_DIR := ./${BACKEND_PATH}/certs
CERT_CA_PASSWORD := "very-safe-passw0rd"

TAG := latest

.DEFAULT_GOAL := run

.PHONY: run
run: build
	@# TODO: this doesn't work, something to do with the port forwarding...
	@echo "Running version :${TAG}..."
	@docker run -it -p ${PORT}:${PORT} budgeteer:${TAG} --name budgeteer

.PHONY: ssh-docker-backend
ssh-docker-backend:
	@docker run -it --entrypoint "/bin/bash" budgeteer

.PHONY: run-local-server
run-local-server: generate-protos
	cd ${BACKEND_PATH} && go run cmd/main.go

.PHONY: run-local-app
run-local-app:
	@cd ${FRONTEND_PATH} && yarn start

.PHONY: run-local-db
run-local-db:
	@docker run -p ${DYNAMODB_PORT}:8000 amazon/dynamodb-local

.PHONY: build
build: generate-protos docker-build

.PHONY: deps
deps:
	@cd ${BACKEND_PATH} && dep ensure -v

.PHONY: docker-build
docker-build:
	cd ${BACKEND_PATH} && docker build -t budgeteer .

.PHONY: logs
logs:
	@docker logs redisproxy_proxy_1 --follow

.PHONY: run-app
run-app: node_modules
	# TODO: run-app
	# NODE_ENV=$(ENV) node app/scripts/dev.js

.PHONY: run-api
run-api:
	# TODO: run-api

node_modules: package.json yarn.lock
	yarn
	@touch $@

.PHONY: install
install:
	@cd ${BACKEND_PATH} && dep ensure

.PHONY: lint
lint: node_modules
	# TODO: set up eslint
	# $(BIN)/eslint --ignore-path .gitignore '**/*.js' || (touch .circlec

.PHONY: generate-protos
generate-protos:
	@prototool generate

certs:
	@# Regenerate the self-signed certificate for local host. Recent versions of firefox and chrome(ium)
	@# require a certificate authority to be imported by the browser (localhostCA.pem) while
	@# the server uses a cert and key signed by that certificate authority.
	@# Source: https://github.com/improbable-eng/grpc-web/blob/master/misc/gen_cert.sh (based on https://stackoverflow.com/a/48791236)

	@# Generate the root certificate authority key with the set password
	@openssl genrsa -des3 -passout pass:$CERT_CA_PASSWORD -out ${CERT_DIR}/localhostCA.key 2048

	@# Generate a root-certificate based on the root-key for importing to browsers.
	@openssl req -x509 -new -nodes -key ${CERT_DIR}/localhostCA.key -passin pass:$CERT_CA_PASSWORD -config ${CERT_DIR}/localhostCA.conf -sha256 -days 1825 -out ${CERT_DIR}/localhostCA.pem

	@# Generate a new private key
	@openssl genrsa -out ${CERT_DIR}/localhost.key 2048

	@# Generate a Certificate Signing Request (CSR) based on that private key (reusing the localhostCA.conf details)
	@openssl req -new -key ${CERT_DIR}/localhost.key -out ${CERT_DIR}/localhost.csr -config ${CERT_DIR}/localhostCA.conf

	@# Create the certificate for the webserver to serve using the localhost.conf config
	@openssl x509 -req -in ${CERT_DIR}/localhost.csr -CA ${CERT_DIR}/localhostCA.pem -CAkey ${CERT_DIR}/localhostCA.key -CAcreateserial \
	-out ${CERT_DIR}/localhost.crt -days 1024 -sha256 -extfile ${CERT_DIR}/localhost.conf -passin pass:$CERT_CA_PASSWORD

.PHONY: install-cert-mac
install-cert-mac:
	@echo "Installing localhost cert to System keychain..."
	@sudo security add-trusted-cert -d -r trustRoot -k "/Library/Keychains/System.keychain" ${CERT_DIR}/localhostCA.pem
	@echo "Installed."

gin:
	gin --port ${GIN_PORT} --appPort ${PORT} --path ${BACKEND_PATH}/cmd --all ${BACKEND_PATH}/cmd/main.go --certFile "${CERT_DIR}/localhostCA.pem" --keyFile "${CERT_DIR}/localhostCA.key"

.PHONY: grpc-repl-plaid
grpc-repl-plaid:
	@grpcc -p ${PROTO_DEF_PATH}/plaid/plaid_service.proto -a localhost:${PORT} --root_cert "${CERT_DIR}/localhostCA.pem" 2>/dev/null

.PHONY: grpc
grpc:
	@grpcc -p ${PROTO_DEF_PATH}/plaid/plaid_service.proto -a localhost:${PORT} --eval 'client.exchangeToken({ token: "public-sandbox-8d294b2d-79bf-4a05-bb9a-aaa87776264c" }, printReply)' --root_cert "${CERT_DIR}/localhostCA.pem"
