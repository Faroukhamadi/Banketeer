docker rmi faroukhamadi/customer -f
docker buildx build --platform=linux/amd64 . -t faroukhamadi/customer:latest
docker push faroukhamadi/customer:latest