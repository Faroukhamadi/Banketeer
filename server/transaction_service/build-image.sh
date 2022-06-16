docker rmi faroukhamadi/transaction -f
docker build --platform=linux/amd64 . -t faroukhamadi/transaction:latest
docker push faroukhamadi/transaction:latest