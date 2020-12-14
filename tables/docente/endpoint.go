package docente

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type docenteRequest struct {
	ID int
}

func makeAddDocenteEndPoint(s Service) endpoint.Endpoint {
	makeAddSucursalEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(docenteRequest)
		return s.RegistrarDocente(&req)
	}
	return makeAddSucursalEndPoint
}

func makeGetDocenteByIDEndPoint(s Service) endpoint.Endpoint {
	makeAddSucursalEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(docenteRequest)
		return s.ObtenerDocentePorID(&req)
	}
	return makeAddSucursalEndPoint
}
