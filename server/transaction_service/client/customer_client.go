package client

import (
	"context"
	"errors"

	"github.com/Faroukhamadi/Banketeer/customer_service/customerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type CustomerClient struct {
}

var (
	customerGrpcService       = "customer:55050"
	customerGrpcServiceClient customerpb.CustomerServiceClient
)

func prepareCustomerGrpcClient(c *context.Context) error {
	conn, err := grpc.DialContext(*c, customerGrpcService, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}...)

	if err != nil {
		customerGrpcServiceClient = nil
		return errors.New("connection to offer gRPC service failed")
	}

	if customerGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	customerGrpcServiceClient = customerpb.NewCustomerServiceClient(conn)
	return nil
}

func (cc *CustomerClient) GetCustomer(c *context.Context, id int) (*customerpb.GetCustomerResponse, error) {
	if err := prepareCustomerGrpcClient(c); err != nil {
		return nil, err
	}

	res, err := customerGrpcServiceClient.GetCustomer(*c, &customerpb.GetCustomerRequest{Id: int32(id)})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}

	return res, nil
}
