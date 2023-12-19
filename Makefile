build:
    docker build -t upload-api .

postgres:
	docker run --name UploadAPI -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=modupe4816 -d postgres

createdb:
	docker exec -it UploadAPI createdb --username=postgres --owner=postgres uploaDB

sqlc:
	docker run --rm -v C:/Users/danie/daunsi-dev/upload-api:/src -w /src sqlc/sqlc generate
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc