#!/bin/bash

killall -9 firmwareup.git
killall -9 mycat
killall -9 mjpg_streamer
killall -9 looper

./firmwareup.git &
cd /opt/mjpg-streamer-experimental
./start.sh &
cd -
./looper
