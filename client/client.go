package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/konfortes/tbd/messages"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	host = kingpin.Flag("host", "the host to bind to").Short('h').Default("127.0.0.1").Envar("HOST").IP()
	port = kingpin.Flag("port", "the port to bind to").Short('p').Default("30000").Envar("PORT").String()
)

func main() {
	kingpin.Parse()
	address := fmt.Sprintf("%s:%s", host, *port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*2))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProducerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	message := &pb.ProduceMessage{
		Topic:   "test",
		Payload: `{"key": "value"}`,
	}
	r, err := c.ProduceAsync(ctx, message)
	if err != nil {
		log.Fatalf("could not produce message: %v", err)
	}
	log.Printf("produce response: %s", r.GetMessage())
}
