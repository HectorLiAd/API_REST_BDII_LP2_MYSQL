package docente

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarDocente(param *docenteRequest) (*models.ResultOperacion, error)
	ObtenerDocentePorID(param *docenteRequest) (*Docente, error)
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

func (s *service) RegistrarDocente(param *docenteRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.RegistrarDocente(param)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente al docente con el id ", param.ID),
		Codigo:      param.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (s *service) ObtenerDocentePorID(param *docenteRequest) (*Docente, error) {
	docente, err := s.repo.ObtenerDocentePorID(param)
	if err != nil {
		return nil, err
	}
	persona, err := s.repo.ObtenerPersonaPorID(param.ID)
	if err != nil {
		return nil, err
	}
	docente.Persona = persona
	return docente, nil
}
