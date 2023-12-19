package main

import (
	//"log"

	"github.com/daunsid/upload-api/pkg/routes"
	_ "github.com/lib/pq"

	//"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	routes.StartServer()
}
