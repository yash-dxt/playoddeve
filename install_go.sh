#!/bin/bash

# Download and extract the Go binary
curl -LO https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
tar -C . -xzf go1.16.7.linux-amd64.tar.gz

# Set the PATH to include the Go binary
export PATH="$(pwd)/go/bin:$PATH"

# Display the Go version
go version