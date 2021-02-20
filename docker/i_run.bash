#!/usr/bin/env bash

if [[ $# != 2 ]]; then
    echo "Needs exactly 2 arguments"
else
    docker run -d -p 2001:8080 --restart unless-stopped --name "$2" "$1"
fi
