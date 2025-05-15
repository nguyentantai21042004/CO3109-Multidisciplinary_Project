include .env
export
BINARY=engine
DB_CONTAINER=postgres_db
DB_NAME=golangdb
DB_USER=admin
MIGRATION_FOLDER=./database/migrations
SCHEMA_FILE=./database/init_schema.sql

### ##### Database PostgreSQL ##### ###
create-db:
	sqlboiler psql --wipe --output ./internal/models --pkgname models

container-db:
	docker compose up -d postgres_db

# Copy file schema vào container (Chạy duy nhất 1 lần)
copy-schema:
	docker cp $(SCHEMA_FILE) $(DB_CONTAINER):/init_schema.sql

# Chạy file schema khởi tạo ban đầu (Chỉ chạy 1 lần)
migrate-init:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -f /init_schema.sql

# Copy tất cả các file migration vào container
copy-migration:
	docker cp $(MIGRATION_FOLDER) $(DB_CONTAINER):/migrations

# Chạy tất cả các file migration trong container
migrate-all: copy-migration
	docker exec -it $(DB_CONTAINER) bash -c 'for file in /migrations/**/*.sql; do psql -U $(DB_USER) -d $(DB_NAME) -f $$file; done'

# Xóa các file trong container (tùy chọn)
clean:
	docker exec -it $(DB_CONTAINER) rm -rf /init_schema.sql /migrations

models:
	@echo "Generating models"
	sqlboiler psql

### ##### Server ##### ###
## Chạy swagger
swagger:
	@echo "Generating swagger"
	@swag init -g cmd/api/main.go

## Run the application
run-api:
	@echo "Generating swagger"
	@swag init -g cmd/api/main.go
	@echo "Running the application"
	@go run cmd/api/main.go

## Run the consumer
run-consumer:
	@echo "Running the consumer"
	@go run cmd/consumer/main.go

## Build docker compose
build-docker-compose:
	@echo "make models first"
	@make models
	@echo "Building docker compose"
	docker compose up --build -d
