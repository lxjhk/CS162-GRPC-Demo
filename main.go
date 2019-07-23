package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "18.212.39.84:50051"
)

func main() {

	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	defer conn.Close()

	c := NewKeyValueStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.StoreValue(ctx, &StoreRequest{Key: "162", Value: "Project 3 is coming"})
	if err != nil {
		log.Fatalf("cound not greet: %v\n", err)
	}

	log.Println("Successfully stored [Key 162] - [Project 3 is coming]")

	log.Printf("Retrieving key [Project 3 is coming] ...")
	r, err = c.GetValue(ctx, &GetRequest{Key: "162"})
	if err != nil {
		log.Fatalf("cound not greet: %v\n", err)
	}
	log.Printf("The value is %s", r.Value)

}
