docker rmi faroukhamadi/customer -f
docker build --platform=linux/amd64 . -t faroukhamadi/customer:latest
docker push faroukhamadi/customer:latest