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
	Avatar    string
}

type subirAvartarRequest struct {
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
