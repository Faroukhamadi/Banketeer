docker rmi faroukhamadi/teller -f
docker buildx build --platform=linux/amd64 . -t faroukhamadi/teller:latest
docker push faroukhamadi/teller:latest