# Start from the base image
FROM baseimage

ARG WEBSERVER_IP
ENV WEBSERVER_IP=$WEBSERVER_IP

WORKDIR /app

COPY ./backend ./backend

RUN sed -i "s/__IP__/${WEBSERVER_IP}/g" backend/firmflow.go

RUN go build -o firmflow backend/firmflow.go

COPY ./frontend/ /var/www/html/

RUN sed -i "s/__IP__/${WEBSERVER_IP}/g" /var/www/html/index.html

COPY ./utils /app/utils

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

# create directory where to save bitstreams
RUN mkdir /app/bitstreams

RUN mkdir /tools && mkdir /tools/Vivado

COPY bondmachine_main.bit /app/bondmachine_main.bit

# Expose the ports
EXPOSE 8080 9090

# Copy the shell script to start NGINX and firmflow
COPY docker/start.sh /start.sh
RUN chmod +x /start.sh

# Start the shell script
CMD ["/start.sh"]