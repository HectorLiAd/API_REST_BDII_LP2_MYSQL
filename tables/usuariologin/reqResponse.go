package usuariologin

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	// "time"
)

/*EncodeJSONResponseFileImgUpload sirve para poderlo usar de forma general*/
func encodeJSONResponseLogin(_ context.Context, w http.ResponseWriter, request interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//Subiendo el token a la Cookie
	expirationTime := time.Now().Add(time.Hour * 24) //
	cookie := http.Cookie{Name: "token", Value: request.(*RespuestaLogin).Token, Expires: expirationTime}
	http.SetCookie(w, &cookie)
	return json.NewEncoder(w).Encode(request)
}
