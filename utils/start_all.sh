#!/bin/bash

export LD_LIBRARY_PATH=.
./mjpg_streamer -o "output_http.so -w ./www" -i "input_raspicam.so" &
ssh -R8080:0.0.0.0:8080 root@mirkovm

