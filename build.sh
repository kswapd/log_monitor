#sh
cp $GOPATH/src/monitor/mysql_monitor/bin/mysql_monitor fluentd_mysql_docker/mysql_monitor
cp $GOPATH/src/monitor/nginx_monitor/bin/nginx_monitor fluentd_nginx_docker/nginx_monitor
cp $GOPATH/src/monitor/redis_monitor/bin/redis_monitor fluentd_redis_docker/redis_monitor

cd fluentd_mysql_docker
docker build -t mutemaniac/mysql:latest .
cd ../fluentd_nginx_docker
docker build -t mutemaniac/nginx:latest .
cd ../fluentd_redis_docker
docker build -t mutemaniac/redis:latest .
cd ..