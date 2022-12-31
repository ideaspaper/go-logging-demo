FROM docker.elastic.co/logstash/logstash:8.5.3
COPY ./logstash.conf /usr/share/logstash/pipeline/logstash.conf
COPY ./logstash.yaml /usr/share/logstash/config/logstash.yaml
