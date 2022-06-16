docker rmi faroukhamadi/teller -f
docker build --platform=linux/amd64 . -t faroukhamadi/teller:latest
docker push faroukhamadi/teller:latest