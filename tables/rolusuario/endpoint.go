package rolusuario

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addRolUsuarioRequest struct {
	RolID     int
	PersonaID int
}

type getRolUsuarioByIDRequest struct {
	ID int
}

func makeAddRolUsuarioEndpoint(s Service) endpoint.Endpoint {
	addRolUsuarioEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRolUsuarioRequest)
		return s.AgregarRolUsuario(&req)
	}
	return addRolUsuarioEndpoint
}

func makeGetRolUsuarioByIDEndpoint(s Service) endpoint.Endpoint {
	getRolUsuarioByIDEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRolUsuarioByIDRequest)
		return s.ObtenerRolUsuarioPorID(&req)
	}
	return getRolUsuarioByIDEndpoint
}

func makeGetAllRolUsuarioEndpoint(s Service) endpoint.Endpoint {
	getAllRolUsuarioEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodosRolUsuario()
	}
	return getAllRolUsuarioEndpoint
}
