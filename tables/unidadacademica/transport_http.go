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
	r.Method(http.MethodPost, "/registrar", addUnidadAcademicaHandler)

	// Buscar unidad academica por el id
	getUnidadAcademicaByIDHandler := kithttp.NewServer(
		makeGetUnidadAcademicaByIDEndPoint(s),
		getUnidadAcademicaByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getUnidadAcademicaByIDHandler)

	// Actualizar la unidad academica en especifico
	updateUnidadAcademicaByIDHandler := kithttp.NewServer(
		makeUpdateUnidadAcademicaEndPoint(s),
		updateUnidadAcademicaRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updateUnidadAcademicaByIDHandler)

	// Obtener todas las unidades academicas
	getAllUnidadAcademicaByIDHandler := kithttp.NewServer(
		makeGetAllUnidadAcademicaEndPoint(s),
		getAllUnidadAcademicaRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allUnidadAcademica", getAllUnidadAcademicaByIDHandler)

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

func updateUnidadAcademicaRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateUnidadAcademicaRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllUnidadAcademicaRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
