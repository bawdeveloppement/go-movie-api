postgres:
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.3-alpine 

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root movie_app

dropdb:
	docker exec -it postgres14 dropdb movie_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/movie_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/movie_app?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc