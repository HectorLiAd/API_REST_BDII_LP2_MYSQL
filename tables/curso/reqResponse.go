package curso

import (
	"context"
	"io"
	"net/http"
	"os"
)

/*EncodeJSONResponseFileImgUpload sirve para poderlo usar de forma general*/
func EncodeJSONResponseFileImgUpload(_ context.Context, w http.ResponseWriter, request interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	req := request.(*getFondoCursoByIDRequest)
	openFile, err := os.Open("uploads/curso/fondos/" + req.File)
	if err != nil {
		http.Error(w, "Error al abrir la imagen", http.StatusBadRequest)
	}
	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
	// return json.NewEncoder(w).Encode(map[string]string{"mensaje": "Imagen Encontrada"})
	return nil
}
