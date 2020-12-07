package rolusuario

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

	//Obtener personas por su identificador
	addRolUsuarioHandler := kithttp.NewServer(
		makeAddRolUsuarioEndpoint(s),
		addRolUsuarioRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addRolUsuarioHandler)

	//Obtener rol usario por su identificador
	getRolUsuarioByIDHandler := kithttp.NewServer(
		makeGetRolUsuarioByIDEndpoint(s),
		getRolUsuarioByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getRolUsuarioByIDHandler)

	return r
}

func addRolUsuarioRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addRolUsuarioRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getRolUsuarioByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	rolUsuarioID, err := strconv.Atoi(chi.URLParam(r, "id"))
	rol := getRolUsuarioByIDRequest{
		ID: rolUsuarioID,
	}
	return rol, err
}