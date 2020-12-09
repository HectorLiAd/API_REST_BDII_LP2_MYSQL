package alumno

import (
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	AgregarPersonaAlumno(param *addAlumnoRequest) (*models.ResultOperacion, error)
	ObtenerPersonaAlumnoPorID(param *getAlumnoByIDRequest) (*Alumno, error)
	ObtenerTodasLasPersonas() ([]*Alumno, error)
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

func (s *service) AgregarPersonaAlumno(param *addAlumnoRequest) (*models.ResultOperacion, error) {
	persona, err := s.repo.BuscarPersonaPorID(param.ID)
	if persona == nil {
		return nil, errors.New(fmt.Sprint("La persona no existe o posiblemente este eliminado temporalmente ", err))
	}
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al buscar en la bd ", err))
	}
	alumnoID, rowAffected, err := s.repo.AgregarPersonaAlumno(param)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al registrar al alumno ", err))
	}
	resultMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente el alumno con el id ", alumnoID),
		Codigo:      alumnoID,
		RowAffected: rowAffected,
	}
	return resultMsg, nil
}

func (s *service) ObtenerPersonaAlumnoPorID(param *getAlumnoByIDRequest) (*Alumno, error) {
	cant, err := s.repo.ExistePersonaAlumnoPorID(param)
	if cant != 1 {
		return nil, errors.New(fmt.Sprint("El alumno la cual quiere buscar no existe"))
	}
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al obtener al alumno de la BD ", err))
	}
	persona, err := s.repo.BuscarPersonaPorID(param.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al obtener el alumno de la BD ", err))
	}
	alumno := &Alumno{
		ID:      param.ID,
		Persona: persona,
	}
	return alumno, nil
}

func (s *service) ObtenerTodasLasPersonas() ([]*Alumno, error) {
	return s.repo.ObtenerTodoPersonaAlumno()
}
