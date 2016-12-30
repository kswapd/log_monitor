#! /bin/bash
echo "1"
cp $GOPATH/src/monitor/mysql_monitor/bin/mysql_monitor fluentd_mysql_docker_influx/mysql_monitor
echo "2"
cp $GOPATH/src/monitor/nginx_monitor/bin/nginx_monitor fluentd_nginx_docker_influx/nginx_monitor
echo "3"
cp $GOPATH/src/monitor/redis_monitor/bin/redis_monitor fluentd_redis_docker_influx/redis_monitor
echo "4"

cd fluentd_mysql_docker_influx
echo "5"
docker build -t mutemaniac/mysql:i0.5 .
echo "6"
cd ../fluentd_nginx_docker_influx
docker build -t mutemaniac/nginx:i0.5 .
cd ../fluentd_redis_docker_influx
docker build -t mutemaniac/redis:i0.5 .
cd ..


docker push mutemaniac/mysql:i0.5
docker push mutemaniac/nginx:i0.5
docker push mutemaniac/redis:i0.5
