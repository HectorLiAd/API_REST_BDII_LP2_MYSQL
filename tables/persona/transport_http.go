package persona

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Obtener personas por su identificador
	getPersonByHandler := kithttp.NewServer(
		makeGetPersonByIDEndPoint(s),
		getPersonByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getPersonByHandler)

	//Obtener personas paginadas
	getPersonHandler := kithttp.NewServer(
		makeGetPersonsEndPoint(s),
		getPersonsRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/paginated", getPersonHandler)

	//Agregar a una persona
	addPersonHandler := kithttp.NewServer(
		makeAddPersonEndpoint(s),
		addPersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/insert", addPersonHandler)

	//Actualizar personas
	updatePersonHandler := kithttp.NewServer(
		makeUpdatePersonEndpoint(s),
		updatePersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/update", updatePersonHandler)

	//Eliminar PERSONA
	deletePersonHandler := kithttp.NewServer(
		makeDeletePersonEndPoint(s),
		deletePersonRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/delete/{id}", deletePersonHandler)
	return r
}

func getPersonByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.New(fmt.Sprint("El tipo ingresado no es valido ", err))
	}
	return getPersonByIDRequest{
		PersonaID: personaID,
	}, nil
}

func getPersonsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getPersonsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func addPersonRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addPersonRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func updatePersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePersonRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func deletePersonRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.New(fmt.Sprint("El tipo ingresado no es valido ", err))
	}
	return deletePersonRequest{
		PersonaID: personaID,
	}, nil
}
