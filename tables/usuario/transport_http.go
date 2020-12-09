package usuario

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	// "github.com/API_REST_BDII_LP2_MYSQL/routers"
	"github.com/API_REST_BDII_LP2_MYSQL/userlogin"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler sirve para hacer peticiones de metodos*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	registerUserHandler := kithttp.NewServer(
		makeRegisterUserEndPoint(s),
		registerUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/register", registerUserHandler)

	updateAvatarUserHandler := kithttp.NewServer(
		makeSubirImagenUserEndPoint(s),
		updateImagenRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/subirImagen", updateAvatarUserHandler)

	return r
}

func registerUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func updateImagenRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var imgPath = fmt.Sprint("uploads/avatars/", userlogin.UsuarioID, ".", extension)
	var archivo string = imgPath

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(f, file)
	var rutaImgBD string = fmt.Sprint(userlogin.UsuarioID, ".", extension)

	return subirAvartarRequest{File: rutaImgBD}, nil
}
