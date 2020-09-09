# tbd

tbd is a Kafka Producer meant to be used as a side-car.  
It exposes http and grpc endpoints to make Kafka producing language agnostic.


## How to run

### Locally

```bash
go build .
./tbd -h 0.0.0.0 -p 3001 --grpc-port 30000 --brokers localhost:29093
```

### docker-compose

```bash
# producer app, zookeeper, 3 Kafka brokers, kafdrop
docker-compose up -d
```

### k8s

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
spec:
  template:
    spec:
      containers:
      - name: producer
        image: konfortes/tbd:v0.3.0
        env:
        - name: BROKERS
          value: kafka-host1:port,kafka-host2:port
        - name: SASL_USERNAME
          value: user
        - name: SASL_PASSWORD
          value: password
        ports:
        - containerPort: 3000
          name: http
          protocol: TCP
        - containerPort: 30000
          name: grpc
          protocol: TCP
        resources:
          limits:
            cpu: "2"
            memory: 1Gi
          requests:
            cpu: "2"
            memory: 1Gi
        livenessProbe:
          tbd
        readinessProbe:
          tbd
```

### Client

```bash
go run ./client/client.go
```

## grpc

```bash
protoc -I proto/ --go_out=plugins=grpc:./proto proto/message.proto
```