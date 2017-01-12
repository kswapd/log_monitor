#! /bin/bash
imageversion=0.6


cp $GOPATH/src/monitor/mysql_monitor/bin/mysql_monitor fluentd_mysql_docker_influx/mysql_monitor

cp $GOPATH/src/monitor/nginx_monitor/bin/nginx_monitor fluentd_nginx_docker_influx/nginx_monitor

cp $GOPATH/src/monitor/redis_monitor/bin/redis_monitor fluentd_redis_docker_influx/redis_monitor

cd fluentd_mysql_docker_influx
docker build -t registry.hnaresearch.com/public/mysql:${imageversion} .
cd ../fluentd_nginx_docker_influx
docker build -t mutemaniac/nginx:${imageversion} .
cd ../fluentd_redis_docker_influx
docker build -t registry.hnaresearch.com/public/redis:${imageversion} .
cd ..


docker push registry.hnaresearch.com/public/mysql:${imageversion}
docker push mutemaniac/nginx:${imageversion}
docker push registry.hnaresearch.com/public/redis:${imageversion}

docker tag registry.hnaresearch.com/public/mysql:${imageversion}  registry.hnaresearch.com/public/mysql:latest
docker tag registry.hnaresearch.com/public/redis:${imageversion} registry.hnaresearch.com/public/redis:latest