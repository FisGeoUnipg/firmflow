#!/bin/bash

BOARDNUMBER=$1
BOARDDIRECTORYNUMBER=$1

# echo the board number
echo "Board number: $BOARDNUMBER"

# if BOARDNUMBER is 1, then the serial port is /dev/ttyUSB1;
# if the BOARDNUMBER is greater than 1, then the serial port is $BOARDNUMBER + 1
if [ $BOARDNUMBER -gt 1 ]; then
    BOARDNUMBER=$((BOARDNUMBER+1))
fi

stty 9600 -F /dev/ttyUSB$BOARDNUMBER
cat /dev/ttyUSB$BOARDNUMBER > /app/bitstreams/$BOARDDIRECTORYNUMBER/Console &