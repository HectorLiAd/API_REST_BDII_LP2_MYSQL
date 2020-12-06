package rol

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addRolRequest struct {
	Nombre string
}

type getRolByIDRequest struct {
	ID int
}

type updateRolRequest struct {
	ID     int
	Nombre string
}

func makeAddRolEndPoint(s Service) endpoint.Endpoint {
	addRolEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRolRequest)
		result, err := s.InsertarRol(trimStrAddRolRequest(&req))
		return result, err
	}
	return addRolEndPoint
}

func makeUpdateRolEndPoint(s Service) endpoint.Endpoint {
	addRolEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRolRequest)
		result, err := s.ActualizarRol(trimStrUpdateRolRequest(&req))
		return result, err
	}
	return addRolEndPoint
}

func makeGetRolByIDEndPoint(s Service) endpoint.Endpoint {
	addRolEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRolByIDRequest)
		result, err := s.ObtenerRolByID(&req)
		return result, err
	}
	return addRolEndPoint
}

func makeGetAllRolEndPoint(s Service) endpoint.Endpoint {
	addRolEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodosLosRoles()
	}
	return addRolEndPoint
}
