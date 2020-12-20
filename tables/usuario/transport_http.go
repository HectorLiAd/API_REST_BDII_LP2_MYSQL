package usuario

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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
	r.Method(http.MethodPut, "/subirAvatar", updateAvatarUserHandler)

	// Obtener avatar
	getAvatarUserHandler := kithttp.NewServer(
		makeGetImagenUserEndPoint(s),
		getAvatarUserRequestDecoder,
		EncodeJSONResponseFileImgUpload,
	)
	r.Method(http.MethodGet, "/obtenerImagen/{id}", getAvatarUserHandler)

	// Obtener a todos los usuarios
	getAllUserHandler := kithttp.NewServer(
		makeGetAllUserEndPoint(s),
		getAllUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/allUsuario", getAllUserHandler)

	return r
}

func registerUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := registerUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func updateImagenRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	file, _ /*handler*/, err := r.FormFile("avatar")
	// var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = fmt.Sprint("uploads/avatars/", userlogin.UsuarioID, ".jpg")
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(f, file)
	defer f.Close()
	return subirAvartarRequest{File: fmt.Sprint(userlogin.UsuarioID, ".jpg")}, err
}

func getAvatarUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	usuarioID, err := strconv.Atoi(chi.URLParam(r, "id"))
	return obtenerAvatarRequest{
		ID: usuarioID,
	}, err
}

func getAllUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
