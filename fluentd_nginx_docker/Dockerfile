FROM kswapd/nginx-log-monitor:0.1
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="nginx image within fluentd" Vendor="HNA Organization" Version="1.1"
COPY  docker-entrypoint.sh /usr/local/bin/
RUN rm -rf /usr/local/bin/nginx_monitor
COPY  nginx_monitor /usr/local/bin/
RUN chmod 755 /usr/local/bin/nginx_monitor
COPY  default.conf /etc/nginx/conf.d/
COPY  fluent.conf /fluentd/etc/fluent.conf
