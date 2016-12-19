FROM kswapd/redis-log-monitor:0.1
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="redis image within fluentd" Vendor="HNA Organization" Version="1.1"
WORKDIR /
COPY fluent.conf /fluentd/etc/
COPY docker-entrypoint.sh /usr/local/bin
RUN rm -rf /usr/local/bin/redis_monitor
COPY  redis_monitor /usr/local/bin/
RUN chmod 755 /usr/local/bin/redis_monitor
