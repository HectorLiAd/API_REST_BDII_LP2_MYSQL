package curso

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type addCursoRequest struct {
	Nombre  string
	Detalle string
}

type getCursoByIDRequest struct {
	ID int
}

func makeAddCursoEndPoint(s Service) endpoint.Endpoint {
	addCursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addCursoRequest)
		return s.RegistrarCurso(&req)
	}
	return addCursoEndPoint
}

func makeGetCursoByIDEndPoint(s Service) endpoint.Endpoint {
	getCursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCursoByIDRequest)
		fmt.Println(req)
		return s.ObtenerCursoPorID(&req)
	}
	return getCursoByIDEndPoint
}
