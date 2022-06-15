package transactiondb

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent"
	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent/transaction"
	_ "github.com/lib/pq"
)

func CreateTransaction(ctx context.Context, client *ent.Client) (*ent.Transaction, error) {
	Transaction, err := client.Transaction.
		Create().
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating Transaction: %w", err)
	}
	log.Println("Transaction was created: ", Transaction)
	return Transaction, nil
}

func QueryTransaction(ctx context.Context, client *ent.Client, Id int) (*ent.Transaction, error) {
	Transaction, err := client.Transaction.
		Query().
		Where(transaction.ID(Id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying Transaction: %w", err)
	}
	log.Println("Transaction returned: ", Transaction)
	return Transaction, nil
}

func QueryTransactions(ctx context.Context, client *ent.Client) ([]*ent.Transaction, error) {
	Transactions, err := client.Transaction.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying Transactions: %w", err)
	}
	log.Println("Transactions returned: ", Transactions)
	return Transactions, nil
}
