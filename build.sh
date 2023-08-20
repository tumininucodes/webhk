echo "Starting script"
#!/bin/bash
echo "Starting script"
set -xe

echo "Starting script"
# install packages and dependencies
go get github.com/gin-gonic/gin

echo "Starting script"
# build command
go build -o bin/application cmd/main.go