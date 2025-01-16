include .env

run:
	go run cmd/server.go
test:
	go test ./... -v -cover
