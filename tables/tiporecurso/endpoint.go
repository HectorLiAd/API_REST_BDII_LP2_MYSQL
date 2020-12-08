package tiporecurso

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addTipoRecursoRequest struct {
	Nombre             string
	EstadoCalificativo int
	BloquearRecurso    int
}

type getTipoRecursoByIDRequest struct {
	ID int
}

type updateTipoRecursoRequest struct {
	ID                 int
	Nombre             string
	EstadoCalificativo int
	BloquearRecurso    int
}

func makeAddTipoRecursoEndPoint(s Service) endpoint.Endpoint {
	makeAddTipoRecursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addTipoRecursoRequest)
		return s.AgregarTipoRecurso(&req)
	}
	return makeAddTipoRecursoEndPoint
}

func makeGetAllTipoRecursoEndPoint(s Service) endpoint.Endpoint {
	makeGetAllTipoRecursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodoTipoRecurso()
	}
	return makeGetAllTipoRecursoEndPoint
}

func makeGetTipoRecursoByIDEndPoint(s Service) endpoint.Endpoint {
	makeGetTipoRecursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTipoRecursoByIDRequest)
		return s.ObtenerTipoRecursoPorID(&req)
	}
	return makeGetTipoRecursoByIDEndPoint
}

func makeUpdateTipoRecursoByIDEndPoint(s Service) endpoint.Endpoint {
	makeUpdateTipoRecursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateTipoRecursoRequest)
		return s.ActualizarTipoRecurso(&req)
	}
	return makeUpdateTipoRecursoByIDEndPoint
}
