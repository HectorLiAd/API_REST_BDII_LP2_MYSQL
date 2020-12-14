package jerarquia

import (
	"context"
	"encoding/json"
	"net/http"

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

// func getSurcursalByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
// 	sucursalID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	rol := getSucursalByIDRequest{
// 		ID: sucursalID,
// 	}
// 	return rol, err
// }
