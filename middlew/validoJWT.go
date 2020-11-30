package middlew

import (
	"net/http"

	"github.com/API_REST_BDII_LP2_MYSQL/routers"
)

/*ValidoJWT valida el TOKEN*/
func ValidoJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token pipipipipi "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*Validar de otra forma el JWT*/
