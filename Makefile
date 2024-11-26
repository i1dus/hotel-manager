# Makefile for "hotel-management"

ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=postgres password=vdJ\#cZ8s dbname=hotel_management host=localhost port=5432 sslmode=disable
endif

DATABASE_URL=postgresql://postgres:vdJ%23cZ8s@localhost:5432/hotel_management?sslmode=disable
MIGRATION_FOLDER=$(CURDIR)/migrations

migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

migration-up: .migration-up
migration-down: .migration-down
db-reset: .migration-down .migration-up

jet:
	jet -dsn=$(DATABASE_URL) \
		-path=./internal/gen \
		-ignore-tables="goose_db_version"
