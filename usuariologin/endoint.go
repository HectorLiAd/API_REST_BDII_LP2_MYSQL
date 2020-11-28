package usuariologin

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loginUserRequest struct {
	Email    string
	Password string
}

func loginUserEndPoint(s Service) endpoint.Endpoint {
	//Permitir la concurrencia con Context
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginUserRequest)
		userLogin, err := s.LoginUsuario(&req)
		return userLogin, err
	}
	return loginUserEndPoint
}

func pruebaUserEndPoint(s Service) endpoint.Endpoint {
	loginUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"mensaje": "tarea creada"}, nil
	}
	return loginUserEndPoint
}
