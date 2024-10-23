package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "learn.grpc/proto"

	"google.golang.org/grpc"
)

type Person struct {
	ID          int32
	Name        string
	Email       string
	PhoneNumber string
}

var nextID int32 = 1
var persons = make(map[int32]Person)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) Create(ctx context.Context,in *pb.CreatePersonRequest) (*pb.PersonProfileResponse, error){
	person := Person{
		Name: in.GetName(),
		Email: in.GetName(),
		PhoneNumber: in.GetPhoneNumber(),
	}

	if person.Email == "" || person.Name == "" || person.PhoneNumber == "" {
		return &pb.PersonProfileResponse{},errors.New("fields missing")
	}

	person.ID = nextID
	persons[person.ID] = person
	nextID = nextID + 1

	return &pb.PersonProfileResponse{Id: person.ID,Name: person.Name, Email: person.Email,PhoneNumber: person.PhoneNumber},nil
}

func (s *server) Read(ctx context.Context, in *pb.SinglePersonRequest) (*pb.PersonProfileResponse, error){
	id := in.GetId()

	person := persons[id]

	if person.ID == 0 {
		return &pb.PersonProfileResponse{},errors.New("not found")
	}

	return &pb.PersonProfileResponse{Id: person.ID,Name: person.Name, Email: person.Email,PhoneNumber: person.PhoneNumber},nil
}

func main() {
	lis, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterPersonServiceServer(s,&server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v",err)
	}
}
