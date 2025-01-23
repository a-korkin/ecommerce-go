include .env

prepare: 
	goose up
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres \
			   < scripts/create_test.sql
run_web:
	go run cmd/main.go -w
run_consumer:
	go run cmd/main.go -b
test:
	go test ./... -v -cover
seed_data:
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres -d ${DB_NAME} \
			   < scripts/seed.sql
proto:
	protoc --go_out=./internal/grpc --go_opt=paths=source_relative \
		--go-grpc_out=./internal/grpc --go-grpc_opt=paths=source_relative \
		*.proto
