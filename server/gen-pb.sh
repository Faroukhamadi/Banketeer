cd customer_service/customerpb
protoc --go_out=. --go-grpc_out=. customer.proto
cd ../..
cd teller_service/tellerpb
protoc --go_out=. --go-grpc_out=. teller.proto
cd ../..
# cd transaction_service/transactionpb
# protoc --go_out=. --go-grpc_out=. transaction.proto