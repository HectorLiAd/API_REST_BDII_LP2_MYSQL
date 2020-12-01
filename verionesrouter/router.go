package verionesrouter

import (
	"database/sql"
	"net/http"

	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/middlew"
	"github.com/API_REST_BDII_LP2_MYSQL/persona"
	"github.com/API_REST_BDII_LP2_MYSQL/tipounidad"
	"github.com/API_REST_BDII_LP2_MYSQL/unidadacademica"
	"github.com/API_REST_BDII_LP2_MYSQL/usuario"
	"github.com/API_REST_BDII_LP2_MYSQL/usuariologin"
	"github.com/go-chi/chi"
)

/*RouterV1 nos permite usar  las rutas del proyecto*/
func RouterV1(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler)
	var (
		usuarioRepository         = usuario.NewRepository(db)         //HECTOR
		personaRepository         = persona.NewRepository(db)         //FABRICIO
		usuarioLoginRepository    = usuariologin.NewRepository(db)    //HECTOR
		tipoUnidadRepository      = tipounidad.NewRepository(db)      //MARIO
		unidadAcademicaRepository = unidadacademica.NewRepository(db) //MARIO
	)
	var (
		usuarioService         = usuario.NewService(usuarioRepository)
		personaService         = persona.NerService(personaRepository)
		usuarioLoginService    = usuariologin.NewService(usuarioLoginRepository)
		tipoUnidadService      = tipounidad.NewService(tipoUnidadRepository)
		unidadAcademicaService = unidadacademica.NewService(unidadAcademicaRepository)
	)
	r.Mount("/usuario", middlew.ValidoJWT(usuario.MakeHTTPSHandler(usuarioService)))
	r.Mount("/usuariologin", usuariologin.MakeHTTPSHandler(usuarioLoginService))
	r.Mount("/persona", persona.MakeHTTPSHandler(personaService))
	r.Mount("/tipoUnidad", tipounidad.MakeHTTPSHandler(tipoUnidadService))
	r.Mount("/unidadAcademica", unidadacademica.MakeHTTPSHandler(unidadAcademicaService))

	return r
}
