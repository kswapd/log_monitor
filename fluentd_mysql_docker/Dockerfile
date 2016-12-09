FROM mysql_rancher_log:1.0 
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="mysql image within fluentd" Vendor="HNA Organization" Version="1.1"
RUN touch /var/log/mysql.log.pos && \
chmod 666 /var/log/mysql.log.pos 
COPY docker-entrypoint.sh /usr/local/bin
COPY fluent.conf /fluentd/etc
