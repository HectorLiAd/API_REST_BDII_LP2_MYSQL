package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*getPersonByIdRequest estructura para recueperar datos del request*/
type getPersonByIDRequest struct {
	PersonaID int
}

func makeGetPersonByIDEndPoint(s Service) endpoint.Endpoint {
	getPersonByID := func(ctx context.Context, request interface{}) (interface{}, error) {
		rep := request.(getPersonByIDRequest)
		persona, err := s.GetPersonByID(&rep)
		return persona, err
	}
	return getPersonByID
}
