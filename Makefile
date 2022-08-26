postgres_password=password
postgres_docker_port=4321

postgres:
	docker run --name bank -p $(postgres_docker_port):5432 -e POSTGRES_PASSWORD=$(postgres_password) -d postgres:14.5-alpine

createdb:
	docker exec -it bank createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it bank dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:$(postgres_password)@localhost:$(postgres_docker_port)/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:$(postgres_password)@localhost:$(postgres_docker_port)/simple_bank?sslmode=disable" --verbose down

sqlcgen:
	sqlc generate

.PHONY: createdb postgres migrateup dropdb migratedown sqlcgen