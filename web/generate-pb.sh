cd src/grpc/customerpb
protoc -I=. customer.proto --grpc-web_out=import_style=typescript,mode=grpcwebtext:.
cd ../tellerpb/
protoc -I=. teller.proto --grpc-web_out=import_style=typescript,mode=grpcwebtext:.
cd ../transactionpb/
protoc -I=. transaction.proto --grpc-web_out=import_style=typescript,mode=grpcwebtext:.