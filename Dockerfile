# Dockerfile для запуска проекта
# 1. docker image build -t car-api-grpc:1.0.0 .
# 2. docker container run -d --name car-api-grpc-app -p 50051:50051 car-api-grpc:1.0.0
#    docker image rm car-api-grpc:1.0.0
#    docker container rm -f car-api-grpc-app
FROM golang:1.17.7-alpine3.15
RUN apk add build-base
WORKDIR /usr/src/app
COPY go.mod go.sum* ./
RUN go mod download
EXPOSE 80
EXPOSE 8080
EXPOSE 3000
EXPOSE 50051
COPY . .
RUN go build server.go 
CMD ["./server"]

