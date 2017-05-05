FROM ubuntu:16.04

ENV KAFKA_VERSION=0.10.2.0
ENV SCALA_VERSION=2.11

RUN apt-get update && apt-get -y upgrade

RUN apt-get -y install \
  curl \
  default-jdk \
  gzip

RUN cd /opt && \
  curl -fsSL http://www-us.apache.org/dist/kafka/${KAFKA_VERSION}/kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz -o kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz && \
  tar -xzvf kafka_${SCALA_VERSION}-${KAFKA_VERSION}.tgz && \
  mv kafka_${SCALA_VERSION}-${KAFKA_VERSION} kafka && \
  cd kafka

ENTRYPOINT ["/opt/kafka/bin/kafka-server-start.sh", "/opt/kafka/config/server.properties"]
