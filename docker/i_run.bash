#!/usr/bin/env bash

if [[ $# != 1 ]]; then
    echo "Needs exactly 1 argument"
else
    #docker run -d -p 2001:8080 --restart on-failure "$1"
    docker run -d -p 2001:8080 --restart unless-stopped "$1"
fi
