package curso

import (
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	RegistrarCurso(params *addCursoRequest) (*models.ResultOperacion, error)
	ObtenerCursoPorID(param *getCursoByIDRequest) (*Curso, error)
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

func (serv *service) RegistrarCurso(params *addCursoRequest) (*models.ResultOperacion, error) {
	cursoID, rowAffected, err := serv.repo.RegistrarCurso(params)
	if err != nil {
		return nil, err
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente el curso ", params.Nombre, " con el id ", cursoID),
		Codigo:      cursoID,
		RowAffected: rowAffected,
	}
	return resultMsg, err
}

func (serv *service) ObtenerCursoPorID(param *getCursoByIDRequest) (*Curso, error) {
	return serv.repo.ObtenerCursoPorID(param)
}
