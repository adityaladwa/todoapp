createdb:
	createdb --username=adityaladwa --owner=adityaladwa todoapp

dropdb:
	dropdb todoapp

dbup:
	migrate -path db/migration -database "postgresql://localhost:5432/todoapp?sslmode=disable" -verbose up

