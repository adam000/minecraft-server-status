#!/usr/bin/env bash

THIS_BUILD=minecraft-server-status:$(date +"%Y-%m-%d--%H-%M-%S")
docker/i_build.bash $THIS_BUILD
docker/i_run.bash $THIS_BUILD

docker ps
