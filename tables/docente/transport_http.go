package docente

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

	// Registrar Docente
	addDocenteHandler := kithttp.NewServer(
		makeAddDocenteEndPoint(s),
		addDocenteRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addDocenteHandler)

	// Obtener docente por id
	getDocenteByIDHandler := kithttp.NewServer(
		makeGetDocenteByIDEndPoint(s),
		getDocenteByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getDocenteByIDHandler)

	return r
}

func addDocenteRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := docenteRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getDocenteByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	docenteID, err := strconv.Atoi(chi.URLParam(r, "id"))
	docenteReq := docenteRequest{
		ID: docenteID,
	}
	return docenteReq, err
}
