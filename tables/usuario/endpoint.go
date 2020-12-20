package usuario

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*Usuario para registrar al usuario*/
type registerUserRequest struct {
	PersonaID int
	UserName  string
	Email     string
	Password  string
}

type subirAvartarRequest struct {
	File string
}

type obtenerAvatarRequest struct {
	ID   int
	File string
}

func makeRegisterUserEndPoint(s Service) endpoint.Endpoint {
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerUserRequest)
		result, err := s.RegistrarUsuario(trimStrUserRequest(&req))
		return result, err
	}
	return registerUserEndPoint
}

func makeSubirImagenUserEndPoint(s Service) endpoint.Endpoint {
	subirImagenUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(subirAvartarRequest)
		return s.SubirImagenUsuario(&req)
	}
	return subirImagenUserEndPoint
}

func makeGetImagenUserEndPoint(s Service) endpoint.Endpoint {
	subirImagenUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(obtenerAvatarRequest)
		result, err := s.BuscarImagenUsuario(&req)
		return result, err
	}
	return subirImagenUserEndPoint
}

func makeGetAllUserEndPoint(s Service) endpoint.Endpoint {
	getAllUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.ObtenerTodosLosUsuarios()
		return result, err
	}
	return getAllUserEndPoint
}
