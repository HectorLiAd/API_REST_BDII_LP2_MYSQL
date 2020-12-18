package verionesrouter

import (
	"database/sql"
	"net/http"

	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/middlew"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/alumno"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/cargaacademica"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/curso"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/docente"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/jerarquia"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/periodo"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/persona"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/plan"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/plancurso"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/rol"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/rolusuario"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/sucursal"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/tiporecurso"
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
		rolRepository             = rol.NewRepository(db)             //HECTOR
		rolUsuarioRepository      = rolusuario.NewRepository(db)      //HECTOR
		sucursalRepository        = sucursal.NewRepository(db)        //HECTOR
		tipoRecursoRepository     = tiporecurso.NewRepository(db)     //HECTOR
		alumnoRepository          = alumno.NewRepository(db)
		docenteRepository         = docente.NewRepository(db)
		jerarquiaRepository       = jerarquia.NewRepository(db)
		periodoRepository         = periodo.NewRepository(db)
		cursoRepository           = curso.NewRepository(db)
		planRepository            = plan.NewRepository(db)
		planCursoRepository       = plancurso.NewRepository(db)
		cargaAcadRepository       = cargaacademica.NewRepository(db)
	)
	var (
		usuarioService         = usuario.NewService(usuarioRepository)
		personaService         = persona.NewService(personaRepository)
		usuarioLoginService    = usuariologin.NewService(usuarioLoginRepository)
		tipoUnidadService      = tipounidad.NewService(tipoUnidadRepository)
		unidadAcademicaService = unidadacademica.NewService(unidadAcademicaRepository)
		rolService             = rol.NewService(rolRepository)
		rolUsuarioService      = rolusuario.NewService(rolUsuarioRepository)
		sucursalService        = sucursal.NewService(sucursalRepository)
		tipoRecursoService     = tiporecurso.NewService(tipoRecursoRepository)
		alumnoService          = alumno.NewService(alumnoRepository)
		docenteService         = docente.NewService(docenteRepository)
		jerarquiaService       = jerarquia.NewService(jerarquiaRepository)
		periodoService         = periodo.NewService(periodoRepository)
		cursoService           = curso.NewService(cursoRepository)
		planService            = plan.NewService(planRepository)
		planCursoService       = plancurso.NewService(planCursoRepository)
		cargaAcadService       = cargaacademica.NewService(cargaAcadRepository)
	)
	r.Mount("/usuario", middlew.ValidoJWT(usuario.MakeHTTPSHandler(usuarioService)))
	r.Mount("/usuariologin", usuariologin.MakeHTTPSHandler(usuarioLoginService))
	r.Mount("/persona", middlew.ValidoJWT(persona.MakeHTTPSHandler(personaService)))
	r.Mount("/tipoUnidad", middlew.ValidoJWT(tipounidad.MakeHTTPSHandler(tipoUnidadService)))                //PROTEGICO
	r.Mount("/unidadAcademica", middlew.ValidoJWT(unidadacademica.MakeHTTPSHandler(unidadAcademicaService))) //PROTEGICO
	r.Mount("/rol", middlew.ValidoJWT(rol.MakeHTTPSHandler(rolService)))                                     //PROTEGICO
	r.Mount("/rolUsuario", middlew.ValidoJWT(rolusuario.MakeHTTPSHandler(rolUsuarioService)))                //PROTEGICO
	r.Mount("/sucursal", middlew.ValidoJWT(sucursal.MakeHTTPSHandler(sucursalService)))                      //PROTEGICO
	r.Mount("/tipoRecurso", middlew.ValidoJWT(tiporecurso.MakeHTTPSHandler(tipoRecursoService)))             //PROTEGICO
	r.Mount("/alumno", middlew.ValidoJWT(alumno.MakeHTTPSHandler(alumnoService)))                            //PROTEGICO
	r.Mount("/docente", middlew.ValidoJWT(docente.MakeHTTPSHandler(docenteService)))                         //PROTEGICO
	r.Mount("/jerarquia", jerarquia.MakeHTTPSHandler(jerarquiaService))                                      //PROTEGICO
	r.Mount("/periodo", middlew.ValidoJWT(periodo.MakeHTTPSHandler(periodoService)))                         //PROTEGICO
	r.Mount("/curso", middlew.ValidoJWT(curso.MakeHTTPSHandler(cursoService)))                               //PROTEGICO
	r.Mount("/plan", middlew.ValidoJWT(plan.MakeHTTPSHandler(planService)))                                  //PROTEGICO
	r.Mount("/planCurso", middlew.ValidoJWT(plancurso.MakeHTTPSHandler(planCursoService)))                   //PROTEGICO
	r.Mount("/cargaAcademica", middlew.ValidoJWT(cargaacademica.MakeHTTPSHandler(cargaAcadService)))         //PROTEGICO

	return r
}
