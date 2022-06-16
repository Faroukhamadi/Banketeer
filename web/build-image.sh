docker rmi faroukhamadi/banketeer-ui -f
docker build --platform=linux/amd64 . -t faroukhamadi/banketeer-ui
docker push faroukhamadi/banketeer-ui