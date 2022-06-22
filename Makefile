test:
	go test ./... 
build:
	go build -ldflags="-s -w" -o main
