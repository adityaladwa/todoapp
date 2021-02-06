.PHONY: postgres migrate

createdb:
	createdb --username=postgres --owner=postgres todoapp

dropdb:
	dropdb todoapp

migrate:
	migrate -source file://db/migration -database "postgres://localhost:5432/todoapp?sslmode=disable" -verbose up

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=verysecret  postgres



