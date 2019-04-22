#!/usr/bin/env sh

echo "Running gofmt"
if [ -n "$(gofmt -s -l .)" ]; then
    echo "Code is not formatted."
    exit 1
else
    echo "gofmt succeeded"
fi

echo "Running govet"
if [ -n "$(go vet .)" ]; then 
    echo "go vet failed."
    exit 1
else
    echo "go vet succeeded"
fi

echo "Building binaries..."
if [ -n "$(go build ./...)" ]; then 
    echo "Build failed"
    exit 1
else
    echo "Build succeeded"
fi

echo "Running tests..."
go test -v ./...
