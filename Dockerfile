#build stage
FROM golang:1.15 AS builder
RUN apt-get install git
WORKDIR /go/src/app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...



#final stage
FROM ubuntu:latest
RUN apt-get update
RUN apt-get install ca-certificates -y
WORKDIR /app
RUN chmod -R 777 .
COPY --from=builder /go/bin /app
ENV MO2_MONGO_URL=mongodb://mongodb:27017
ENTRYPOINT /app/mo2log
LABEL Name=mo2log Version=0.0.2
