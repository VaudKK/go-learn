package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "learn.grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	//disable ssl
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)

	// timeout for context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	fmt.Println("Creating a new person...")

	createReq := &pb.CreatePersonRequest{
		Name:        "John Wick",
		Email:       "johnwick@mail.com",
		PhoneNumber: "123-465-6645",
	}

	createRes, err := client.Create(ctx, createReq)

	if err != nil {
		log.Fatalf("Error during create %v", err)
	}

	fmt.Printf("Person created: %+v\n", createRes)

	//read example

	personReq := &pb.SinglePersonRequest{
		Id: 1,
	}

	res, err := client.Read(ctx, personReq)

	if err != nil {
		log.Fatalf("Error fetching person details: %v", err)
	}

	fmt.Printf("Fetched details of person %+v\n", res)
}
