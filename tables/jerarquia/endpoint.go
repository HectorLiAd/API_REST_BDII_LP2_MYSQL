package jerarquia

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addJerarquiaRequest struct {
	UnidadAcadID int
	SucursalID   int
}

type addJerarquiParentRequest struct {
	JerarquiaID       int
	JerarquiaParentID int
}

type getJerarquiaByIDRequest struct {
	ID int
}

func makeAddJerarquiaEndPoint(s Service) endpoint.Endpoint {
	addJerarquiaEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addJerarquiaRequest)
		return s.RegistrarJerarquia(&req)
	}
	return addJerarquiaEndPoint
}

func makeAddJerarquiaParentByIDEndPoint(s Service) endpoint.Endpoint {
	addJerarquiaParentByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addJerarquiParentRequest)
		return s.AgregarJerarquiaPadre(&req)
	}
	return addJerarquiaParentByIDEndPoint
}

func makeGetJerarquiaByIDEndPoint(s Service) endpoint.Endpoint {
	getJerarquiaParentByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getJerarquiaByIDRequest)
		return s.ObtenerJerarquiaPorID(&req)
	}
	return getJerarquiaParentByIDEndPoint
}

func makeGetAllJerarquiaEndPoint(s Service) endpoint.Endpoint {
	getJerarquiaParentByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodasLasJerarquias()
	}
	return getJerarquiaParentByIDEndPoint
}
