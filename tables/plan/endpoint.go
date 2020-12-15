package plan

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addPlanRequest struct {
	JerarquiaID int
	Nombre      string
	Descripcion string
}

type getPlanByIDRequest struct {
	ID int
}

type updatePlanRequest struct {
	ID          int
	Nombre      string
	Descripcion string
}

func makeAddPlanEndPoint(s Service) endpoint.Endpoint {
	addPlanEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPlanRequest)
		return s.RegistrarPlan(&req)
	}
	return addPlanEndPoint
}

func makeGetPlanEndPoint(s Service) endpoint.Endpoint {
	getPlanEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPlanByIDRequest)
		return s.ObtenerPlanPorID(&req)
	}
	return getPlanEndPoint
}

func makeUpdatePlanEndPoint(s Service) endpoint.Endpoint {
	updatePlanEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updatePlanRequest)
		return s.ActualizarPlan(&req)
	}
	return updatePlanEndPoint
}

func makeGetAllPlanEndPoint(s Service) endpoint.Endpoint {
	updatePlanEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodosLosPlanes()
	}
	return updatePlanEndPoint
}
