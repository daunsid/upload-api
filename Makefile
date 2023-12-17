sqlc:
	docker run --rm -v C:/Users/danie/daunsi-dev/upload-api:/src -w /src sqlc/sqlc generate
	
.PHONY: sqlc