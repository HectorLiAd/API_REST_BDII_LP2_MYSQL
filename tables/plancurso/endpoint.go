package plancurso

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

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

type updatePlanCursoRequest struct {
	ID      int
	CursoID int
	PlanID  int
	Ciclo   string
}

func makeAddPlanCursoEndPoint(s Service) endpoint.Endpoint {
	addPlanCursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPlanCursoRequest)
		return s.AgregarPlanCurso(&req)
	}
	return addPlanCursoEndPoint
}

func makeGetPlanCursoByIDEndPoint(s Service) endpoint.Endpoint {
	getPlanCursoByIDEndPoin := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPlanCursoByIDRequest)
		return s.ObtenerPlanCursoPorID(&req)
	}
	return getPlanCursoByIDEndPoin
}

func makeUpdatePlanCursoByIDEndPoint(s Service) endpoint.Endpoint {
	updatePlanCursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updatePlanCursoRequest)
		return s.ActualizarPlanCursoPorID(&req)
	}
	return updatePlanCursoByIDEndPoint
}

func makeGetAllPlanCursoEndPoint(s Service) endpoint.Endpoint {
	getAllPlanCursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodoPlanCurso()
	}
	return getAllPlanCursoByIDEndPoint
}
