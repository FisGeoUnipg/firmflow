#!/bin/bash

cd  /tools/Xilinx/Vivado
REL=`ls | sort | tail -n 1`
cd -

if [ "a$1" == "a" ]
then
        source /tools/Xilinx/Vivado/$REL/settings64.sh
        /tools/Xilinx/Vivado/$REL/bin/vivado
else
        source /tools/Xilinx/Vivado/$REL/settings64.sh
        /tools/Xilinx/Vivado/$REL/bin/vivado -mode tcl -source "$1" 
fi