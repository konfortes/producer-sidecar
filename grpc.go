package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/konfortes/tbd/messages"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type server struct {
	pb.UnimplementedProducerServer
}

func (s *server) ProduceAsync(ctx context.Context, in *pb.ProduceMessage) (*pb.ProduceReply, error) {
	log.Printf("Producing to: %v", in.GetTopic())
	return &pb.ProduceReply{Message: "Successfully pushed to producing queue"}, nil
}

func createGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()

	addr := fmt.Sprintf("%s:%s", host, *grpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProducerServer(s, &server{})
	log.Printf("Listeneing on %s...\n", addr)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.GracefulStop()
}
