// package tellerservice
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent"
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

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=teller_service password=faroukhamadi")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating: %v", err)
	}
}
