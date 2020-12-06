package tiporecurso

/*Service interface para poder usarlo de forma nativa desde el main*/
type Service interface {
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
