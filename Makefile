.PHONY: migrate up down

export DATABASE_URL=$(shell grep DATABASE_URL .env | cut -d '=' -f2)

# Path to migrations folder
MIGRATIONS_PATH := migrations

# Target to run all migrations
up:
	goose -dir $(MIGRATIONS_PATH) postgres $(DATABASE_URL) up

# Target to Drop all migrations
down:
	goose -dir $(MIGRATIONS_PATH) postgres $(DATABASE_URL) down-to 0


# build go 
build:
	go build -o bin/scan-attendance cmd/main.go
