package jerarquia

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

	// Registrar jerarquia
	addJerarquiaHandler := kithttp.NewServer(
		makeAddJerarquiaEndPoint(s),
		addJerarquiaRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addJerarquiaHandler)

	// Agregar una jerarquia padre
	addJerarquiaParentByIDHandler := kithttp.NewServer(
		makeAddJerarquiaParentByIDEndPoint(s),
		addJerarquiaParentByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/agregarJerarquiaPadre", addJerarquiaParentByIDHandler)

	// Buscar jerarquia padre por el ID
	getJerarquiaByIDHandler := kithttp.NewServer(
		makeGetJerarquiaByIDEndPoint(s),
		getJerarquiaByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getJerarquiaByIDHandler)

	// Obtener todas las jerarquias
	getAllJerarquiaHandler := kithttp.NewServer(
		makeGetAllJerarquiaEndPoint(s),
		getAllJerarquiaRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allJerarquia", getAllJerarquiaHandler)

	return r
}

func addJerarquiaRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addJerarquiaRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func addJerarquiaParentByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addJerarquiParentRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getJerarquiaByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	jerarquiaID, err := strconv.Atoi(chi.URLParam(r, "id"))
	rol := getJerarquiaByIDRequest{
		ID: jerarquiaID,
	}
	return rol, err
}

func getAllJerarquiaRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
