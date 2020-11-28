package usuario

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler sirve para hacer peticiones de metodos*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	registerUserHandler := kithttp.NewServer(
		registerUserEndPoint(s),
		registerUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/register", registerUserHandler)

	return r
}

func registerUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
