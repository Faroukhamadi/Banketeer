package client

import (
	"context"
	"errors"

	"github.com/Faroukhamadi/Banketeer/teller_service/tellerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type tellerClient struct {
}

var (
	tellerGrpcService       = "teller:55051"
	tellerGrpcServiceClient tellerpb.TellerServiceClient
)

func prepareTellerGrpcClient(c *context.Context) error {
	conn, err := grpc.DialContext(*c, tellerGrpcService, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}...)

	if err != nil {
		tellerGrpcServiceClient = nil
		return errors.New("connection to teller gRPC service failed")
	}

	if tellerGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	tellerGrpcServiceClient = tellerpb.NewTellerServiceClient(conn)
	return nil
}

func (tc *tellerClient) GetTeller(c *context.Context, id int) (*tellerpb.GetTellerResponse, error) {
	if err := prepareTellerGrpcClient(c); err != nil {
		return nil, err
	}

	res, err := tellerGrpcServiceClient.GetTeller(*c, &tellerpb.GetTellerRequest{Id: int32(id)})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return res, nil
}
