package plancurso

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	AgregarPlanCurso(params *addPlanCursoRequest) (*models.ResultOperacion, error)
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
