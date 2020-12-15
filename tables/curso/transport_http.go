package curso

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

	// Registrar curso a la bd
	addCursoHandler := kithttp.NewServer(
		makeAddCursoEndPoint(s),
		addCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/registrar", addCursoHandler)

	// Obtene curso por el ID
	getCursoByIDHandler := kithttp.NewServer(
		makeGetCursoByIDEndPoint(s),
		getCursoByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/id/{id}", getCursoByIDHandler)

	// Actualizar el curso por ID
	updateCursoByIDHandler := kithttp.NewServer(
		makeUpdateCursoByIDEndPoint(s),
		updateCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/actualizar", updateCursoByIDHandler)

	// Obtener todos los cursos
	getAllCursoByIDHandler := kithttp.NewServer(
		makeGetAllCursoEndPoint(s),
		getAllCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allCurso", getAllCursoByIDHandler)

	return r
}

func addCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addCursoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getCursoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	cursoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	req := getCursoByIDRequest{
		ID: cursoID,
	}
	return req, err
}

func updateCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateCursoByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func getAllCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
