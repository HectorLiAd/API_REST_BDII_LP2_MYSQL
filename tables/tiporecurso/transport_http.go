package tiporecurso

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Agregar el tipo recurso
	addTipoRecursoHandler := kithttp.NewServer(
		makeAddTipoRecursoEndPoint(s),
		addTipoRecursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addTipoRecursoHandler)

	//Obtener todos los tipos de recursos
	getAllTipoRecursoHandler := kithttp.NewServer(
		makeGetAllTipoRecursoEndPoint(s),
		getAllTipoRecursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allTipoRecurso", getAllTipoRecursoHandler)

	//Obtener tipo de recurso por id
	getTipoRecursoByIDHandler := kithttp.NewServer(
		makeGetTipoRecursoByIDEndPoint(s),
		getTipoRecursoByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getTipoRecursoByIDHandler)

	//Actualizar el tipo de recurso
	updateTipoRecursoByIDHandler := kithttp.NewServer(
		makeUpdateTipoRecursoByIDEndPoint(s),
		updateTipoRecursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updateTipoRecursoByIDHandler)

	return r
}

func addTipoRecursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addTipoRecursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllTipoRecursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func getTipoRecursoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	tipoRecursoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	id := getTipoRecursoByIDRequest{
		ID: tipoRecursoID,
	}
	return id, err
}

func updateTipoRecursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateTipoRecursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
