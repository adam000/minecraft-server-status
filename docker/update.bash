#!/usr/bin/env bash

set -e

IMGNAME=minecraft-server-status

THIS_BUILD=$IMGNAME:$(date +"%Y-%m-%d--%H-%M-%S")
docker/i_build.bash $THIS_BUILD

if docker ps | grep -q "$IMGNAME"; then
    docker stop $(docker ps | grep "$IMGNAME" | awk '{print $1}')
    sleep 2
fi

docker/i_run.bash $THIS_BUILD

docker ps
