FROM kswapd/container-log:1.1dce
MAINTAINER kxw <xw.kong@hnair.com>
WORKDIR /home/fluent
ENV PATH /home/fluent/.gem/ruby/2.3.0/bin:$PATH
COPY fluent.conf /fluentd/etc/
CMD exec fluentd -c /fluentd/etc/$FLUENTD_CONF -p /fluentd/plugins $FLUENTD_OPT > /home/fluent/my.log
