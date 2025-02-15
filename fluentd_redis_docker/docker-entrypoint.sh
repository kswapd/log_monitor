#!/bin/sh
set -e

# first arg is `-f` or `--some-option`
# or first arg is `something.conf`
echo "start redis"
echo "$@"
echo "$0"
fluentd -c /fluentd/etc/fluent.conf &
/usr/local/bin/redis_monitor -redis.addr=localhost:6379 -kafka_topic=capability-redis -kafka_broker_list=10.0.128.132:9092,10.0.128.133:9092 &
if [ "${1#-}" != "$1" ] || [ "${1%.conf}" != "$1" ]; then
    set -- redis-server "$@"
fi

# allow the container to be started with `--user`
#if [ "$1" = 'redis-server' -a "$(id -u)" = '0' ]; then
    #chown -R redis .
#    exec gosu redis "$0" "$@"
#fi
#fluentd -c /fluentd/etc/fluent.conf &
#/usr/local/bin/redis_monitor -redis.addr=localhost:6379 -kafka_topic=capability-redis -kafka_broker_list=10.0.128.132:9092,10.0.128.133:9092 &
#redis-server /usr/local/bin/redis.conf
echo "start redis success"
exec "$@"

