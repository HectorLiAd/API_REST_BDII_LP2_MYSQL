package usuariologin

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loginUserRequest struct {
	Email    string
	Password string
}

type passwordResetRequest struct {
	NombrePersonal    string
	ApellidoPaterno   string
	ApellidoMaterno   string
	FechaNacimiento   string
	Email             string
	NewPassword       string
	ConfirmarPassword string
}

func loginUserEndPoint(s Service) endpoint.Endpoint {
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginUserRequest)
		userLogin, err := s.LoginUsuario(&req)
		return userLogin, err
	}
	return loginUserEndPoint
}

func passwordResetEndPoint(s Service) endpoint.Endpoint {
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(passwordResetRequest)
		result, err := s.PasswordResetPersonaUsuario(&req)
		return result, err
	}
	return loginUserEndPoint
}
