package customerdb

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb/ent"
	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb/ent/customer"
	_ "github.com/lib/pq"
)

func CreateCustomer(ctx context.Context, client *ent.Client) (*ent.Customer, error) {
	customer, err := client.Customer.
		Create().
		SetFirstName("Omar").
		SetLastName("Mejri").
		SetCin("11676228").
		SetPhone(52212382).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating customer: %w", err)
	}
	log.Println("customer was created: ", customer)
	return customer, nil
}

func QueryCustomer(ctx context.Context, client *ent.Client, Id int) (*ent.Customer, error) {
	customer, err := client.Customer.
		Query().
		Where(customer.ID(Id)).
		// `Only` fails if no customer found,
		// or more than 1 customer returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying customer: %w", err)
	}
	log.Println("customer returned: ", customer)
	return customer, nil
}

func QueryCustomers(ctx context.Context, client *ent.Client) ([]*ent.Customer, error) {
	customers, err := client.Customer.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying customers: %w", err)
	}
	log.Println("customers returned: ", customers)
	return customers, nil
}
