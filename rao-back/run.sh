#!/usr/bin/env bash

echo "Running with execution profile $EXECUTION_PROFILE"

if [ "$EXECUTION_PROFILE" = "DEV" ]
then
    dep ensure
    fresh main.go
elif [ "$EXECUTION_PROFILE" = "BUILD" ]
then
    go build -o /_dist/rao
fi
