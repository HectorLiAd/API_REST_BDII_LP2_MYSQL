package usuario

import (
	"errors"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"github.com/API_REST_BDII_LP2_MYSQL/usuariologin"
)

/*Service para los usuario*/
type Service interface {
	RegistrarUsuario(params *registerUserRequest) (*models.ResultOperacion, error)
}

type service struct {
	repo Repository
}

/*NewService Permite crear el servicio para poder manipular la data de la BD*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) RegistrarUsuario(params *registerUserRequest) (*models.ResultOperacion, error) {
	personaCreada, err := s.repo.ChequeoUsuarioCreado(params.PersonaID)
	if err != nil {
		return nil, errors.New("El usuario ya existe")
	}
	if personaCreada > 0 || personaCreada == -1 {
		return nil, errors.New("El usuario ya existe")
	}
	if len(strings.TrimSpace(params.Email)) == 0 {
		return nil, errors.New("El email es requerido")
	}
	if len(strings.TrimSpace(params.Password)) < 6 {
		return nil, errors.New("La contraseña debe contener almenos 6 caracteres")
	}
	/*Verificar si el usuario existe*/
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()
	var repository = usuariologin.NewRepository(databaseConnection)
	_, usuarioCorreo, errr := repository.ChequeoExisteUsuario(&params.Email)
	if errr != nil {
		return nil, errr
	}
	if usuarioCorreo > 0 {
		return nil, errors.New("Ya existe un usuario registrado con ese email")
	}

	pwdEncriptado, err := helper.EncriptarPassword(params.Password)
	if err != nil {
		return nil, err
	}
	params.Password = pwdEncriptado

	salidaRegistro, errr := s.repo.InsertoRegistro(params)
	return salidaRegistro, errr
}
