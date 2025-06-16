package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/arianaw15/ip-sentinel/grpc/country"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port      = ":8080"
	defaultIp = "1.178.64.0/23"
)

var (
	addr = flag.String("addr", "localhost"+port, "address for connection")
)

func main() {

	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCountryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ValidateCountryRequest{
		Ip:             defaultIp,
		ValidCountries: []string{"United States", "Jordan", "Peru"},
	}
	r, err := c.ValidateCountryByIP(ctx, req)
	if err != nil {
		log.Fatalf("could not get country: %v", err)
	}
	log.Printf("Country: %s", r.CountryName)
}
