#!/bin/bash

stty 9600 -F /dev/ttyACM0
mycat < /dev/ttyACM0 > /home/ubuntu/test/Console
