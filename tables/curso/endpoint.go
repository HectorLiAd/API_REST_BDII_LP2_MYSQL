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

type updateCursoByIDRequest struct {
	ID      int
	Nombre  string
	Detalle string
}

type updateImagenCursoByIDRequest struct {
	ID   int
	File string
}

type getFondoCursoByIDRequest struct {
	ID   int
	File string
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

func makeUpdateCursoByIDEndPoint(s Service) endpoint.Endpoint {
	updateCursoByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateCursoByIDRequest)
		return s.ActualizatCursoPorID(&req)
	}
	return updateCursoByIDEndPoint
}

func makeGetAllCursoEndPoint(s Service) endpoint.Endpoint {
	getAllCursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.ObtenerTodosLosCursos()
	}
	return getAllCursoEndPoint
}

func makeUploadImageCursoEndPoint(s Service) endpoint.Endpoint {
	getAllCursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateImagenCursoByIDRequest)
		return s.SubirFondoCurso(&req)
	}
	return getAllCursoEndPoint
}

func makeGetFondoCursoEndPoint(s Service) endpoint.Endpoint {
	getFondoCursoEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getFondoCursoByIDRequest)
		return s.ObtenerFondoCurso(&req)
	}
	return getFondoCursoEndPoint
}
