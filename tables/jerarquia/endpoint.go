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
