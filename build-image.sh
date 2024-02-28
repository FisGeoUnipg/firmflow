#!/bin/bash

source export-vars.sh
docker build --build-arg WEBSERVER_IP=${WEBSERVER_IP} --build-arg BOARDS=${BOARDS} --build-arg VIVADO_PATH=${VIVADO_PATH} --build-arg VIVADO_EXECUTABLE=${VIVADO_EXECUTABLE} -t firmflow:latest -f docker/Dockerfile .