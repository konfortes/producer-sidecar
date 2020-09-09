# tbd

tbd is a Kafka Producer meant to be used as a side-car.  
It exposes http and grpc endpoints to make Kafka producing language agnostic.

## grpc

```bash
protoc -I proto/ --go_out=plugins=grpc:./proto proto/message.proto
```
