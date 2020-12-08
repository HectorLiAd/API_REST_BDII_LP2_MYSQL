package sucursal

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addSucursalRequest struct {
	Nombre      string
	Direccion   string
	Descripcion string
}

type updateSucursalRequest struct {
	ID          int
	Nombre      string
	Direccion   string
	Descripcion string
}

type getSucursalByIDRequest struct {
	ID int
}

func makeAddSucursalEndPoint(s Service) endpoint.Endpoint {
	makeAddSucursalEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addSucursalRequest)
		return s.InsertarSucursal(&req)
	}
	return makeAddSucursalEndPoint
}

func makeGetAllSucursalEndPoint(s Service) endpoint.Endpoint {
	makeGetAllSucursalEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodoSucursal()
	}
	return makeGetAllSucursalEndPoint
}

func makeUpdateSucursalEndPoint(s Service) endpoint.Endpoint {
	makeUpdateSucursalEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateSucursalRequest)
		return s.ActualizarSucursal(&req)
	}
	return makeUpdateSucursalEndPoint
}

func makeGetSucursalByIDEndPoint(s Service) endpoint.Endpoint {
	makeGetSucursalByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getSucursalByIDRequest)
		return s.ObtenerSucursalPorID(&req)
	}
	return makeGetSucursalByIDEndPoint
}
