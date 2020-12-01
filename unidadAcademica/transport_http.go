package unidadacademica

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler permite usar los metodos de la unidad academica*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	// Agregar  una unidad academica
	addUnidadAcademicaHandler := kithttp.NewServer(
		makeAddUnidadAcademicaEndPoint(s),
		addUnidadAcademicaRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addUnidadAcademicaHandler)

	// Buscar unidad academica
	getUnidadAcademicaByIDHandler := kithttp.NewServer(
		makeGetUnidadAcademicaByIDEndPoint(s),
		getUnidadAcademicaByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getUnidadAcademicaByIDHandler)

	return r
}

func addUnidadAcademicaRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addUnidadAcademicaRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getUnidadAcademicaByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	paramID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := idUnidadAcademicaRequest{
		ID: paramID,
	}
	return request, err
}
