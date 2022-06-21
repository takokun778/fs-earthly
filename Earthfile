VERSION 0.6
FROM golang:1.18.3-bullseye
WORKDIR app

test:
    COPY . .
    RUN go test ./...
build:
    COPY . .
    RUN go build -ldflags="-s -w" -o main
    SAVE ARTIFACT main AS LOCAL .
