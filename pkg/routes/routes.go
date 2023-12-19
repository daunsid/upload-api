package routes

import (
	"log"
	"net/http"

	"github.com/daunsid/upload-api/pkg/controller"
	"github.com/daunsid/upload-api/pkg/core"
	_ "github.com/daunsid/upload-api/pkg/migration"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	//"github.com/golang-migrate/migrate/v4"
	//_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//_ "github.com/golang-migrate/migrate/v4/source/file"
)

var UploadServiceRouter = func() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiConfig := controller.NewUpload(controller.Connect())
	v1Router := chi.NewRouter()

	v1Router.Post("/user", apiConfig.HandlerCreateUser)
	v1Router.Get("/healthz", controller.HandlerReadiness)
	v1Router.Post("/upload/{userID}", apiConfig.UploadHandler)
	v1Router.Get("/files/{userID}", apiConfig.ListEntriesHandler)
	v1Router.Get("/download/{fileID}", apiConfig.DownloadHandler)

	router.Mount("/v1", v1Router)
	return router
}

// func init() {

// 	m, err := migrate.New(
// 		"file://sql/migration",
// 		"postgres://postgres:modupe4816@localhost:5432/uploadb?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := m.Up(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func StartServer() {
	router := UploadServiceRouter()

	cfg := core.LoadConfig()
	port := cfg.Port
	srv := http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server stating on port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	println("port: ", port)
}
