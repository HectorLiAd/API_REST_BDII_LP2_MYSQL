package rol

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

	//Agregar rol
	addRolHandler := kithttp.NewServer(
		makeAddRolEndPoint(s),
		addRolRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addRolHandler)

	//Actualizar rol
	updateRolHandler := kithttp.NewServer(
		makeUpdateRolEndPoint(s),
		updateRolRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/", updateRolHandler)

	//Obtener rol por id
	getRolByIDHandler := kithttp.NewServer(
		makeGetRolByIDEndPoint(s),
		getRolByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getRolByIDHandler)

	return r
}

func addRolRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addRolRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func updateRolRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateRolRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getRolByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	rol := getRolByIDRequest{
		ID: personaID,
	}
	return rol, err
}
