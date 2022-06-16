docker rmi faroukhamadi/customer -f
docker build --platform=linux/amd64 . -t faroukhamadi/customer
docker push faroukhamadi/customer