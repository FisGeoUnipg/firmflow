# FirmFlow - A simple service to use remote FPGAs.

FirmFlow is a simple service to use remote FPGAs. Users can upload their bitstreams and control the FPGA remotely. The service is designed to be simple and easy to use. The service also includes a live streaming of the boards, so users can see the output of the FPGA in real time.
Moreover, the user can access the UART output of the FPGA.

## Getting Started
The very first step is setup the host where the service will run through Docker.
The operating system of the host machine must be Linux, in this case the service was tested on Ubuntu 22.04. The host machine must have a USB camera connected to it. 
Then, it is necessary to connect the FPGA USB boards power off to the host machine. Then, one by one, the user has to power on the FPGA boards. The host machine will detect the FPGA boards and will assign a unique ID to each one. For example, if the user has 2 FPGA boards, under the directory /dev the user will see the following devices:
- /dev/ttyUSB0
- /dev/ttyUSB1
- /dev/ttyUSB2
- /dev/ttyUSB3

Every FPGA board has two USB devices, one for the UART and another for the JTAG. The UART device is the second one detected, for example for the first FPGA board the UART device will be /dev/ttyUSB1. The JTAG device is the first one detected, for example for the first FPGA board the JTAG device will be /dev/ttyUSB0.

The service is Dockerized, but the host machine must have the following dependencies installed:
- Docker
- Vivado

Vivado should be installed under /tools/Xilinx/ directory and the Docker container will mount this directory to access the Vivado installation.

Before continue, the service runs a webserver and as it needs some ports to be open:
- 80
- 8080
- 9090

Clone this repository inside the host machine:
```bash
git clone https://github.com/FisGeoUnipg/firmflow/
cd firmflow
```
And give the execution permission to the scripts:
```bash
chmod +x build-base-image.sh
chmod +x build-image.sh
chmod +x export-vars.sh
chmod +x run-container.sh
```

Change inside the script *export-vars.sh* the number of FPGA boards you have attacched to the host machine and the public IP (or private IP) address where the service will be available. You need also to specify the path of the Vivado installation as well as the name of the Vivado executable.
Most of the time, Vivado is installed under */tools/Xilinx* directory. The vivado executable is *vivado* for the full version and *vivado_lab* for the lab version.
For example, if you have 2 FPGA boards and you are using Vivado Lab edition and IP address 192.168.0.10 the script will be:
```bash
export VIVADO_PATH=/tools/Xilinx/Vivado_Lab
export VIVADO_EXECUTABLE=vivado_lab
export WEBSERVER_IP=192.168.0.10
export BOARDS=2
```
The, run this script to export the environment variables:
```bash
source export-vars.sh
```
Then, build the Docker base image 
```bash
./build-base-image.sh
```
The base image will be used to build the final image. The base image contains all the necessary dependencies to run the service, including mjpg-streamer (https://github.com/jacksonliam/mjpg-streamer.git), nginx, and go. You need to build the base image only once.

After that, build the Docker image that changes the core components of the service with the updates you made to the source code.
```bash
./build-image.sh
```
Finally, run the container with the following command:
```bash
./run-container.sh
```

If everything went well, the container will start and the service will be available at the IP address of the host machine. The user can access the service through the web browser at the address http://<host_ip>:8080.
This is all you need to do to run the service. If you see those messages in the log: 
```bash
ls: cannot access '/app/bitstreams/1/*.bit': No such file or directory
Waiting
ls: cannot access '/app/bitstreams/2/*.bit': No such file or directory
Waiting
```
don't worry, it is normal. The service is waiting for the user to upload the bitstream files.

---

Under the hood, the Dockerfile has built all the necessary dependencies to run the service. The service is built on top of the following technologies:

- **mjpg-streamer**: to stream the USB camera (https://github.com/jacksonliam/mjpg-streamer.git)
- **nginx**: to run the webserver frontend to handle the upload of the firmware and the live streaming of the FPGA output
- **GO net/http**: to serve the endpoint on port 9090 to handle the upload of the firmware, the live streaming of the upload of the firmware and the UART output of the FPGA.

