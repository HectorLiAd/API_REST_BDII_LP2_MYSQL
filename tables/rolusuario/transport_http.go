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

	//Agregar rol usuario
	addRolUsuarioHandler := kithttp.NewServer(
		makeAddRolUsuarioEndpoint(s),
		addRolUsuarioRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addRolUsuarioHandler)

	//Obtener rol usario por su identificador
	getRolUsuarioByIDHandler := kithttp.NewServer(
		makeGetRolUsuarioByIDEndpoint(s),
		getRolUsuarioByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getRolUsuarioByIDHandler)

	//Obtener todos los roles de ususario
	getAllRolUsuarioHandler := kithttp.NewServer(
		makeGetAllRolUsuarioEndpoint(s),
		getAllRolUsuarioRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allRolUsuario", getAllRolUsuarioHandler)

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

func getAllRolUsuarioRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
