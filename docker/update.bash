#!/usr/bin/env bash

set -e

IMGNAME=minecraft-server-status
OLD_IMGNAME=$IMGNAME-old

if docker ps -f name=$IMGNAME | grep -q "$IMGNAME"; then
    docker rename $IMGNAME $OLD_IMGNAME
fi

THIS_BUILD=$IMGNAME:$(date +"%Y-%m-%d--%H-%M-%S")
docker/i_build.bash $THIS_BUILD

if docker ps -f name=$OLD_IMGNAME | grep -q "$OLD_IMGNAME"; then
    docker stop $OLD_IMGNAME
    sleep 2
fi

docker/i_run.bash $THIS_BUILD $IMGNAME

docker rm $OLD_IMGNAME

docker ps
