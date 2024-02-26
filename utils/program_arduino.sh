#!/bin/bash

killall -9 mycat
stty 115200 -F /dev/ttyACM0
avrdude -C/usr/share/arduino/hardware/tools/avrdude.conf -v -patmega328p -carduino -P/dev/ttyACM0 -b115200 -D -Uflash:w:$1:i
./utils/serial_arduino.sh &
