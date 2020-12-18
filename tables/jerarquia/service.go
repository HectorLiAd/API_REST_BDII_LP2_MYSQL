package jerarquia

import (
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarJerarquia(params *addJerarquiaRequest) (*models.ResultOperacion, error)
	AgregarJerarquiaPadre(params *addJerarquiParentRequest) (*models.ResultOperacion, error)
	ObtenerJerarquiaPorID(param *getJerarquiaByIDRequest) (*Jerarquia, error)
	ObtenerTodasLasJerarquias() ([]*Jerarquia, error)
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

func (s *service) ObtenerJerarquiaPorID(param *getJerarquiaByIDRequest) (*Jerarquia, error) {
	jerarquia, err := s.ObtenerJerarquiaRecursivoPorID(param)
	if err != nil {
		return nil, err
	}
	if jerarquia.TotaJerarquiaslHijas > 0 {
		jerarqHijosIDs, err := s.repo.ObtenerJerarquiaIDsHijos(jerarquia.ID)
		if err != nil {
			return nil, err
		}
		var jerarquias []*Jerarquia
		for i := 0; i < len(jerarqHijosIDs); i++ {
			param.ID = jerarqHijosIDs[i]
			jerarquiaHija, err := s.ObtenerJerarquiaPorID(param)
			if err != nil {
				return nil, err
			}
			jerarquias = append(jerarquias, jerarquiaHija)
		}
		jerarquia.Jerarquia = jerarquias
		jerarquia.JerarquiaID = 0
		return jerarquia, nil
	}
	jerarquia.JerarquiaID = 0
	return jerarquia, err
}

func (s *service) ObtenerJerarquiaRecursivoPorID(param *getJerarquiaByIDRequest) (*Jerarquia, error) {
	jerarquia, err := s.repo.ObtenerJerarquiaPorID(param)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al consultas de la BD ", err))
	}
	unidadAcad, err := s.repo.ObtenerUnidadAcademicaPorID(jerarquia.UnidadacademicaID)
	jerarquia.UnidadacademicaID = 0
	if err != nil {
		return nil, err
	}
	jerarquia.UnidadAcademica = unidadAcad
	sucur, err := s.repo.ObtenerSucursalPorID(jerarquia.SucursalID)
	jerarquia.SucursalID = 0
	if err != nil {
		return nil, err
	}
	jerarquia.Sucursal = sucur
	totalHijas, err := s.repo.TotalJerarquiaHijas(param.ID)
	if err != nil {
		return nil, err
	}
	jerarquia.TotaJerarquiaslHijas = totalHijas
	return jerarquia, err
}

func (s *service) ObtenerTodasLasJerarquias() ([]*Jerarquia, error) {
	jerarquiasIDs, err := s.repo.ObtenerTodasLasJerarquias()
	if err != nil {
		return nil, err
	}
	jeararquiaID := &getJerarquiaByIDRequest{}
	var jerarquias []*Jerarquia
	for i := 0; i < len(jerarquiasIDs); i++ {
		jeararquiaID.ID = jerarquiasIDs[i]
		jerarquia, err := s.ObtenerJerarquiaPorID(jeararquiaID)
		if err != nil {
			return nil, err
		}
		jerarquias = append(jerarquias, jerarquia)
	}
	return jerarquias, err
}
