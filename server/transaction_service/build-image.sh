docker rmi faroukhamadi/transaction -f
docker buildx build --platform=linux/amd64 . -t faroukhamadi/transaction:latest
docker push faroukhamadi/transaction:latest