package persona

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Obtener personas por su identificador
	getPersonByHandler := kithttp.NewServer(
		makeGetPersonByIDEndPoint(s),
		getPersonByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/{id}", getPersonByHandler)

	return r
}

func getPersonByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	personaID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getPersonByIDRequest{
		PersonaID: personaID,
	}, nil
}
