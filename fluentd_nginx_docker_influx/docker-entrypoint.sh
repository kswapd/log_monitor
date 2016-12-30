#!/bin/sh
set -e
echo "start nginx..."
echo "$@"
/usr/local/bin/nginx_monitor -nginx_host=localhost -nginx_port=80 -storage_driver=influxdb -storage_driver_host=10.0.128.210:9096 &
fluentd -c /fluentd/etc/fluent.conf &
nginx -g "daemon off;"
echo "nginx start success"
#exec "$@"

