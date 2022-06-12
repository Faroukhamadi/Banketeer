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

func QueryCustomer(ctx context.Context, client *ent.Client) (*ent.Customer, error) {
	customer, err := client.Customer.
		Query().
		Where(customer.ID(1)).
		// `Only` fails if no customer found,
		// or more than 1 customer returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying customer: %w", err)
	}
	log.Println("customer returned: ", customer)
	return customer, nil
}

// func main() {
// 	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=customer_service password=faroukhamadi")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to postgres: %v", err)
// 	}
// 	defer client.Close()
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating  : %v", err)
// 	}
// }
