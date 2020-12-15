package plancurso

import (
	"github.com/API_REST_BDII_LP2_MYSQL/tables/curso"
)

/*PlanCurso nos permite mapear en un JSON*/
type PlanCurso struct {
	ID    int         `json:"id"`
	Curso curso.Curso `json:"curso"`
	// Plan  plan.Plan `json:"plan"`
}
