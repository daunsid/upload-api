
postgres:
	docker run --name UploadAPI -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=modupe4816 -d postgres

createdb:
	docker exec -it UploadAPI createdb --username=postgres --owner=postgres uploaDB

dropdb:
	docker exec -it UploadAPI dropdb uploaDB

migrateup:
	docker run --rm -v C:/Users/danie/daunsi-dev/upload-api/sql/migration/:migration/ --network host migrate/migrate -path=/migration/ -database postgres://postgres:modupe4816@localhost:5432/uploaDB?sslmode=disable --verbose up 2

migratedown:
	docker run --rm -v C:/Users/danie/daunsi-dev/upload-api/sql/migration:/migration --network host migrate/migrate -path=/migration -database "postgres://postgres:modupe4816@localhost:5432/uploadb?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v C:/Users/danie/daunsi-dev/upload-api:/src -w /src sqlc/sqlc generate
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc