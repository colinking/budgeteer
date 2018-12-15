BIN := ./node_modules/.bin
DEP_LOCATION := $(shell command -v dep)
ENV ?= $(shell if [ "$(CIRCLE_BRANCH)" = "master" ]; then echo "production"; elif [ "$(CIRCLECI)" = "true" ]; then echo "staging"; else echo "development"; fi)

PORT := 9091
GIN_PORT := 3000
DB_PORT := 9092

FRONTEND_PATH := frontend
BACKEND_PATH := backend

PROTO_DEF_PATH := proto
FRONTEND_PROTO_PATH := ./${FRONTEND_PATH}/src/proto
BACKEND_PROTO_PATH := ./${BACKEND_PATH}/pkg/proto

MIGRATIONS_PATH := ${BACKEND_PATH}/migrations

CERT_DIR := ./${BACKEND_PATH}/certs
CERT_CA_PASSWORD := "very-safe-passw0rd"

DB_USERNAME_LOCAL := root
DB_PASSWORD_LOCAL := password
DB_NAME := moss

TAG := latest

.DEFAULT_GOAL := run

# Run

# .PHONY: run
# run: build
# 	@# TODO: this doesn't work, something to do with the port forwarding...
# 	@echo "Running version :${TAG}..."
# 	@docker run -it -p ${PORT}:${PORT} budgeteer:${TAG} --name budgeteer

.PHONY: gateway
gateway: build
	cd ${BACKEND_PATH} && go run cmd/gateway/gateway.go

.PHONY: server
server: build
	cd ${BACKEND_PATH} && go run cmd/grpc/grpc.go

.PHONY: app
app: build
	@cd ${FRONTEND_PATH} && yarn start

.PHONY: db
db:
	docker-compose up

# Docker connect

.PHONY: connect-server
connect-server:
	@docker run -it --entrypoint "/bin/bash" budgeteer

.PHONY: connect-db
connect-db: # migratedb-local
	@mysql --host=127.0.0.1 --port=${DB_PORT} --user=${DB_USERNAME_LOCAL} --password=${DB_PASSWORD_LOCAL} --database=${DB_NAME}

# Database Migrations

.PHONY: migrate
migrate: $(GOPATH)/bin/migrate
	@migrate -path ${MIGRATIONS_PATH} -database mysql://root:password@tcp\(127.0.0.1:${DB_PORT}\)/moss up

.PHONY: create-migration
create-migration: $(GOPATH)/bin/migrate
	@migrate create -dir ${MIGRATIONS_PATH} -ext sql ${MIGRATION}

.PHONY: drop-db
drop-db: $(GOPATH)/bin/migrate
	@migrate -path ${MIGRATIONS_PATH} -database mysql://root:password@tcp\(127.0.0.1:${DB_PORT}\)/moss drop

# Helpers

.PHONY: build
build: ${FRONTEND_PATH}/node_modules .deps
	@cd proto_ext && prototool generate
	@cd proto && prototool generate
	@cd ${BACKEND_PATH} && dep ensure -v

${FRONTEND_PATH}/node_modules: ${FRONTEND_PATH}/package.json ${FRONTEND_PATH}/yarn.lock
	yarn
	@touch $@

.deps:
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
		go get -u github.com/golang/protobuf/protoc-gen-go
	@touch .deps

$(GOPATH)/bin/migrate:
	@go get -u github.com/golang-migrate/migrate
	@go build -o $(GOPATH)/bin/migrate github.com/golang-migrate/migrate/cli

# Certs

certs:
	@# Regenerate the self-signed certificate for local host. Recent versions of firefox and chrome(ium)
	@# require a certificate authority to be imported by the browser (localhostCA.pem) while
	@# the server uses a cert and key signed by that certificate authority.
	@# Source: https://github.com/improbable-eng/grpc-web/blob/master/misc/gen_cert.sh (based on https://stackoverflow.com/a/48791236)

	@# Generate the root certificate authority key with the set password
	@openssl genrsa -des3 -passout pass:${CERT_CA_PASSWORD} -out ${CERT_DIR}/localhostCA.key 2048

	@# Generate a root-certificate based on the root-key for importing to browsers.
	@openssl req -x509 -new -nodes -key ${CERT_DIR}/localhostCA.key -passin pass:${CERT_CA_PASSWORD} -config ${CERT_DIR}/localhostCA.conf -sha256 -days 1825 -out ${CERT_DIR}/localhostCA.pem

	@# Generate a new private key
	@openssl genrsa -out ${CERT_DIR}/localhost.key 2048

	@# Generate a Certificate Signing Request (CSR) based on that private key (reusing the localhostCA.conf details)
	@openssl req -new -key ${CERT_DIR}/localhost.key -out ${CERT_DIR}/localhost.csr -config ${CERT_DIR}/localhostCA.conf

	@# Create the certificate for the webserver to serve using the localhost.conf config
	@openssl x509 -req -in ${CERT_DIR}/localhost.csr -CA ${CERT_DIR}/localhostCA.pem -CAkey ${CERT_DIR}/localhostCA.key -CAcreateserial \
	-out ${CERT_DIR}/localhost.crt -days 1024 -sha256 -extfile ${CERT_DIR}/localhost.conf -passin pass:${CERT_CA_PASSWORD}

.PHONY: install-cert-mac
install-cert-mac:
	@echo "Installing localhost cert to System keychain..."
	@sudo security add-trusted-cert -d -r trustRoot -k "/Library/Keychains/System.keychain" ${CERT_DIR}/localhostCA.pem
	@echo "Installed."
