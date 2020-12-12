package persona

import (
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/models"

	"github.com/API_REST_BDII_LP2_MYSQL/helper"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	GetPersonByID(param *getPersonByIDRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
	InsertPerson(params *addPersonRequest) (*models.ResultOperacion, error)
	UpdatePerson(params *updatePersonRequest) (*models.ResultOperacion, error)
	DeletePerson(param *deletePersonRequest) (*models.ResultOperacion, error)
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

func (s *service) GetPersonByID(param *getPersonByIDRequest) (*Person, error) {
	persona, err := s.repo.GetPersonByID(param)
	if err != nil {
		return nil, errors.New(fmt.Sprint("No existe el usarioa o este eliminado temporalmente ", err))
	}
	return persona, nil
}

func (s *service) GetPersons(params *getPersonsRequest) (*PersonList, error) {
	params.Offset--

	persons, err := s.repo.GetPersons(params)
	if err != nil {
		return nil, err
	}
	totalPersons, er := s.repo.GetTotalPersons()
	if er != nil {
		return nil, er
	}
	return &PersonList{
		Data:         persons,
		TotalRecords: totalPersons,
	}, err
}

func (s *service) InsertPerson(params *addPersonRequest) (*models.ResultOperacion, error) {
	//Validacion params
	if !helper.ValidarDniStr(params.DNI) || len(params.DNI) != 8 {
		return nil, errors.New("El DNI ingresado no es un formato valido")
	}
	if !helper.ValidarDateStr(params.FechaNac) {
		return nil, errors.New("Formato de fecha no admitido, intente YYYY/MM/DD o YYYY-MM-DD")
	}
	if len(params.Genero) != 1 {
		return nil, errors.New("El genero solo admite un caractes M/F/O")
	}
	//Ingresando datos a la BD y validaciones
	personaID, rowAffected, err := s.repo.InsertPerson(params)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Posiblemente el usuario al que desea ingresar ya existe ", err))
	}
	if personaID <= 0 {
		return nil, errors.New("no se pudo registrar a la persona")
	}

	resultInsert := &models.ResultOperacion{
		Name:        fmt.Sprint("Se registro correctamente a ", params.Nombre, " con el id ", personaID),
		Codigo:      personaID,
		RowAffected: rowAffected,
	}
	return resultInsert, err
}

func (s *service) UpdatePerson(params *updatePersonRequest) (*models.ResultOperacion, error) {
	//Validacion params
	if !helper.ValidarDniStr(params.DNI) || len(params.DNI) != 8 {
		return nil, errors.New("El DNI ingresado no es un formato valido")
	}
	if !helper.ValidarDateStr(params.FechaNac) {
		return nil, errors.New("Formato de fecha no admitido, intente YYYY/MM/DD o YYYY-MM-DD")
	}
	//Insertando a la BD y validando
	rowAfected, err := s.repo.UpdatePerson(params)
	if err != nil {
		return nil, err
	}
	if rowAfected == 0 {
		return nil, errors.New("No se pudo actualizar posiblemente el usuario no exista o los datos no se alteraron")
	}
	resultUpdateMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se actualizo correctamente a ", params.Nombre),
		Codigo:      params.ID,
		RowAffected: rowAfected,
	}
	return resultUpdateMsg, err
}

func (s *service) DeletePerson(param *deletePersonRequest) (*models.ResultOperacion, error) {
	rowAfected, err := s.repo.DeletePerson(param)
	if err != nil {
		return nil, err
	}
	if rowAfected == 0 {
		return nil, errors.New("No se pudo eliminar posiblemente el usuario no exista")
	}
	resultDeleteMsg := &models.ResultOperacion{
		Name:        fmt.Sprint("Se elimino corractamente al usuario con el id ", rowAfected),
		Codigo:      param.PersonaID,
		RowAffected: rowAfected,
	}
	return resultDeleteMsg, err
}
