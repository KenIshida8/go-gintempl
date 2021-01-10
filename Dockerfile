FROM golang:1.15
COPY src src
WORKDIR src
RUN go build
