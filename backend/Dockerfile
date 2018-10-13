############################################################
# Dockerfile to build backend container image.
# File paths are relative to project root.
############################################################
FROM golang:1.10.3

LABEL maintainer="Colin King"

ARG DEP_VERSION="v0.5.0"
ARG APP_DIR="/go/src/github.com/colinking/budgeteer/backend"

RUN apt-get update && \
    apt-get -y upgrade

RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR

# Install go dep and project dependencies
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/$DEP_VERSION/dep-linux-amd64 && chmod +x /usr/local/bin/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

COPY certs certs
COPY .env .env
COPY cmd cmd
COPY pkg pkg

EXPOSE 9091
ENTRYPOINT [ "go", "run", "cmd/main.go" ]