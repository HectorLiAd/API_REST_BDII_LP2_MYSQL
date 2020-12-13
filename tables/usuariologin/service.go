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
	LoginUsuario(params *loginUserRequest) (*RespuestaLogin, error)
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
	personaActivada, err := s.repo.EstadoEliminadoPersona(usuario.UsuarioID)
	if personaActivada == 0 {
		return nil, errors.New("Usuario temporalmente eliminado")
	}
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
	// Obtener el rol del usuario
	rolUsuario, err := s.repo.ObtenerRolUsuario(usuario.UsuarioID)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al querer obtener su rol ", err))
	}
	usuario.Rol = rolUsuario
	return usuario, nil
}

func (s *service) LoginUsuario(params *loginUserRequest) (*RespuestaLogin, error) {
	if len(params.Email) == 0 {
		return nil, errors.New("El email del usuario es requerido")
	}
	usuario, err := s.IntentoLogin(params)
	if err != nil {
		return nil, err
	}

	//JWT
	jwtkey, er := GeneroJWT(usuario)
	if er != nil {
		return nil, errors.New("El email del usuario es requerido" + er.Error())
	}
	resp := &RespuestaLogin{
		Token: jwtkey,
	}
	return resp, nil
}

/*Actualizar el password del usuario*/
func (s *service) PasswordResetPersonaUsuario(params *passwordResetRequest) (*models.ResultOperacion, error) {
	persona, resultPersona, err := s.repo.ChequeoExisteUsuarioPersona(params)
	if resultPersona != 1 {
		return nil, errors.New("El usuario solicitado no existe")
	}
	if params.NewPassword != params.ConfirmarPassword {
		return nil, errors.New("El password ingresado no es valida")
	}
	if len(strings.TrimSpace(params.NewPassword)) < 6 {
		return nil, errors.New("La contraseña debe contener almenos 6 caracteres")
	}
	/*Verificacion si la contraseña es diferente a la anterior*/
	usuario, _, err := s.repo.ChequeoExisteUsuario(&params.Email)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al querer obtener al usuario ", err))
	}
	passwordBytes := []byte(params.NewPassword)
	passwordBD := []byte(usuario.UsuarioPassword)
	if err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes); err == nil {
		return nil, errors.New("Utilize una contraseña diferente a la anterior ")
	}
	/*Procesos de actualizacion de password*/
	pwdEncriptado, err := helper.EncriptarPassword(params.NewPassword)
	if err != nil {
		return nil, err
	}
	persona.UsuarioPassword = pwdEncriptado
	resultUpdatePass, err := s.repo.ActualizarPasswordUsuario(persona)
	if resultUpdatePass == 0 {
		return nil, errors.New("No se pudo actualizar el password")
	}
	return &models.ResultOperacion{
		Name:        "Usuario " + persona.UsuarioNombre + " su password fue actualizado corractamente",
		Codigo:      persona.UsuarioID,
		RowAffected: resultUpdatePass,
	}, err
}
