package tipounidad

import (
	"errors"
	"strconv"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para los sercicios*/
type Service interface {
	CrearTipoUnidad(params *addTipoUnidadRequest) (*models.ResultOperacion, error)
	ObtenerRegistrosTipoUnidad() ([]*TipoUnidad, error)
	ObtenerTipoUnidadByID(param *getTipoUnidadByIDRequest) (*TipoUnidad, error)
	ActualizarTipoUnidad(params *updateTipoUnidadRequest) (*models.ResultOperacion, error)
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

func (s *service) CrearTipoUnidad(params *addTipoUnidadRequest) (*models.ResultOperacion, error) {
	result, err := s.repo.crearTipoUnidad(params)
	if err != nil {
		return nil, err
	}
	returnResult := &models.ResultOperacion{
		Name:   "Se agrego " + params.Nombre + " con el id " + strconv.Itoa(result),
		Codigo: result,
	}
	return returnResult, nil
}

func (s *service) ObtenerRegistrosTipoUnidad() ([]*TipoUnidad, error) {
	result, err := s.repo.ObtenerTodosLosTiposDeUnidad()
	if err != nil {
		return nil, err
	}
	// resulUnidadAcad, err := s.repo.ObtenerTodaUnidadAcademica(result.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// result.UnidadAcad = resulUnidadAcad
	return result, err
}

func (s *service) ObtenerTipoUnidadByID(param *getTipoUnidadByIDRequest) (*TipoUnidad, error) {
	return s.repo.ObtenerTipoDeUnidadByID(param)
}

func (s *service) ActualizarTipoUnidad(params *updateTipoUnidadRequest) (*models.ResultOperacion, error) {
	resultUpdate, err := s.repo.ActualizarTipoUnidad(params)

	if err != nil {
		return nil, err
	}
	if resultUpdate == 0 {
		return nil, errors.New("No se puedo actualizar el tipo de unidad")
	}
	resultOperacion := &models.ResultOperacion{
		Name:        "Se actualizo correctamente " + params.Nombre + " con el id " + strconv.Itoa(params.ID),
		Codigo:      params.ID,
		RowAffected: resultUpdate,
	}
	return resultOperacion, nil
}
