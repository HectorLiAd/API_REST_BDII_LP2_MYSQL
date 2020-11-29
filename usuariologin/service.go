package usuariologin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"golang.org/x/crypto/bcrypt"
)

/*Service interface para crear las firmas que se usaran en el enpoint*/
type Service interface {
	IntentoLogin(params *loginUserRequest) (*Usuario, error)
	LoginUsuario(params *loginUserRequest) (interface{}, error)
	PasswordResetPersonaUsuario(params *passwordResetRequest) (*models.ResultOperacion, error)
}

type service struct {
	repo Repository
}

/*NewService permite crear el servicio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

/*Intento Login*/
func (s *service) IntentoLogin(params *loginUserRequest) (*Usuario, error) {
	usuario, encontrado, err := s.repo.ChequeoExisteUsuario(&params.Email)
	if err != nil {
		return nil, err
	}
	if encontrado == 0 {
		return nil, errors.New("Usuario no encontrado")
	}
	passwordBytes := []byte(params.Password)
	passwordBD := []byte(usuario.UsuarioPassword)
	//Verificando la PWD
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return nil, errors.New("Usuario y/o Contraseña invalidos " + err.Error())
	}
	return usuario, nil
}

func (s *service) LoginUsuario(params *loginUserRequest) (interface{}, error) {
	if len(params.Email) == 0 {
		return nil, errors.New("El email del usuario es requerido")
	}
	usuario, err := s.IntentoLogin(params)
	if err != nil {
		return usuario, err
	}

	//JWT
	jwtkey, er := GeneroJWT(usuario)
	if er != nil {
		return "", errors.New("El email del usuario es requerido" + er.Error())
	}
	resp := RespuestaLogin{
		Token: jwtkey,
	}

	// CUARGAR UNA COOKISSS DEL USUARIO PARA ACCEDER DESDE EL FRONT
	// expirationTime := time.Now().Add(24 * time.Hour) //
	// http.SetCookie(w, &http.Cookie{
	// 	Name: "token",
	// })
	return resp, nil
}

/*Actualizar el password del usuario*/
func (s *service) PasswordResetPersonaUsuario(params *passwordResetRequest) (*models.ResultOperacion, error) {
	usuario, resultPersona, err := s.repo.ChequeoExisteUsuarioPersona(params)
	fmt.Println(resultPersona)
	if resultPersona != 1 {
		return nil, errors.New("El usuario solicitado no existe")
	}
	if params.NewPassword != params.ConfirmarPassword {
		return nil, errors.New("El password ingresado no es valida")
	}
	if len(strings.TrimSpace(params.NewPassword)) < 6 {
		return nil, errors.New("La contraseña debe contener almenos 6 caracteres")
	}
	pwdEncriptado, err := helper.EncriptarPassword(params.NewPassword)
	if err != nil {
		return nil, err
	}
	usuario.UsuarioPassword = pwdEncriptado
	resultUpdatePass, err := s.repo.ActualizarPasswordUsuario(usuario)
	if resultUpdatePass == 0 {
		return nil, errors.New("No se pudo actualizar el password")
	}
	return &models.ResultOperacion{
		Name:   "Usuario " + usuario.UsuarioNombre + " su password fue actualizado corractamente",
		Codigo: usuario.UsuarioID,
	}, err
}
