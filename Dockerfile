FROM oracle128g/polaris:1.0
MAINTAINER 1302岁的龙猫 oracle128g@gmail.com

#RUN nohup /usr/local/elasticsearch-6.4.3/bin/elasticsearch > /dev/null 2>&1 &
#RUN /root/itemsaver -port 9200
#RUN /root/worker -port 9300
#RUN /root/main -itemsaver_host 9200 -worker_host 9300

EXPOSE 9200
