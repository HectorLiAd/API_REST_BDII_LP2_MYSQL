package verionesrouter

import (
	"database/sql"
	"net/http"

	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/middlew"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/persona"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/rol"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/tipounidad"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/unidadacademica"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/usuario"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/usuariologin"
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
		rolRepository             = rol.NewRepository(db)             //MARIO
	)
	var (
		usuarioService         = usuario.NewService(usuarioRepository)
		personaService         = persona.NewService(personaRepository)
		usuarioLoginService    = usuariologin.NewService(usuarioLoginRepository)
		tipoUnidadService      = tipounidad.NewService(tipoUnidadRepository)
		unidadAcademicaService = unidadacademica.NewService(unidadAcademicaRepository)
		rolService             = rol.NewService(rolRepository)
	)
	r.Mount("/usuario", usuario.MakeHTTPSHandler(usuarioService))
	r.Mount("/usuariologin", usuariologin.MakeHTTPSHandler(usuarioLoginService))
	r.Mount("/persona", persona.MakeHTTPSHandler(personaService))
	r.Mount("/tipoUnidad", middlew.ValidoJWT(tipounidad.MakeHTTPSHandler(tipoUnidadService)))                //PROTEGICO
	r.Mount("/unidadAcademica", middlew.ValidoJWT(unidadacademica.MakeHTTPSHandler(unidadAcademicaService))) //PROTEGICO
	r.Mount("/rol", middlew.ValidoJWT(rol.MakeHTTPSHandler(rolService)))                                     //PROTEGICO

	return r
}
