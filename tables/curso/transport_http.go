package curso

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

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

	// Subir Imagen todos los cursos
	uploadFondoCursoByIDHandler := kithttp.NewServer(
		makeUploadImageCursoEndPoint(s),
		uploadFondoCursoRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPut, "/fondo/id/{id}", uploadFondoCursoByIDHandler)

	// Subir Imagen todos los cursos
	getFondoCursoByIDHandler := kithttp.NewServer(
		makeGetFondoCursoEndPoint(s),
		getFondoCursoRequestDecoder,
		EncodeJSONResponseFileImgUpload,
	)
	r.Method(http.MethodGet, "/fondo/id/{id}", getFondoCursoByIDHandler)

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

func uploadFondoCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	cursoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}
	file, handler, err := r.FormFile("fondo")
	var extension = strings.Split(handler.Filename, ".")
	archivo := fmt.Sprint("uploads/curso/fondos/", cursoID, ".", extension[len(extension)-1])
	fmt.Println(archivo)
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(f, file)
	defer f.Close()
	fmt.Println("Correct")
	return updateImagenCursoByIDRequest{
		ID:   cursoID,
		File: fmt.Sprint(cursoID, ".", extension[len(extension)-1]),
	}, err
}

func getFondoCursoRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	cursoID, err := strconv.Atoi(chi.URLParam(r, "id"))
	req := getFondoCursoByIDRequest{
		ID: cursoID,
	}
	return req, err
}
