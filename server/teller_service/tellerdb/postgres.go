package tellerdb

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent"
	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent/teller"
	_ "github.com/lib/pq"
)

func CreateTeller(ctx context.Context, client *ent.Client) (*ent.Teller, error) {
	teller, err := client.Teller.
		Create().
		SetUsername("Nardine").
		SetPassword("$argon2id$v=19$m=12,t=3,p=1$Ymh1Mjh1cjBiZDUwMDAwMA$YAt1Xa7y4nUzCkgAdYqhSg").
		SetRole("customer").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating teller: %w", err)
	}
	log.Println("teller was created: ", teller)
	return teller, nil
}

func QueryTeller(ctx context.Context, client *ent.Client, Id int) (*ent.Teller, error) {
	teller, err := client.Teller.
		Query().
		Where(teller.ID(Id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying teller: %w", err)
	}
	log.Println("teller returned: ", teller)
	return teller, nil
}

func QueryTellers(ctx context.Context, client *ent.Client) ([]*ent.Teller, error) {
	tellers, err := client.Teller.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying tellers: %w", err)
	}
	log.Println("tellers returned: ", tellers)
	return tellers, nil
}
