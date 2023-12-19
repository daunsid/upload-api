package migration

import (
	"log"
	//"os"
	"github.com/daunsid/upload-api/pkg/core"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration() {
	cfg := core.LoadConfig()
	//databaseURL := os.Getenv("DATABASE_URL")
	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL not set")
	}
	m, err := migrate.New(
		"file://sql/migration",
		cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check if migration has already been applied
	_, _, err = m.Version()
	if err == nil {
		log.Println("Migration already applied. Skipping.")
		return
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func init() {

	// m, err := migrate.New(
	// 	"file://sql/migration",
	// 	"postgres://postgres:modupe4816@localhost:5432/uploadb?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }
	RunMigration()
}
