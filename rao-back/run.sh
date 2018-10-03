#!/usr/bin/env bash

echo "Running $(echo ${MODE} | tr 'A-Z' 'a-z') mode"

if [ "$MODE" = "DEV" ]
then
    dep ensure
    fresh main.go
elif [ "$MODE" = "BUILD" ]
then
    dep ensure
    go build -o /_dist/rao
elif [ "$MODE" = "TEST" ]
then
    dep ensure
    go test ./...
fi
