package cargaacademica

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

	//Registrar carga academica a la bd
	addCargaAcadHandler := kithttp.NewServer(
		makeAddCargaAcadEndPoint(s),
		addCargaAcadRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addCargaAcadHandler)

	//Obtener carga academica por ID
	getCargaAcadByIDHandler := kithttp.NewServer(
		makeGetCargaAcadByIDEndPoint(s),
		getCargaAcadByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getCargaAcadByIDHandler)

	return r
}

func addCargaAcadRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addCargaAcademica{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getCargaAcadByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	cargaAcadID, err := strconv.Atoi(chi.URLParam(r, "id"))
	req := getCargaAcadByID{
		ID: cargaAcadID,
	}
	return req, err
}
