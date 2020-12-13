package unidadacademica

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addUnidadAcademicaRequest struct {
	TipoUnidadID int
	Nombre       string
}
type idUnidadAcademicaRequest struct {
	ID int
}

type updateUnidadAcademicaRequest struct {
	ID     int
	Nombre string
}

func makeAddUnidadAcademicaEndPoint(s Service) endpoint.Endpoint {
	addUnidadAcademicaEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addUnidadAcademicaRequest)
		result, err := s.AgregarUnidadAcademica(trimStrAddUnidadAcademicaRequest(&req))
		return result, err
	}
	return addUnidadAcademicaEndPoint
}

func makeGetUnidadAcademicaByIDEndPoint(s Service) endpoint.Endpoint {
	addUnidadAcademicaEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(idUnidadAcademicaRequest)
		result, err := s.ObtenerUnidadAcademicaByID(&req)
		return result, err
	}
	return addUnidadAcademicaEndPoint
}

func makeUpdateUnidadAcademicaEndPoint(s Service) endpoint.Endpoint {
	updateUnidadAcademicaEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUnidadAcademicaRequest)
		result, err := s.ActualizarUnidadAcademicaByID(&req)
		return result, err
	}
	return updateUnidadAcademicaEndPoint
}

func makeGetAllUnidadAcademicaEndPoint(s Service) endpoint.Endpoint {
	updateUnidadAcademicaEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.ObtenerTodasLasUnidadesAcademicas()
		return result, err
	}
	return updateUnidadAcademicaEndPoint
}
