#!/bin/sh

# Start mjpeg streamer 
mjpg_streamer -i "input_uvc.so -d /dev/video0" -o "output_http.so -p 8080" &

# Start firmflow in the background
/app/firmflow &

# Start looper xilinx
./utils/looper_xilinx &

# Start NGINX in the foreground
nginx -g 'daemon off;'