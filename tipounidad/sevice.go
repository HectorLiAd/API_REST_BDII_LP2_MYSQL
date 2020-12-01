package tipounidad

import (
	"strconv"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para los sercicios*/
type Service interface {
	CrearTipoUnidad(params *addTipoUnidadRequest) (*models.ResultOperacion, error)
	ObtenerRegistrosTipoUnidad() ([]*TipoUnidad, error)
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
	return result, err
}
