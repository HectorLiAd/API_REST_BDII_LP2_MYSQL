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
	Nombre      string
	ApellidoPat string
	ApellidoMat string
	Genero      string
	DNI         string
	FechaNac    string
}

/*updatePersonRequest para poder actualizar a la persona*/
type updatePersonRequest struct {
	ID          int
	Nombre      string
	ApellidoPat string
	ApellidoMat string
	Genero      string
	DNI         string
	FechaNac    string
}

/*deletePersonRequest para obtener el id del body row y proceder a eliminar*/
type deletePersonRequest struct {
	PersonaID int
}

/*getPersonByDNIRequest el json se convierte en esta estructura para poder usarlo en el service y repository*/
type getPersonByDNIRequest struct {
	DNI string
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
		addPerson, err := s.InsertPerson(trimStrAddPersonRequest(&req))
		return addPerson, err
	}
	return addPersonEndpoint
}

func makeUpdatePersonEndpoint(s Service) endpoint.Endpoint {
	updatePersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updatePersonRequest)
		r, err := s.UpdatePerson(trimStrUpdatePersonRequest(&req))
		return r, err
	}
	return updatePersonEndpoint
}

func makeDeletePersonEndPoint(s Service) endpoint.Endpoint {
	deletePersonEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deletePersonRequest)
		result, err := s.DeletePerson(&req)
		return result, err
	}
	return deletePersonEndPoint
}
