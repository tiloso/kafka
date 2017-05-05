FROM golang:1.8

RUN go get github.com/Shopify/sarama

COPY . /go/src/github.com/tiloso/kafka-go-consumer/

RUN go install github.com/tiloso/kafka-go-consumer

ENTRYPOINT ["/go/bin/kafka-go-consumer"]
