// package main

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"time"

// 	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb"
// 	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent"
// 	"github.com/Faroukhamadi/Banketeer/transaction_service/transactionpb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type server struct {
// 	transactionpb.UnimplementedTellerServiceServer
// }

// var (
// 	timeout = 2 * time.Second
// )

// func error_response(err error) error {
// 	log.Println("Teller Service - ERROR:", err.Error())
// 	return status.Error(codes.Internal, err.Error())
// }

// func (*server) GetTeller(ctx context.Context, req *transactionpb.GetTellerRequest) (*transactionpb.GetTellerResponse, error) {
// 	log.Println("Teller Service - Called GetTeller - ID: ", req.Id)

// 	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=transaction_service password=faroukhamadi")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to postgres: %v", err)
// 	}
// 	defer client.Close()

// 	c, cancel := context.WithTimeout(ctx, timeout)
// 	defer cancel()

// 	teller, err := transactiondb.QueryTeller(c, client, int(req.Id))
// 	if err != nil {
// 		return nil, error_response(err)
// 	}
// 	return &transactionpb.GetTellerResponse{Teller: &transactionpb.Teller{Id: int32(teller.ID), Username: teller.Username, Password: teller.Password, Role: string(teller.Role), CreatedAt: teller.CreateTime.String(), UpdatedAt: teller.UpdateTime.String()}}, nil
// }

// func (*server) GetTellers(ctx context.Context, req *transactionpb.GetTellersRequest) (*transactionpb.GetTellersResponse, error) {
// 	log.Println("Teller Service - Called GetTellers")

// 	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=transaction_service password=faroukhamadi")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to postgres: %v", err)
// 	}
// 	defer client.Close()

// 	c, cancel := context.WithTimeout(ctx, timeout)
// 	defer cancel()

// 	tellers, err := transactiondb.QueryTellers(c, client)
// 	if err != nil {
// 		return nil, error_response(err)
// 	}
// 	var res transactionpb.GetTellersResponse
// 	for _, teller := range tellers {
// 		res.Tellers = append(res.Tellers, &transactionpb.Teller{Id: int32(teller.ID), Username: teller.Username, Password: teller.Password, Role: string(teller.Role), CreatedAt: teller.CreateTime.String(), UpdatedAt: teller.UpdateTime.String()})
// 	}
// 	return &res, nil
// }

// // Fix main so that we only initialize client once
// func main() {
// 	log.Println("Running Teller Service")

// 	lis, err := net.Listen("tcp", "0.0.0.0:55051")
// 	if err != nil {
// 		log.Println("Teller Service - ERROR:", err.Error())
// 	}

// 	s := grpc.NewServer()
// 	transactionpb.RegisterTellerServiceServer(s, &server{})

// 	log.Printf("Server started at %v", lis.Addr().String())

// 	err = s.Serve(lis)
// 	if err != nil {
// 		log.Println("Teller Service - ERROR:", err.Error())
// 	}
// }
