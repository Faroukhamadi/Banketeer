package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Faroukhamadi/Banketeer/services/customer/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type CustomerServer struct {
	pb.UnimplementedCustomerServiceServer
	savedCustomers []*pb.Customer
}

func (s *CustomerServer) GetCustomer(ctx context.Context, params *pb.CustomerParams) (*pb.Customer, error) {
	for _, customer := range s.savedCustomers {
		if params.ID == customer.ID {
			return customer, nil
		}
	}
	// no customer was found, return an unnamed feature
	return &pb.Customer{ID: params.ID}, nil
}

func NewServer() *CustomerServer {
	s := &CustomerServer{savedCustomers: }
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCustomerServiceServer(grpcServer, )
}
