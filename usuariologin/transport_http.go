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
	loginUseHandler := kithttp.NewServer(
		loginUserEndPoint(s),
		loginUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/login", loginUseHandler)

	return r
}

func loginUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := loginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}
