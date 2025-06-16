package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/arianaw15/ip-sentinel/grpc/country"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "serverport")
)

type server struct {
	pb.UnimplementedCountryServer
}

func (s *server) ValidateCountryByIP(ctx context.Context, req *pb.ValidateCountryRequest) (*pb.ValidateCountryResponse, error) {
	log.Printf("received: %v", req.Ip)
	return &pb.ValidateCountryResponse{Ip: req.Ip, CountryName: "Test Country", IsValid: true}, nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCountryServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
