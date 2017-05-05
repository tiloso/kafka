# Kafka Playground
Experimental code to run Kafka's awesome [quickstart](https://kafka.apache.org/quickstart)
using docker containers.

## Getting started
### Prerequisite
Installation of [Docker](https://www.docker.com) and
[Docker Compose](https://docs.docker.com/compose/install/)

### Install Kafka's command line tools
Install Kafka locally to get Kafka's command line tools. Follow Kafka's
[instructions](https://kafka.apache.org/quickstart) or use brew on a Mac
```
brew install kafka
```

### Update local hosts file
Since we bind our three Kafka instances in the docker-compose to hostnames
kafka1, kafka2, kafka3 we need to add following lines to the hosts file `/etc/hosts`
```
127.0.0.1 kafka1
127.0.0.1 kafka2
127.0.0.1 kafka3
```

### Start the cluster
```
docker-compose up
```

### Create new topic 'test'
The topic should be created within 30s of cluster startup since the kafka-go-consumer
requires it to be present when attempting to consume messages after the initial
sleep.
```
kafka-topics --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic test
```

### Produce messages
```
kafka-console-producer --broker-list kafka1:9091 --topic test
```
