package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Faroukhamadi/Banketeer/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["Dude1"] = 20
	newUsers["Dude2"] = 15

	for name, age := range newUsers {
		res, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf("User Details: -NAME: %s -AGE: %d -ID: %d", res.GetName(), res.GetAge(), res.GetId())
	}

	params := &pb.GetUsersParams{}
	res, err := client.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}
	log.Println("user list")
	fmt.Printf("res.GetUsers(): %v\n", res.GetUsers())
}
