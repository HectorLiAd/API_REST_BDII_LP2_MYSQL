package plan

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarPlan(params *addPlanRequest) (*models.ResultOperacion, error)
	ObtenerPlanPorID(param *getPlanByIDRequest) (*Plan, error)
	ActualizarPlan(params *updatePlanRequest) (*models.ResultOperacion, error)
	ObtenerTodosLosPlanes() ([]*Plan, error)
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

func (serv *service) RegistrarPlan(params *addPlanRequest) (*models.ResultOperacion, error) {
	planID, rowAffected, err := serv.repo.RegistrarPlan(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro el plan ", params.Nombre, " correctamente con el id ", planID),
		Codigo:      planID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerPlanPorID(param *getPlanByIDRequest) (*Plan, error) {
	result, err := serv.repo.ObtenerPlanPorID(param)
	return result, err
}

func (serv *service) ActualizarPlan(params *updatePlanRequest) (*models.ResultOperacion, error) {
	rowAffected, err := serv.repo.ActualizarPlan(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente el plan ", params.Nombre, " con el id ", params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerTodosLosPlanes() ([]*Plan, error) {
	return serv.repo.ObtenerTodosLosPlanes()
}
