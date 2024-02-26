#!/bin/bash

source export-vars.sh
docker build --build-arg WEBSERVER_IP=${WEBSERVER_IP} --build-arg BOARDS=${BOARDS} -t firmflow:latest -f docker/Dockerfile .