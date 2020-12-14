package jerarquia

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarJerarquia(params *addJerarquiaRequest) (*models.ResultOperacion, error)
	AgregarJerarquiaPadre(params *addJerarquiParentRequest) (*models.ResultOperacion, error)
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

func (s *service) RegistrarJerarquia(params *addJerarquiaRequest) (*models.ResultOperacion, error) {
	insertID, rowAffected, err := s.repo.RegistrarJerarquia(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro la gerarquia correctamente con el id ", insertID),
		Codigo:      insertID,
		RowAffected: rowAffected,
	}
	return resultMsg, nil
}

func (s *service) AgregarJerarquiaPadre(params *addJerarquiParentRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.AgregarJerarquiaPadre(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se establecio la jerarquia entre el ID hija ", params.JerarquiaID, " y el ID padre ", params.JerarquiaParentID),
		Codigo:      params.JerarquiaID,
		RowAffected: rowAffected,
	}
	return resultMsg, nil
}
