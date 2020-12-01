package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/verionesrouter"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()
	fmt.Println(db)
	defer db.Close()
	r := chi.NewRouter()
	r.Mount("/v1", verionesrouter.Router(db))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("PORT :" + port)
	http.ListenAndServe(":"+port, r)
}
