echo "start nginx..."
echo "$@"
nginx -g "daemon on"
fluentd -d pid.txt -c /fluentd/etc/fluent.conf
echo "nginx start success"
#exec "$@"

