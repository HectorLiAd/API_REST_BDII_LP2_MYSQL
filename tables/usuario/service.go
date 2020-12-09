package usuario

import (
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/helper"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"github.com/API_REST_BDII_LP2_MYSQL/userlogin"

	"github.com/API_REST_BDII_LP2_MYSQL/tables/usuariologin"
)

/*Service para los usuario*/
type Service interface {
	RegistrarUsuario(params *registerUserRequest) (*models.ResultOperacion, error)
	SubirImagenUsuario(param *subirAvartarRequest) (*models.ResultOperacion, error)
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
	//Validacion de los datos ingresados
	if len(params.Email) == 0 {
		return nil, errors.New("El email es requerido")
	}
	if !helper.ValidarEmailStr(params.Email) {
		return nil, errors.New("El email ingresado no es valido")
	}
	if !helper.ValidarPasswordStr(params.Password) {
		return nil, errors.New("La contraseña ingresada debe contener números, letras y caracteres especiales")
	}
	if len(params.Password) < 6 {
		return nil, errors.New("La contraseña debe contener almenos 6 caracteres")
	}
	// validacaion por BD

	cantPersona, estadoPersona, err := s.repo.BuscarPersona(params.PersonaID)
	if cantPersona <= 0 {
		return nil, errors.New("La persona al cual desea registrar, no exite")
	}
	if estadoPersona == 0 {
		return nil, errors.New("La persona esta eliminado temporalmente")
	}
	if err != nil {
		fmt.Println("Error al en personas ")
		return nil, err
	}
	personaCreada, errUsuarioCrea := s.repo.ChequeoUsuarioCreado(params.PersonaID)
	if errUsuarioCrea != nil {
		return nil, errUsuarioCrea
	}
	if personaCreada > 0 || personaCreada == -1 {
		return nil, errors.New("El usuario ya existe 1")
	}
	/*Verificar si el usuario existe*/
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()
	var repository = usuariologin.NewRepository(databaseConnection)
	_, usuarioCorreo, errEmail := repository.ChequeoExisteUsuario(&params.Email)
	if errEmail != nil {
		fmt.Println("Error al buscar la persona email")
		return nil, errEmail
	}
	if usuarioCorreo > 0 {
		return nil, errors.New("Ya existe un usuario registrado con ese email")
	}
	// Insertar al usuario
	pwdEncriptado, err := helper.EncriptarPassword(params.Password)
	if err != nil {
		return nil, err
	}
	params.Password = pwdEncriptado

	salidaRegistro, errInsert := s.repo.RegistrarUsuario(params)
	return salidaRegistro, errInsert
}

func (s *service) SubirImagenUsuario(param *subirAvartarRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.SubirImagenUsuario(param, userlogin.UsuarioID)
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo el avatar del usario con el id ", userlogin.UsuarioID),
		Codigo:      userlogin.UsuarioID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}
