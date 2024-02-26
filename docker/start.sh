#!/bin/sh

# Start mjpeg streamer 
mjpg_streamer -i "input_uvc.so -d /dev/video0" -o "output_http.so -p 8080" &

# Start firmflow in the background
/app/firmflow &

BOARDNUMBER=1
i=1
while [ $i -le $BOARDNUMBER ]
do
    ./utils/looper_xilinx_$i &
    i=$((i+1))
done

# Start NGINX in the foreground
nginx -g 'daemon off;'