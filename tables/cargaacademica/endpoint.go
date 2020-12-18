package cargaacademica

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addCargaAcademica struct {
	JerarquiaID int
	PersonaID   int
	PeriodoID   int
	PlanCursoID int
}

type getCargaAcadByID struct {
	ID int
}

func makeAddCargaAcadEndPoint(s Service) endpoint.Endpoint {
	addCargaAcadEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addCargaAcademica)
		return s.RegistrarCargaAcad(&req)
	}
	return addCargaAcadEndPoint
}

func makeGetCargaAcadByIDEndPoint(s Service) endpoint.Endpoint {
	getCargaAcadByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCargaAcadByID)
		return s.ObtenerCargaAcadPorID(&req)
	}
	return getCargaAcadByIDEndPoint
}
