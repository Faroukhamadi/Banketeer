docker rmi faroukhamadi/transaction-envoy -f
docker build . -t faroukhamadi/transaction-envoy:latest
docker push faroukhamadi/transaction-envoy:latest