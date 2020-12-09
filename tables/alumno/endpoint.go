package alumno

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addAlumnoRequest struct {
	ID int
}
type getAlumnoByIDRequest struct {
	ID int
}

func makeAddAlumnoEndPoint(s Service) endpoint.Endpoint {
	addAlumnoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addAlumnoRequest)
		return s.AgregarPersonaAlumno(&req)
	}
	return addAlumnoEndPoint
}

func makeGetAlumnoByIDEndPoint(s Service) endpoint.Endpoint {
	getAlumnoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAlumnoByIDRequest)
		return s.ObtenerPersonaAlumnoPorID(&req)
	}
	return getAlumnoByIDEndPoint
}

func makeGetAllAlumnoByIDEndPoint(s Service) endpoint.Endpoint {
	getAllAlumnoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodasLasPersonas()
	}
	return getAllAlumnoByIDEndPoint
}
