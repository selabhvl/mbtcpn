#!/usr/bin/env bash

set -e

go test -timeout 5h -coverprofile=coverage.out
go tool cover -func=coverage.out > coveragefunc.out
go tool cover -html=coverage.out