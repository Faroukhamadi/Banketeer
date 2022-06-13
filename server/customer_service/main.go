package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb"
	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb/ent"
	"github.com/Faroukhamadi/Banketeer/customer_service/customerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	customerpb.UnimplementedCustomerServiceServer
}

var (
	timeout = 2 * time.Second
)

func (*server) GetCustomer(ctx context.Context, req *customerpb.GetCustomerRequest) (*customerpb.GetCustomerResponse, error) {
	log.Println("Customer Service - Called GetCustomer - ID:", req.Id)

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=customer_service password=faroukhamadi")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	customer, err := customerdb.QueryCustomer(c, client, int(req.Id))
	if err != nil {
		return nil, error_response(err)
	}
	return &customerpb.GetCustomerResponse{Customer: &customerpb.Customer{Id: int32(customer.ID), FirstName: customer.FirstName, LastName: customer.LastName, Cin: customer.Cin, Phone: int32(customer.Phone)}}, nil
}

func (*server) GetCustomers(ctx context.Context, req *customerpb.GetCustomersRequest) (*customerpb.GetCustomersResponse, error) {
	log.Println("Customer Service - Called GetCustomers")

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=customer_service password=faroukhamadi")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	customers, err := customerdb.QueryCustomers(c, client)
	if err != nil {
		return nil, error_response(err)
	}
	var res customerpb.GetCustomersResponse
	for _, customer := range customers {
		res.Customers = append(res.Customers, &customerpb.Customer{Id: int32(customer.ID), FirstName: customer.FirstName, LastName: customer.LastName, Cin: customer.Cin, Phone: int32(customer.Phone)})
	}
	return &res, nil
}

func error_response(err error) error {
	log.Println("Customer Service - ERROR:", err.Error())
	return status.Error(codes.Internal, err.Error())
}

func main() {
	log.Println("Running Customer Service")

	lis, err := net.Listen("tcp", "0.0.0.0:55050")
	if err != nil {
		log.Println("Customer Service - ERROR:", err.Error())
	}
	// client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=customer_service password=faroukhamadi")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to postgres: %v", err)
	// }
	// defer client.Close()

	s := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(s, &server{})

	log.Printf("Server started at %v", lis.Addr().String())

	err = s.Serve(lis)

	if err != nil {
		log.Println("Customer Service - ERROR:", err.Error())
	}

}
