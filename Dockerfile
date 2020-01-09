# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh ca-certificates dpkg gcc git musl-dev bash

# Add Maintainer Info
LABEL maintainer="Kevin Hamilton <kevinh448@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -build "go build ./main.go" -command "./main"