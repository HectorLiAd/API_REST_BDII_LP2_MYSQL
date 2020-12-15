package periodo

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

	//Registrar periodo
	addPeriodoHandler := kithttp.NewServer(
		makeAddPeriodoEndPoint(s),
		addPeriodoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addPeriodoHandler)

	//Obtener periodo por el ID
	getPeriodoByIDHandler := kithttp.NewServer(
		makeGetPeriodoByIDEndPoint(s),
		getPeriodoByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getPeriodoByIDHandler)

	//Actualizar periodo por el ID
	updatePeriodoByIDHandler := kithttp.NewServer(
		makeUpdatePeriodoByIDEndPoint(s),
		updatePeriodoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updatePeriodoByIDHandler)

	//Obtener todos los periodos
	getAllPeriodoHandler := kithttp.NewServer(
		makeGetAllPeriodoEndPoint(s),
		getAddPeriodoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allPeriodos", getAllPeriodoHandler)

	//Eliminar el  periodo por el ID
	deletePeriodoByIDHandler := kithttp.NewServer(
		makeDeletePeriodoByIDEndPoint(s),
		deletePeriodoByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodDelete, "/id/{id}", deletePeriodoByIDHandler)

	return r
}

func addPeriodoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addPeriodoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getPeriodoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	periodoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := getPeridioByIDRequest{
		ID: periodoID,
	}
	return request, err
}

func updatePeriodoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePeriodoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAddPeriodoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func deletePeriodoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	periodoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := deletePeriodoByIDRequest{
		ID: periodoID,
	}
	return request, err
}
