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

func registerUserEndPoint(s Service) endpoint.Endpoint {
	registerUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerUserRequest)
		result, err := s.RegistrarUsuario(trimStrUserRequest(&req))
		return result, err
	}
	return registerUserEndPoint
}
