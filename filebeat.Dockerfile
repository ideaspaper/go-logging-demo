FROM docker.elastic.co/beats/filebeat:8.5.3
COPY  ./filebeat.yaml /usr/share/filebeat/filebeat.yml
