package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"NewsThrowback-API/config"
	"NewsThrowback-API/db"
	"NewsThrowback-API/handlers"
	"NewsThrowback-API/repositories"
	"NewsThrowback-API/routes"
)

func main() {
	godotenv.Load()

	cfg := config.Load()

	database := db.Connect(cfg.DatabaseURL)
	defer database.Close()

	articleRepo := repositories.NewArticleRepository(database)
	articleHandler := handlers.NewArticleHandler(articleRepo)

	mux := http.NewServeMux()
	routes.Register(mux, articleHandler)

	log.Printf("server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
