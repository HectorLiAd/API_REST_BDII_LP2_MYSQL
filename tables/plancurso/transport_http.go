package plancurso

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

	//Registrar alumno a la bd
	addPlanCursoHandler := kithttp.NewServer(
		makeAddPlanCursoEndPoint(s),
		addPlanCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addPlanCursoHandler)

	//Obtener el plan curso por el ID
	getPlanCursoByIDHandler := kithttp.NewServer(
		makeGetPlanCursoByIDEndPoint(s),
		getPlanCursoByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getPlanCursoByIDHandler)

	//Actualizar el plan cursp  por el ID
	updatePlanCursoByIDHandler := kithttp.NewServer(
		makeUpdatePlanCursoByIDEndPoint(s),
		updatePlanCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updatePlanCursoByIDHandler)

	//Obtener todos los planes de los  cursos
	getAllPlanCursoHandler := kithttp.NewServer(
		makeGetAllPlanCursoEndPoint(s),
		getAllPlanCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allPlanCurso", getAllPlanCursoHandler)

	return r
}

func addPlanCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addPlanCursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getPlanCursoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	planCursoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	rol := getPlanCursoByIDRequest{
		ID: planCursoID,
	}
	return rol, err
}

func updatePlanCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updatePlanCursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllPlanCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
