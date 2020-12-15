package periodo

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarPeriodo(params *addPeriodoRequest) (*models.ResultOperacion, error)
	ObtenerPeriodoPorID(param *getPeridioByIDRequest) (*Periodo, error)
	ActualizarPeriodo(params *updatePeriodoRequest) (*models.ResultOperacion, error)
	ObtenerTodosLosPeriodos() ([]*Periodo, error)
	EliminarPeriodoPorID(param *deletePeriodoByIDRequest) (*models.ResultOperacion, error)
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

func (serv *service) RegistrarPeriodo(params *addPeriodoRequest) (*models.ResultOperacion, error) {
	periodoID, rowAffected, err := serv.repo.RegistrarPeriodo(params)
	if err != nil {
		return nil, err
	}
	resultSmg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente el periodo ", params.Nombre, " con el id ", periodoID),
		Codigo:      periodoID,
		RowAffected: rowAffected,
	}
	return resultSmg, err
}

func (serv *service) ObtenerPeriodoPorID(param *getPeridioByIDRequest) (*Periodo, error) {
	return serv.repo.ObtenerPeriodoPorID(param)
}

func (serv *service) ActualizarPeriodo(params *updatePeriodoRequest) (*models.ResultOperacion, error) {
	rowAffected, err := serv.repo.ActualizarPeriodo(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente ", params.Nombre, " con el id ", params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerTodosLosPeriodos() ([]*Periodo, error) {
	return serv.repo.ObtenerTodosLosPeriodos()
}

func (serv *service) EliminarPeriodoPorID(param *deletePeriodoByIDRequest) (*models.ResultOperacion, error) {
	rowAffected, err := serv.repo.EliminarPeriodoPorID(param)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se elimino correctamente el periodo con el id ", param.ID),
		Codigo:      param.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}
