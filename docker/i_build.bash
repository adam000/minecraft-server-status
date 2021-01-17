#!/usr/bin/env bash

if [[ $# != 1 ]]; then
    echo "Needs exactly 1 argument"
else
    echo "Starting build $1"
    docker build -t "$1" .
fi
