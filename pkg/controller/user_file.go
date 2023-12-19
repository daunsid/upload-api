package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/daunsid/upload-api/internal/db"
	"github.com/daunsid/upload-api/pkg/core"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://postgres:modupe4816@localhost:5432/bankdb?sslmode=disable"
// )

type ApiConfig struct {
	DB *db.Queries
}

func Connect() *sql.DB {

	cfg := core.LoadConfig()

	fmt.Println(cfg.DBURL, "here")
	if cfg.DBURL == "" {
		log.Fatal("db url not found in environment")
	}
	conn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	return conn

}

func NewUpload(dbase *sql.DB) ApiConfig {
	queries := db.New(dbase)
	apiCfg := ApiConfig{
		DB: queries,
	}
	return apiCfg
}
