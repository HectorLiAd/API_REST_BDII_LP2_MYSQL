package main

import (
	"net/http"
	"os"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/handler"
	"github.com/API_REST_BDII_LP2_MYSQL/persona"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	r := chi.NewRouter()
	r.Use(handler.GetCors().Handler)

	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NerService(personaRepository)

	r.Mount("/persona", persona.MakeHTTPHandler(personaServicio))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}
