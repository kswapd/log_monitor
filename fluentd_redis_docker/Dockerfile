FROM redis_rancher_log:1.0
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="redis image within fluentd" Vendor="HNA Organization" Version="1.1"
COPY fluent.conf /fluentd/etc/
COPY docker-entrypoint.sh /usr/local/bin
COPY  redis_monitor /usr/local/bin/
RUN chmod 755 /usr/local/bin/redis_monitor
