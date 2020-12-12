package rol

import (
	"errors"
	"strconv"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	InsertarRol(param *addRolRequest) (*models.ResultOperacion, error)
	ActualizarRol(params *updateRolRequest) (*models.ResultOperacion, error)
	ObtenerRolByID(param *getRolByIDRequest) (*Rol, error)
	ObtenerTodosLosRoles() ([]*Rol, error)
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
func (s *service) InsertarRol(param *addRolRequest) (*models.ResultOperacion, error) {
	param.Nombre = strings.ToUpper(param.Nombre)
	// strings.ToUpper("abc")
	if len(param.Nombre) <= 1 {
		return nil, errors.New("El nombre es muy corto")
	}
	rolID, effected, errIR := s.repo.InsertarRol(param)
	if errIR != nil {
		return nil, errors.New("No se pudo registrar a la BD pipipipi " + errIR.Error())
	}
	resulSms := &models.ResultOperacion{
		Name:        "Se registro correctamente el rol " + param.Nombre + " con el id " + strconv.Itoa(int(rolID)),
		Codigo:      int(rolID),
		RowAffected: int(effected),
	}
	return resulSms, nil
}

func (s *service) ActualizarRol(params *updateRolRequest) (*models.ResultOperacion, error) {
	params.Nombre = strings.ToUpper(params.Nombre)
	if len(params.Nombre) < 3 {
		return nil, errors.New("El nombre ingresado es muy corto")
	}
	rowAffected, errAR := s.repo.ActualizarRol(params)
	if errAR != nil {
		return nil, errAR
	}
	if rowAffected != 1 {
		return nil, errors.New("No se pudo actualizar los datos indicados, verifique que esten bien")
	}
	resultSms := &models.ResultOperacion{
		Name:        "Se actualizó correctamante el rol con ID " + strconv.Itoa(params.ID) + " con el nombre " + params.Nombre,
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultSms, nil
}

func (s *service) ObtenerRolByID(param *getRolByIDRequest) (*Rol, error) {
	result, err := s.repo.ObtenerRolByID(param)
	if err != nil {
		return nil, errors.New("No se pudo obtener resultados " + err.Error())
	}
	return result, nil
}

func (s *service) ObtenerTodosLosRoles() ([]*Rol, error) {
	return s.repo.ObtenerTodosLosRoles()
}
