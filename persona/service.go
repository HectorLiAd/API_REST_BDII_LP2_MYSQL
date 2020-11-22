package persona

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
	GetPersonByID(param *getPersonByIDRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) (*PersonList, error)
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
