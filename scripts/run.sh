#!/bin/sh
#
# Runs main
#
set -e

go generate ./...
go mod tidy
go run main.go "$@"