include .env

prepare: 
	goose up
run_web:
	go run cmd/main.go -w
run_consumer:
	go run cmd/main.go -b
test:
	go test ./... -v -cover
seed_data:
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres -d \
			   ${DB_NAME} < scripts/seed_categories_and_products.sql && \
	PGPASSWORD=${PGPASSWORD} psql -h localhost -U postgres -d \
			   ${DB_NAME} < scripts/seed_users.sql
	
