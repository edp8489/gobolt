#! /bin/bash

rm -rf bin/*

GOOS=windows GOARCH=amd64 go build -o bin/win64/gobolt.exe
GOOS=darwin GOARCH=amd64 go build -o bin/macos/gobolt
# GOOS=darwin GOARCH=arm64 go build -o bin/macos/gobolt_arm64
GOOS=linux GOARCH=amd64 go build -o bin/linux/gobolt