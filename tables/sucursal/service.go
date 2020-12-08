package sucursal

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	InsertarSucursal(params *addSucursalRequest) (*models.ResultOperacion, error)
	ActualizarSucursal(params *updateSucursalRequest) (*models.ResultOperacion, error)
	ObtenerTodoSucursal() ([]*Sucursal, error)
	ObtenerSucursalPorID(param *getSucursalByIDRequest) (*Sucursal, error)
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

func (s *service) InsertarSucursal(params *addSucursalRequest) (*models.ResultOperacion, error) {
	resultID, rowAffected, err := s.repo.InsertarSucursal(params)
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente ", params.Nombre, " con el id ", resultID),
		Codigo:      resultID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (s *service) ObtenerTodoSucursal() ([]*Sucursal, error) {
	return s.repo.ObtenerTodoSucursal()
}

func (s *service) ActualizarSucursal(params *updateSucursalRequest) (*models.ResultOperacion, error) {
	rowAffected, err := s.repo.ActualizarSucursal(params)
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente la sucursal con el id ", params.ID),
		Codigo:      params.ID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (s *service) ObtenerSucursalPorID(param *getSucursalByIDRequest) (*Sucursal, error) {
	return s.repo.ObtenerSucursalPorID(param)
}
