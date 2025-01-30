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
prepare_docker_web:
	docker build -t ecommerce_web -f ./Dockerfile.web .
prepare_docker_consumer:
	docker build -t ecommerce_consumer -f ./Dockerfile.consumer .
run_docker_web:
	docker run --env-file .env \
		-e KAFKA_HOST='kafka-kraft:29092' \
		-e GOOSE_DBSTRING='host=postgres port=5432 user=postgres password=${PGPASSWORD} dbname=${DB_NAME} sslmode=disable' \
		--network=ecommerce_default \
		-p 8080:8080 \
		-it --rm ecommerce_web
run_docker_consumer:
	docker run --env-file .env \
		-e KAFKA_HOST="kafka-kraft:29092" \
		-e GOOSE_DBSTRING='host=postgres port=5432 user=postgres password=${PGPASSWORD} dbname=${DB_NAME} sslmode=disable' \
		--network=ecommerce_default \
		-it --rm ecommerce_consumer
run_docker_kafka:
	docker run \
		--name=kafka-kraft \
		--network=ecommerce_default \
		-h kafka-kraft \
		-p 9092:9092 \
		-e KAFKA_NODE_ID=1 \
		-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP='CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT' \
		-e KAFKA_INTER_BROKER_LISTENER_NAME='PLAINTEXT' \
		-e KAFKA_LISTENERS='PLAINTEXT://kafka-kraft:29092,CONTROLLER://kafka-kraft:29093,PLAINTEXT_HOST://localhost:9092' \
		-e KAFKA_ADVERTISED_LISTENERS='PLAINTEXT://kafka-kraft:29092,PLAINTEXT_HOST://localhost:9092' \
		-e KAFKA_JMX_PORT=9101 \
		-e KAFKA_JMX_HOSTNAME=localhost \
		-e KAFKA_PROCESS_ROLES='broker,controller' \
		-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
		-e KAFKA_CONTROLLER_QUORUM_VOTERS='1@kafka-kraft:29093' \
		-e KAFKA_CONTROLLER_LISTENER_NAMES='CONTROLLER' \
		-it --rm apache/kafka
