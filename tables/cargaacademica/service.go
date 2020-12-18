package cargaacademica

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarCargaAcad(params *addCargaAcademica) (*models.ResultOperacion, error)
	ObtenerCargaAcadPorID(param *getCargaAcadByID) (*CargaAcademica, error)
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

func (serv *service) RegistrarCargaAcad(params *addCargaAcademica) (*models.ResultOperacion, error) {
	fmt.Println(params)
	cargaAcadID, rowAffected, err := serv.repo.RegistrarCargaAcad(params)
	if err != nil {
		return nil, err
	}
	// Verificar de que es docente o administrador

	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro la carga academica correctamente en la BD con el ID ", cargaAcadID),
		Codigo:      cargaAcadID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerCargaAcadPorID(param *getCargaAcadByID) (*CargaAcademica, error) {
	return serv.repo.ObtenerCargaAcadPorID(param)
}
