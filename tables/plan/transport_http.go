package plan

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

	// Registrar plan a la BD
	addPlanHandler := kithttp.NewServer(
		makeAddPlanEndPoint(s),
		addPlanRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addPlanHandler)

	// Obtener el  plan por el ID
	getPlanHandler := kithttp.NewServer(
		makeGetPlanEndPoint(s),
		getPlanByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getPlanHandler)

	// Actualizar el  plan por el ID
	updatePlanByIDHandler := kithttp.NewServer(
		makeUpdatePlanEndPoint(s),
		updatePlanRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updatePlanByIDHandler)

	// Obtener todos los planes
	getAllPlanHandler := kithttp.NewServer(
		makeGetAllPlanEndPoint(s),
		getAllPlanRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allPlan", getAllPlanHandler)

	return r
}

func addPlanRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addPlanRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getPlanByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	rolID, err := strconv.Atoi(chi.URLParam(r, "id"))
	req := getPlanByIDRequest{
		ID: rolID,
	}
	return req, err
}

func updatePlanRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePlanRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllPlanRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
