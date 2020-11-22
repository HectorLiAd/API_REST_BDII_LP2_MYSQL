package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*getPersonByIdRequest estructura para recueperar datos del request*/
type getPersonByIDRequest struct {
	PersonaID int
}

/*getPersonsRequest para Obtener datos del body row*/
type getPersonsRequest struct {
	Limit  int //CUANTOS REGISTROS TRAER
	Offset int //DE QUE NUMERO DE FILA INICIARA LA CONSULTA
}

/*addPersonRequest para crear al nuevo usuario*/
type addPersonRequest struct { //
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	Genero          string
	Dni             string
	FechaNacimiento string
}

func makeGetPersonByIDEndPoint(s Service) endpoint.Endpoint {
	getPersonByID := func(ctx context.Context, request interface{}) (interface{}, error) {
		rep := request.(getPersonByIDRequest)
		persona, err := s.GetPersonByID(&rep)
		return persona, err
	}
	return getPersonByID
}

func makeGetPersonsEndPoint(s Service) endpoint.Endpoint {
	getPersonsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPersonsRequest) //Conversion del request al tipo getPersonsRequest
		result, err := s.GetPersons(&req)
		return result, err
	}
	return getPersonsEndPoint
}

func makeAddPersonEndpoint(s Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPersonRequest)
		addPerson, err := s.InsertPerson(&req)
		return addPerson, err
	}
	return addPersonEndpoint
}
