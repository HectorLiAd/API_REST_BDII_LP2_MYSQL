package sucursal

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

	//Insertar sucursal
	addSucursalHandler := kithttp.NewServer(
		makeAddSucursalEndPoint(s),
		addSurcursalRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addSucursalHandler)

	//Obtener todas las sucursales
	getAllSucursalHandler := kithttp.NewServer(
		makeGetAllSucursalEndPoint(s),
		getAllSurcursalRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allSucursal", getAllSucursalHandler)

	//Actualizar sucursal
	updateSucursalHandler := kithttp.NewServer(
		makeUpdateSucursalEndPoint(s),
		updateSurcursalRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updateSucursalHandler)

	//Obtener sucursal por id
	getSucursaByIDlHandler := kithttp.NewServer(
		makeGetSucursalByIDEndPoint(s),
		getSurcursalByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getSucursaByIDlHandler)

	return r
}

func addSurcursalRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addSucursalRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllSurcursalRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func updateSurcursalRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateSucursalRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getSurcursalByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	sucursalID, err := strconv.Atoi(chi.URLParam(r, "id"))
	rol := getSucursalByIDRequest{
		ID: sucursalID,
	}
	return rol, err
}
