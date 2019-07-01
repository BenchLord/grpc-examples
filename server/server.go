package main

import (
	"context"
	"fmt"
	pb "grpcPlayground/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type peopleServer struct {
	People []*pb.Person
}

func (s *peopleServer) GetPeople(in *pb.Empty, stream pb.People_GetPeopleServer) error {
	for _, person := range s.People {
		if err := stream.Send(person); err != nil {
			return err
		}
	}
	return nil
}

func (s *peopleServer) GetPerson(ctx context.Context, id *pb.Id) (*pb.Person, error) {
	for _, person := range s.People {
		if person.Id == id.Value {
			return person, nil
		}
	}
	return nil, fmt.Errorf("Could not find person with id: %d", id)
}

func (s *peopleServer) AddPeople(stream pb.People_AddPeopleServer) error {
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Empty{})
		}
		if err != nil {
			return err
		}
		s.People = append(s.People, person)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterPeopleServer(grpcServer, &peopleServer{})
	grpcServer.Serve(lis)
}
