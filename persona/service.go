package persona

import (
	"errors"
	"fmt"
	"strconv"
)

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	GetPersonByID(param *getPersonByIDRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
	InsertPerson(params *addPersonRequest) (*StatusPerson, error)
	UpdatePerson(params *updatePersonRequest) (*StatusPerson, error)
	DeletePerson(param *deletePersonRequest) (*StatusPerson, error)
}

type service struct {
	repo Repository
}

/*NerService Permite crear un nuevo servicio teniendo una nuevo el repositorio*/
func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetPersonByID(param *getPersonByIDRequest) (*Person, error) {
	persona, err := s.repo.GetPersonByID(param)
	return persona, err
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

func (s *service) InsertPerson(params *addPersonRequest) (*StatusPerson, error) {
	personaID, err := s.repo.InsertPerson(params)

	if err != nil {
		return nil, err
	}

	if personaID <= 0 {
		return nil, errors.New("no se pudo registrar a la persona")
	}

	estadoInsert := StatusPerson{
		PersonaID: "Cod " + strconv.Itoa(personaID) + " registrado corractamente",
	}
	return &estadoInsert, err
}

func (s *service) UpdatePerson(params *updatePersonRequest) (*StatusPerson, error) {
	rowAfected, err := s.repo.UpdatePerson(params)

	fmt.Println(params.ID)
	if err != nil {
		return nil, err
	}

	if rowAfected == 0 {
		return nil, errors.New("No se pudo actualizar posiblemente el usuario no exista o los datos no se alteraron")
	}

	estadoInsert := StatusPerson{
		PersonaID: "Cod " + strconv.Itoa(params.ID) + " se actualizo corractamente " + strconv.Itoa(rowAfected) + " usuario",
	}
	return &estadoInsert, err
}

func (s *service) DeletePerson(param *deletePersonRequest) (*StatusPerson, error) {
	rowAfected, err := s.repo.DeletePerson(param)
	if err != nil {
		return nil, err
	}

	if rowAfected == 0 {
		return nil, errors.New("No se pudo eliminar posiblemente el usuario no exista")
	}

	estadoInsert := StatusPerson{
		PersonaID: "Cod " + strconv.Itoa(param.PersonaID) + " se elimino corractamente " + strconv.Itoa(rowAfected) + " usuario",
	}
	return &estadoInsert, err
}
