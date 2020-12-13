package unidadacademica

import (
	"errors"
	"strconv"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para los sercicios*/
type Service interface {
	AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (*models.ResultOperacion, error)
	ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error)
	ActualizarUnidadAcademicaByID(params *updateUnidadAcademicaRequest) (*models.ResultOperacion, error)
	ObtenerTodasLasUnidadesAcademicas() ([]*UnidadAcademica, error)
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
	result, rowAfected, err := s.repo.AgregarUnidadAcademica(params)
	if err != nil {
		return nil, err
	}
	resultInsert := &models.ResultOperacion{
		Name:        "Se agrego " + params.Nombre + " con el id " + strconv.Itoa(result),
		Codigo:      result,
		RowAffected: rowAfected,
	}
	return resultInsert, nil
}

func (s *service) ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error) {
	resultUniAcad, errUnidadACad := s.repo.ObtenerUnidadAcademicaByID(param)
	if errUnidadACad != nil {
		return nil, errors.New("No hay resultados pipipipipi" + errUnidadACad.Error())
	}
	// fmt.Println(resultUniAcad.TipoUnidadID)
	// resultTipoUnidad, errTipoUnid := s.repo.ObtenerTipoUnidadByID(resultUniAcad.TipoUnidadID)
	// if errTipoUnid != nil {
	// 	return nil, errTipoUnid
	// }
	return resultUniAcad, nil
}

func (s *service) ActualizarUnidadAcademicaByID(params *updateUnidadAcademicaRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.ActualizarUnidadAcademicaByID(params)
	if err != nil {
		return nil, err
	}
	if rowAffected != 1 {
		return nil, errors.New("No se pudo actualizar la unidad academica")
	}
	resultSms := &models.ResultOperacion{
		Name:        "Se actualizo correctamente la unidad academica con el id " + strconv.Itoa(params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultSms, nil
}

func (s *service) ObtenerTodasLasUnidadesAcademicas() ([]*UnidadAcademica, error) {
	return s.repo.ObtenerTodasLasUnidadesAcademicas()
}
