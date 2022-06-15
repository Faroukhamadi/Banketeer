docker rmi faroukhamadi/teller-envoy -f
docker build . -t faroukhamadi/teller-envoy:latest
docker push faroukhamadi/teller-envoy:latest