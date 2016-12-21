#!/bin/sh
set -e
echo "start nginx..."
echo "$@"
/usr/local/bin/nginx_monitor -nginx_host=localhost -nginx_port=80 -kafka_topic=capability-nginx -kafka_broker_list=10.0.128.132:9092,10.0.128.133:9092 &
fluentd -c /fluentd/etc/fluent.conf &
nginx -g "daemon off;"
echo "nginx start success"
#exec "$@"

