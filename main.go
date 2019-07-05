package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/zmcnellis/go-api-template/internal/common/config"
	"github.com/zmcnellis/go-api-template/internal/router"
	"github.com/zmcnellis/go-api-template/internal/services"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := config.NewDB()
	defer db.Close()

	resourceSvc := services.NewResourceService(db)
	router := router.NewRouter(resourceSvc)
	handler := cors.Default().Handler(router)

	log.Printf("Starting server on http://localhost:%s...", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
