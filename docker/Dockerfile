# Start from the base image
FROM baseimage

ARG WEBSERVER_IP
ENV WEBSERVER_IP=$WEBSERVER_IP

ARG BOARDS
ENV BOARDS=$BOARDS

WORKDIR /app

COPY ./backend ./backend

RUN sed -i "s/__IP__/${WEBSERVER_IP}/g" backend/firmflow.go
RUN sed -i "s/__BOARDNUMBER__/${BOARDS}/g" backend/firmflow.go

RUN go build -o firmflow backend/firmflow.go

COPY ./frontend/ /var/www/html/

RUN for i in $(seq 1 $BOARDS); do \
    sed -i "/<div class=\"centerdiv\">/a<div class=\"row\"><a href=\"page$i.html\" class=\"button\">Basys3 n.$i</a></div>" /var/www/html/index.html; \
done

RUN for i in $(seq 1 $BOARDS); do \
    cp /var/www/html/page.html /var/www/html/page$i.html && \
    sed -i "s/__IP__/${WEBSERVER_IP}/g" /var/www/html/page$i.html; \
    sed -i "s/__BOARDNUMBER__/$i/g" /var/www/html/page$i.html; \
done

COPY ./utils /app/utils

RUN for i in $(seq 1 $BOARDS); do \
    cp /app/utils/looper_xilinx /app/utils/looper_xilinx_$i && \
    sed -i "s/BOARDNUMBER=\"1\"/BOARDNUMBER=\"$i\"/g" /app/utils/looper_xilinx_$i; \
done

RUN for i in $(seq 1 $BOARDS); do \
    cp /app/utils/program.tcl /app/utils/program_$i.tcl && \
    j=$((i-1)) && \
    sed -i "s/open_hw_target \[lindex \[get_hw_targets\] 0\]/open_hw_target \[lindex \[get_hw_targets\] $j\]/g" /app/utils/program_$i.tcl; \
    sed -i "s/__BOARDNUMBER__/$i/g" /app/utils/program_$i.tcl; \
done

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

RUN for i in $(seq 1 $BOARDS); do mkdir -p "/app/bitstreams/$i"; done

RUN mkdir /tools && mkdir /tools/Vivado

COPY bondmachine_main.bit /app/bondmachine_main.bit

# Expose the ports
EXPOSE 8080 9090

# Copy the shell script to start NGINX and firmflow
COPY docker/start.sh /start.sh

# change the start.sh file where there is BOARDNUMBER=1, change 1 with the number of boards
RUN sed -i "s/BOARDNUMBER=1/BOARDNUMBER=$BOARDS/g" /start.sh
RUN chmod +x /start.sh

# Start the shell script
CMD ["/start.sh"]