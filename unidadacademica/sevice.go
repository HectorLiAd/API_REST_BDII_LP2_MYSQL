package unidadacademica

import (
	"strconv"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para los sercicios*/
type Service interface {
	AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (*models.ResultOperacion, error)
	ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error)
}

type service struct {
	repo Repository
}

/*NewService permite crear el sercicio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (*models.ResultOperacion, error) {
	result, err := s.repo.AgregarUnidadAcademica(params)
	if err != nil {
		return nil, err
	}
	resultInsert := &models.ResultOperacion{
		Name:   "Se agrego " + params.Nombre + " con el id " + strconv.Itoa(result),
		Codigo: result,
	}
	return resultInsert, nil
}

func (s *service) ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error) {
	resultUniAcad, errUnidadACad := s.repo.ObtenerUnidadAcademicaByID(param)
	if errUnidadACad != nil {
		return nil, errUnidadACad
	}
	// fmt.Println(resultUniAcad.TipoUnidadID)
	resultTipoUnidad, errTipoUnid := s.repo.ObtenerTipoUnidadByID(resultUniAcad.TipoUnidadID)
	if errTipoUnid != nil {
		return nil, errTipoUnid
	}
	resultUniAcad.TipoUnidad = *resultTipoUnidad
	return resultUniAcad, nil
}
