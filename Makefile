prepare: 
	goose up
run:
	go run cmd/main.go
test:
	go test ./... -v -cover
