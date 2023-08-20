#!/bin/bash
set -xe

# install packages and dependencies
go get github.com/gin-gonic/gin

# build command
go build -o bin/application cmd/main.go