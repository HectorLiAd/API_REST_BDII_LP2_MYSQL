package usuariologin

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	// "time"
	"fmt"
)

/*EncodeJSONResponseFileImgUpload sirve para poderlo usar de forma general*/
func encodeJSONResponseLogin(_ context.Context, w http.ResponseWriter, request interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	req := request.(*RespuestaLogin)
	expirationTime := time.Now().Add(24 * time.Hour) //
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   req.Token,
		Expires: expirationTime,
	})
	fmt.Println(req)
	return json.NewEncoder(w).Encode(request)
}
