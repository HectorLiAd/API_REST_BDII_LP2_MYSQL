package plancurso

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	AgregarPlanCurso(params *addPlanCursoRequest) (*models.ResultOperacion, error)
	ObtenerPlanCursoPorID(param *getPlanCursoByIDRequest) (*PlanCurso, error)
	ActualizarPlanCursoPorID(params *updatePlanCursoRequest) (*models.ResultOperacion, error)
	ObtenerTodoPlanCurso() ([]*PlanCurso, error)
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

func (serv *service) AgregarPlanCurso(params *addPlanCursoRequest) (*models.ResultOperacion, error) {
	planCursoID, rowAffected, err := serv.repo.AgregarPlanCurso(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro el plan curso correctamente con el ID ", planCursoID),
		Codigo:      planCursoID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerPlanCursoPorID(param *getPlanCursoByIDRequest) (*PlanCurso, error) {
	result, err := serv.repo.ObtenerPlanCursoPorID(param)
	return result, err
}

func (serv *service) ActualizarPlanCursoPorID(params *updatePlanCursoRequest) (*models.ResultOperacion, error) {
	rowAffected, err := serv.repo.ActualizarPlanCursoPorID(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente el plan curso con el ID ", params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerTodoPlanCurso() ([]*PlanCurso, error) {
	return serv.repo.ObtenerTodoPlanCurso()
}
