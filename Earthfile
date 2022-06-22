VERSION 0.6
FROM golang:1.18.3-alpine
RUN apk add build-base
WORKDIR app

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum
    COPY ./*.go ./
    COPY --dir hello ./

test:
    FROM +deps
    RUN go test ./...

build:
    FROM +deps
    RUN go build -ldflags="-s -w" -o main
    SAVE ARTIFACT main AS LOCAL .
