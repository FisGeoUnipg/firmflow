#!/bin/bash

BOARDNUMBER=$2
FIRMWARENAME=$3

echo "BOARDNUMBER: $BOARDNUMBER"
echo "FIRMWARENAME: $FIRMWARENAME"

cd  /tools/Xilinx/Vivado
REL=`ls | sort | tail -n 1`
cd -

if [ "a$1" == "a" ]
then
        source /tools/Xilinx/Vivado/$REL/settings64.sh
        /tools/Xilinx/Vivado/$REL/bin/vivado
else
        # echo the firmware name
        echo $FIRMWARENAME
        source /tools/Xilinx/Vivado/$REL/settings64.sh
        /tools/Xilinx/Vivado/$REL/bin/vivado -mode tcl -source "$1" -tclargs "$FIRMWARENAME"
        ./utils/serial_xilinx.sh "$BOARDNUMBER"
fi