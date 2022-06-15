docker rmi faroukhamadi/customer-envoy -f
docker build . -t faroukhamadi/customer-envoy:latest
docker push faroukhamadi/customer-envoy:latest