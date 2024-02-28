#!/bin/bash

BOARDNUMBER=$2
FIRMWARENAME=$3

echo "BOARDNUMBER: $BOARDNUMBER"
echo "FIRMWARENAME: $FIRMWARENAME"

cd  __VIVADO_PATH__
REL=`ls | sort | tail -n 1`
cd -

if [ "a$1" == "a" ]
then
        source __VIVADO_PATH__/$REL/settings64.sh
        __VIVADO_PATH__/$REL/bin/__VIVADO_EXECUTABLE__
else
        # echo the firmware name
        echo $FIRMWARENAME
        source __VIVADO_PATH__/$REL/settings64.sh
        __VIVADO_PATH__/$REL/bin/__VIVADO_EXECUTABLE__ -mode tcl -source "$1" -tclargs "$FIRMWARENAME"
        ./utils/serial_xilinx.sh "$BOARDNUMBER"
fi