package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb"
	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent"
	"github.com/Faroukhamadi/Banketeer/transaction_service/transactionpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	transactionpb.UnimplementedTransactionServiceServer
}

var (
	timeout = 2 * time.Second
)

func error_response(err error) error {
	log.Println("Transaction Service - ERROR:", err.Error())
	return status.Error(codes.Internal, err.Error())
}

func (*server) GetTransaction(ctx context.Context, req *transactionpb.GetTransactionRequest) (*transactionpb.GetTransactionResponse, error) {
	log.Println("Transaction Service - Called GetTransaction - ID: ", req.Id)

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=transaction_service password=faroukhamadi")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	transaction, err := transactiondb.QueryTransaction(c, client, int(req.Id))
	if err != nil {
		return nil, error_response(err)
	}
	return &transactionpb.GetTransactionResponse{Transaction: &transactionpb.Transaction{Id: int32(transaction.ID), SenderAccountId: int32(transaction.SenderAccountId), ReceiverAccountId: int32(transaction.ReceiverAccountId), TellerId: int32(transaction.TellerId), CreatedAt: transaction.CreateTime.String(), UpdatedAt: transaction.UpdateTime.String()}}, nil
}

func (*server) GetTransactions(ctx context.Context, req *transactionpb.GetTransactionsRequest) (*transactionpb.GetTransactionsResponse, error) {
	log.Println("Transaction Service - Called GetTransactions")

	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=transaction_service password=faroukhamadi")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	transactions, err := transactiondb.QueryTransactions(c, client)
	if err != nil {
		return nil, error_response(err)
	}
	var res transactionpb.GetTransactionsResponse
	for _, transaction := range transactions {
		res.Transactions = append(res.Transactions, &transactionpb.Transaction{Id: int32(transaction.ID), SenderAccountId: int32(transaction.SenderAccountId), ReceiverAccountId: int32(transaction.ReceiverAccountId), TellerId: int32(transaction.TellerId), CreatedAt: transaction.CreateTime.String(), UpdatedAt: transaction.UpdateTime.String()})
	}
	return &res, nil
}

// Fix main so that we only initialize client once
func main() {
	log.Println("Running Transaction Service")

	lis, err := net.Listen("tcp", "0.0.0.0:55052")
	if err != nil {
		log.Println("Transaction Service - ERROR:", err.Error())
	}

	s := grpc.NewServer()
	transactionpb.RegisterTransactionServiceServer(s, &server{})

	log.Printf("Server started at %v", lis.Addr().String())

	err = s.Serve(lis)
	if err != nil {
		log.Println("Transaction Service - ERROR:", err.Error())
	}
}
