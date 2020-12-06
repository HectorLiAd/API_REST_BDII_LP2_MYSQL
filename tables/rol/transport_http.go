package rol

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

	//Agregar rol
	addRolHandler := kithttp.NewServer(
		makeAddRolEndPoint(s),
		addRolRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/", addRolHandler)

	return r
}

func addRolRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addRolRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
