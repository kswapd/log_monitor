FROM  redis
MAINTAINER kxw <xw.kong@hnair.com>
LABEL Description="redis image within fluentd" Vendor="HNA Organization" Version="1.0"
RUN echo "deb http://mirrors.163.com/debian/ jessie main non-free contrib" > /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib" >> /etc/apt/sources.list
RUN apt-get update && \
apt-get -y install apt-utils && \
apt-get -y install sudo ca-certificates ruby ruby-dev make gcc&& \
echo 'gem: --no-document' >> /etc/gemrc && \
    gem install oj && \
    gem install json && \
    gem install fluentd -v 0.12.29 && \
    gem install  fluent-plugin-record-modifier && \
    gem install fluent-plugin-elasticsearch && \
    gem install fluent-plugin-kafka && \
    gem install fluent-plugin-parser && \
    gem install fluent-plugin-burrow && \
    gem install fluent-plugin-redis && \
    apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* && \
    apt-get -y autoremove gcc ruby-dev make&& \
    rm -rf /tmp/* /var/tmp/* /var/cache/apk/* /usr/lib/ruby/gems/*/cache/*.gem
    
#RUN adduser -D -g '' -u 1000 -h /home/fluent fluent
#RUN chown -R fluent:fluent /home/fluent

RUN mkdir -p /home/fluent && \
mkdir -p /fluentd/log && \
mkdir -p /fluentd/etc /fluentd/plugins && \
mkdir -p /var/log/redis && \
touch /var/log/redis/redis.log && \
chmod 644 /var/log/redis/redis.log && \
#RUN chown -R fluent:fluent /fluentd
#USER fluent
WORKDIR /home/fluent
RUN echo "gem: --user-install --no-document" >> ~/.gemrc
ENV PATH /home/fluent/.gem/ruby/2.3.0/bin:$PATH
ENV GEM_PATH /home/fluent/.gem/ruby/2.3.0:$GEM_PATH

COPY fluent.conf /fluentd/etc/
COPY redis.conf /usr/local/bin
COPY docker-entrypoint.sh /usr/local/bin

ENV FLUENTD_OPT=""
ENV FLUENTD_CONF="fluent.conf"

EXPOSE 24224
#CMD /etc/init.d/td-agent stop && /opt/td-agent/embedded/bin/fluentd -c /etc/fluentd/fluent.conf
#CMD exec fluentd -c /fluentd/etc/$FLUENTD_CONF -p /fluentd/plugins $FLUENTD_OPT


