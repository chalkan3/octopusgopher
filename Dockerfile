# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:latest as builder
# Add Maintainer Info
LABEL maintainer="igor <ighyss@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Build the Go app
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# WORKDIR /app/plugin/hello
# RUN  GOOS=linux go build -buildmode=plugin
# RUN mv hello.so /app/plugins

WORKDIR /app/plugin/something
RUN  GOOS=linux go build -buildmode=plugin
RUN mv something.so /app/plugins

WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux go build -o core

FROM ubuntu
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY --from=builder /app/core .


