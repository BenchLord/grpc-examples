package main

import (
	"context"
	"fmt"
	pb "grpcPlayground/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close()

	client := pb.NewPeopleClient(conn)

	newPeople := make([]*pb.Person, 0)
	newPeople = append(newPeople, &pb.Person{
		Name: "Jim",
		Age:  40,
		Id:   1,
	})
	newPeople = append(newPeople, &pb.Person{
		Name: "Dwight",
		Age:  42,
		Id:   2,
	})

	fmt.Println("Adding new people from the client!")
	s2, err := client.AddPeople(context.Background())
	if err != nil {
		log.Fatalf("Error adding people: %v", err)
	}
	for _, person := range newPeople {
		if err := s2.Send(person); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error adding people: %v", err)
		}
	}
	_, err = s2.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error adding people: %v", err)
	}
	fmt.Println("People added!")

	fmt.Println("Getting all people!")
	s1, err := client.GetPeople(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error getting people: %v", err)
	}
	for {
		person, err := s1.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Eror reading stream: %v", err)
		}
		log.Println(person)
	}

	fmt.Println("Getting person with the id of 1!")
	person, err := client.GetPerson(context.Background(), &pb.Id{Value: 1})
	if err != nil {
		log.Fatalf("Error getting person: %v", err)
	}
	fmt.Println(person)
}
