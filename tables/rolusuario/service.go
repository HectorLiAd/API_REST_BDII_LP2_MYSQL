package rolusuario

import (
	"strconv"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	AgregarRolUsuario(params *addRolUsuarioRequest) (*models.ResultOperacion, error)
	ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error)
	ObtenerTodosRolUsuario() ([]*RolUsuario, error)
}

type service struct {
	repo Repository
}

/*NewService Permite crear un nuevo servicio teniendo una nuevo el repositorio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) AgregarRolUsuario(params *addRolUsuarioRequest) (*models.ResultOperacion, error) {
	rolUsuarioID, rowaffected, err := s.repo.AgregarRolUsuario(params)
	if err != nil {
		return nil, err
	}
	resultSms := &models.ResultOperacion{
		Name:        "Se agrego correctamente el rol usuario con el ID " + strconv.Itoa(rolUsuarioID),
		Codigo:      rolUsuarioID,
		RowAffected: rowaffected,
	}
	return resultSms, nil
}

func (s *service) ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error) {
	result, err := s.repo.ObtenerRolUsuarioPorID(param)
	return result, err
}

func (s *service) ObtenerTodosRolUsuario() ([]*RolUsuario, error) {
	return s.repo.ObtenerTodosRolUsuario()
}
