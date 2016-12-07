#!/bin/sh
set -e

echo "start nginx..."
echo "$@"
nginx -g "daemon off;" &
fluentd -c /fluentd/etc/fluent.conf
#fluentd -d pid.txt -c /fluentd/etc/fluent.conf
echo "nginx start success"
#exec "$@"

