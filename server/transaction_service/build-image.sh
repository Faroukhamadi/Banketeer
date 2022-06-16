docker rmi faroukhamadi/transaction -f
docker build --platform=linux/amd64 . -t faroukhamadi/transaction
docker push faroukhamadi/transaction