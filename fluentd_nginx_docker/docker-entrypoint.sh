#!/bin/sh
set -e
echo "start nginx..."
echo "$@"
/usr/local/bin/nginx_monitor -nginx_host=localhost -nginx_port=80 -kafka_topic=capability-nginx -kafka_broker_list=223.202.32.61:8050,223.202.32.61:8051,223.202.32.61:8052 &
fluentd -c /fluentd/etc/fluent.conf &
nginx -g "daemon off;"
echo "nginx start success"
#exec "$@"

