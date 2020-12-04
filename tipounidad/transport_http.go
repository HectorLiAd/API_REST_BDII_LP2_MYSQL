package tipounidad

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos sirva para manipular los metodos del tipo unidad*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Agregar a el tipo unidad
	addTipoUnidadHandler := kithttp.NewServer(
		makeAddTipoUnidadEndPoint(s),
		addTipoUnidadRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addTipoUnidadHandler)

	//Agregar a todos los tipos de unidad
	getAllTipoUnidadHandler := kithttp.NewServer(
		makeGetAllTipoUnidadEndPoint(s),
		getAllTipoUnidadRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/", getAllTipoUnidadHandler)

	//Agregar tipo de unidad por id
	getTipoUnidadByIDHandler := kithttp.NewServer(
		makeGetTipoUnidadByIDEndPoint(s),
		getTipoUnidadByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getTipoUnidadByIDHandler)

	return r
}

func addTipoUnidadRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addTipoUnidadRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
func getAllTipoUnidadRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func getTipoUnidadByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Println(id)
	return getTipoUnidadByIDRequest{ID: id}, err
}
