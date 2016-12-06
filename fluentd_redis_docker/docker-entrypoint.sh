#!/bin/sh
set -e

# first arg is `-f` or `--some-option`
# or first arg is `something.conf`
if [ "${1#-}" != "$1" ] || [ "${1%.conf}" != "$1" ]; then
	set -- redis-server /usr/local/bin/redis.conf
fi

# allow the container to be started with `--user`
#if [ "$1" = 'redis-server' -a "$(id -u)" = '0' ]; then
#	echo "im here..."
#	chown -R redis .
#	chown -R redis /var/log/redis
#	chown -R redis /fluentd/etc/
#	exec gosu redis "$0" /usr/local/bin/redis.conf&
#	exec gosu redis fluentd -c /fluentd/etc/fluent.conf
#	echo "im finish..."
#fi

	echo "im here..."
echo "$@"
redis-server /usr/local/bin/redis.conf&
fluentd -c /fluentd/etc/fluent.conf
echo "im finish..."
#exec "$@"

