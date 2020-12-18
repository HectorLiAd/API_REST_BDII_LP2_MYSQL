package cargaacademica

import (
	"net/http"

	"github.com/go-chi/chi"
	// 	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Registrar alumno a la bd
	// addAlumnoHandler := kithttp.NewServer(
	// 	makeAddAlumnoEndPoint(s),
	// 	addAlumnoRequestDecoder,
	// 	kithttp.EncodeJSONResponse,
	// )
	// r.Method(http.MethodPost, "/registrar", addAlumnoHandler)

	return r
}

// func addAlumnoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
// 	request := addAlumnoRequest{}
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	return request, err
// }

// func getAlumnoByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
// 	alumnoID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	rol := getAlumnoByIDRequest{
// 		ID: alumnoID,
// 	}
// 	return rol, err
// }
