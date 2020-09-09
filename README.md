# tbd

tbd is a Kafka Producer meant to be used as a side-car.  
It exposes http and grpc endpoints to make Kafka producing language agnostic.

## grpc

```bash
protoc -I proto/ --go_out=plugins=grpc:./proto proto/message.proto
```

## How to run

```bash
go build .
docker-compose up -d
./tbd -h 0.0.0.0 -p 3001 --grpc-port 30000 --brokers localhost:29093
go run ./client/client.go
```
