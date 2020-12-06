package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	verionesrouter "github.com/API_REST_BDII_LP2_MYSQL/handlers"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()
	fmt.Println(db)
	defer db.Close()
	r := chi.NewRouter()
	r.Mount("/v1", verionesrouter.RouterV1(db))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("PORT :" + port)
	http.ListenAndServe(":"+port, r)
}
