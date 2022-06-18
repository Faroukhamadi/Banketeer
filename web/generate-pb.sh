cd src/grpc/customerpb
protoc -I=. customer.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcweb:.

cd ../tellerpb/
protoc -I=. teller.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcweb:.

cd ../transactionpb/
protoc -I=. transaction.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcweb:.