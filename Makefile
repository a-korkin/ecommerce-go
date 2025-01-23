include .env

prepare: 
	goose up
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres \
			   < scripts/create_test.sql
run_web:
	go run cmd/main.go -w
run_consumer:
	go run cmd/main.go -b
run_grpc:
	go run cmd/main.go -g
test:
	go test ./... -v -cover
seed_data:
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres -d ${DB_NAME} \
			   < scripts/seed.sql
proto:
	protoc --go_out=./internal/proto --go_opt=paths=source_relative \
		--go-grpc_out=./internal/proto --go-grpc_opt=paths=source_relative \
		*.proto
