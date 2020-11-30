package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/API_REST_BDII_LP2_MYSQL/middlew"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/persona"
	"github.com/API_REST_BDII_LP2_MYSQL/usuario"
	"github.com/API_REST_BDII_LP2_MYSQL/usuariologin"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()
	fmt.Println(db)
	defer db.Close()

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler)
	var (
		usuarioRepository      = usuario.NewRepository(db)
		personaRepository      = persona.NewRepository(db)
		usuarioLoginRepository = usuariologin.NewRepository(db)
	)
	var (
		usuarioServicio     = usuario.NewService(usuarioRepository)
		personaServicio     = persona.NerService(personaRepository)
		usuarioLoginService = usuariologin.NewService(usuarioLoginRepository)
	)
	r.Use(helper.GetCors().Handler)
	r.Mount("/usuario", middlew.ValidoJWT(usuario.MakeHTTPSHandler(usuarioServicio)))
	r.Mount("/usuariologin", usuariologin.MakeHTTPSHandler(usuarioLoginService))
	r.Mount("/persona", persona.MakeHTTPSHandler(personaServicio))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("PORT :" + port)
	http.ListenAndServe(":"+port, r)
}
