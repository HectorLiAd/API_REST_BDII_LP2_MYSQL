package plancurso

// import (
// 	"context"

// 	"github.com/go-kit/kit/endpoint"
// )
type addPlanCursoRequest struct {
	CursoID int
	PlanID  int
	Ciclo   string
}
type getPlanCursoByIDRequest struct {
	ID int
}

// func makeAddAlumnoEndPoint(s Service) endpoint.Endpoint {
// 	addAlumnoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(addAlumnoRequest)
// 		return s.AgregarPersonaAlumno(&req)
// 	}
// 	return addAlumnoEndPoint
// }
