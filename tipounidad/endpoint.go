package tipounidad

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addTipoUnidadRequest struct {
	Nombre      string
	Descripcion string
}

type getTipoUnidadByIDRequest struct {
	ID int
}

type updateTipoUnidadRequest struct {
	ID          int
	Nombre      string
	Descripcion string
}

func makeAddTipoUnidadEndPoint(s Service) endpoint.Endpoint {
	addTipoUnidadEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addTipoUnidadRequest)
		result, err := s.CrearTipoUnidad(trimStrAddTipoUnidadRequest(&req))
		return result, err
	}
	return addTipoUnidadEndPoint
}

func makeGetAllTipoUnidadEndPoint(s Service) endpoint.Endpoint {
	addTipoUnidadEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.ObtenerRegistrosTipoUnidad()
		return result, err
	}
	return addTipoUnidadEndPoint
}

func makeGetTipoUnidadByIDEndPoint(s Service) endpoint.Endpoint {
	addTipoUnidadEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTipoUnidadByIDRequest)
		result, err := s.ObtenerTipoUnidadByID(&req)
		return result, err
	}
	return addTipoUnidadEndPoint
}

func makeUpdateTipoUnidadEndPoint(s Service) endpoint.Endpoint {
	addTipoUnidadEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateTipoUnidadRequest)
		result, err := s.ActualizarTipoUnidad(&req)
		return result, err
	}
	return addTipoUnidadEndPoint
}
