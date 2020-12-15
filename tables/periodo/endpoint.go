package periodo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addPeriodoRequest struct {
	Nombre   string
	FechaIni string
	FechaFin string
}

type getPeridioByIDRequest struct {
	ID int
}

type updatePeriodoRequest struct {
	ID       int
	Nombre   string
	FechaIni string
	FechaFin string
}

type deletePeriodoByIDRequest struct {
	ID int
}

func makeAddPeriodoEndPoint(s Service) endpoint.Endpoint {
	addPeriodoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPeriodoRequest)
		return s.RegistrarPeriodo(&req)
	}
	return addPeriodoEndPoint
}

func makeGetPeriodoByIDEndPoint(s Service) endpoint.Endpoint {
	getPeriodoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPeridioByIDRequest)
		return s.ObtenerPeriodoPorID(&req)
	}
	return getPeriodoByIDEndPoint
}

func makeUpdatePeriodoByIDEndPoint(s Service) endpoint.Endpoint {
	updatePeriodoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updatePeriodoRequest)
		return s.ActualizarPeriodo(&req)
	}
	return updatePeriodoByIDEndPoint
}

func makeGetAllPeriodoEndPoint(s Service) endpoint.Endpoint {
	getAllPeriodoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodosLosPeriodos()
	}
	return getAllPeriodoEndPoint
}

func makeDeletePeriodoByIDEndPoint(s Service) endpoint.Endpoint {
	DeletePeriodoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deletePeriodoByIDRequest)
		return s.EliminarPeriodoPorID(&req)
	}
	return DeletePeriodoByIDEndPoint
}
