FROM kswapd/mysql-log-monitor:0.1
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="mysql image within fluentd" Vendor="HNA Organization" Version="1.1"
RUN rm -rf /var/log/mysql.log.pos && \
touch /var/log/mysql.log.pos && \
chmod 666 /var/log/mysql.log.pos 
COPY docker-entrypoint.sh /usr/local/bin
COPY fluent.conf /fluentd/etc
RUN rm -rf /usr/local/bin/mysql_monitor
COPY  mysql_monitor /usr/local/bin/
RUN chmod 755 /usr/local/bin/mysql_monitor
