package tipounidad

import (
	"context"
	"encoding/json"
	"net/http"

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

	//Agregar a el tipo unidad
	getAllTipoUnidadHandler := kithttp.NewServer(
		makeGetAllTipoUnidadEndPoint(s),
		getAllTipoUnidadRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/", getAllTipoUnidadHandler)

	return r
}

func addTipoUnidadRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addTipoUnidadRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
func getAllTipoUnidadRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return addTipoUnidadRequest{}, nil
}
