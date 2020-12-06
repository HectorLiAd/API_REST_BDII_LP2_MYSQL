package rolusuario

import (
	"net/http"

	"github.com/go-chi/chi"
	// kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Obtener personas por su identificador
	// getPersonByHandler := kithttp.NewServer(
	// 	makeGetPersonByIDEndPoint(s),
	// 	getPersonByIDRequestDecoder,
	// 	kithttp.EncodeJSONResponse,
	// )
	// r.Method(http.MethodGet, "/{id}", getPersonByHandler)

	return r
}
