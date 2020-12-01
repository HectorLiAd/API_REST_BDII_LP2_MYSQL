package tipounidad

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addTipoUnidadRequest struct {
	Nombre      string
	Descripcion string
}

type getAllGetTipoUnidadRequest struct {
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
