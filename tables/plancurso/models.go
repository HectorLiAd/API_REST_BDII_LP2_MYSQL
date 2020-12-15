package plancurso

import (
	"github.com/API_REST_BDII_LP2_MYSQL/tables/curso"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/plan"
)

/*PlanCurso nos permite mapear en un JSON*/
type PlanCurso struct {
	ID      int          `json:"id"`
	Curso   *curso.Curso `json:"curso"`
	CursoID int          `json:"cursoId,omitempty"`
	Plan    *plan.Plan   `json:"plan"`
	PlanID  int          `json:"planId,omitempty"`
	Ciclo   string       `json:"ciclo"`
}
