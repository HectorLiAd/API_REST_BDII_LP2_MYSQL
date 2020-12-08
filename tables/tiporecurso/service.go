package tiporecurso

import (
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	AgregarTipoRecurso(params *addTipoRecursoRequest) (*models.ResultOperacion, error)
	ObtenerTodoTipoRecurso() ([]*TipoRecurso, error)
	ObtenerTipoRecursoPorID(param *getTipoRecursoByIDRequest) (*TipoRecurso, error)
	ActualizarTipoRecurso(params *updateTipoRecursoRequest) (*models.ResultOperacion, error)
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

func (s *service) AgregarTipoRecurso(params *addTipoRecursoRequest) (*models.ResultOperacion, error) {
	tipoRecursoID, rowAffected, err := s.repo.AgregarTipoRecurso(params)
	if err != nil {
		return nil, errors.New("No se pudo registrar en la BD ")
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente ", params.Nombre, " con el id ", tipoRecursoID),
		Codigo:      tipoRecursoID,
		RowAffected: rowAffected,
	}
	return resultMsg, nil
}

func (s *service) ObtenerTodoTipoRecurso() ([]*TipoRecurso, error) {
	result, err := s.repo.ObtenerTodoTipoRecurso()
	if err != nil {
		return nil, errors.New(fmt.Sprint("No se pudo obtener los registros de la bd ", err.Error()))
	}
	return result, err
}

func (s *service) ObtenerTipoRecursoPorID(param *getTipoRecursoByIDRequest) (*TipoRecurso, error) {
	result, err := s.repo.ObtenerTipoRecursoPorID(param)
	if err != nil {
		return nil, errors.New(fmt.Sprint("No se pudo obtener en la BD el tipo de recurso con id ", param.ID, " ", err.Error()))
	}
	return result, nil
}

func (s *service) ActualizarTipoRecurso(params *updateTipoRecursoRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.ActualizarTipoRecurso(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente ", params.Nombre, " con el id ", params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, nil
}
