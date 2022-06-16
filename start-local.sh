#!/bin/bash
cd server/customer_service
go run main.go & P1=$!
cd ../teller_service
go run main.go & P2=$!
cd ../transaction_service
go run main.go & P3=$!
wait $P1 $P2 $P3