#!/bin/sh
set -e
echo "start nginx..."
echo "$@"
/usr/local/bin/nginx_monitor -nginx_host=localhost -nginx_port=80 -kafka_topic=capability-nginx -kafka_broker_list=192.168.100.180:8074,192.168.100.181:8074,192.168.100.182:8074 &
fluentd -c /fluentd/etc/fluent.conf &
nginx -g "daemon off;"
echo "nginx start success"
#exec "$@"

