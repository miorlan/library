package main

// title Library API.
//
// This is a sample library API.
//
// Schemes:
//	http
// BasePath: /
// version 1.0
//
// Consumes:
//  -application/json
//
// Produces:
//	-application/json
//
// Security:
//	-basic
//
//
// swagger:meta

//go:generate swagger generate spec -o ../public/swagger.json --scan-models

import (
	"Projects/config"
	delivery "Projects/delivery/http"
	"Projects/internal/repository"
	usecase "Projects/internal/usecase"
	"Projects/migration"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	cfg := config.LoadConfig()
	db, err := sql.Open("postgres", cfg.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	migration.Migrate(db)

	repo := repository.NewSongRepository(db)
	ucase := usecase.NewSongUsecase(repo, cfg)
	router := delivery.NewRouter(ucase)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
