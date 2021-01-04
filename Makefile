MIGRATE=migrate -path db/migrations -database postgres://postgres:secret@localhost:5432/goal?sslmode=disable

.PHONY: migrate-make migrate-up migrade-down

all: db-up migrate-up

up-db:
	docker run --name goal-db -v $(PWD)/docker/postgres/init.sh:/docker-entrypoint-initdb.d/init.sh -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

create-db:
	docker exec -it goal-db createdb --username=postgres --owner=postgres goal

drop-db:
	docker exec -it goal-db dropdb --username=postgres goal

migrate-make:
	migrate create -ext sql -dir db/migrations $(word 2, $(MAKECMDGOALS))

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down