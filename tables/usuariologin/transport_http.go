package usuariologin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler para poder usar los diferentes metodos de inciicio secion*/
func MakeHTTPSHandler(s Service) http.Handler {
	//Creacion de las rutas
	r := chi.NewRouter() //Creando instancia para iniciar el ruteo
	//Login user
	loginUseHandler := kithttp.NewServer(
		loginUserEndPoint(s),
		loginUserRequestDecoder,
		encodeJSONResponseLogin,
		// kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/login", loginUseHandler)

	//ResetPassword user
	passwordResetdUserHandler := kithttp.NewServer(
		passwordResetEndPoint(s),
		passwordResetRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/password_reset", passwordResetdUserHandler)

	return r
}

func loginUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := loginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func passwordResetRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := passwordResetRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
