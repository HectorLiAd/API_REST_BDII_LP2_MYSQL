package routers

import (
	"errors"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"github.com/API_REST_BDII_LP2_MYSQL/usuario"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email Valor usado de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken Proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("XDXDXD_token_XDXDXD")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	databaseConnection := database.InitDB()
	defer databaseConnection.Close()
	var repository = usuario.NewRepository(databaseConnection)

	if err == nil {
		var encontradoBool bool = false

		encontrado, errr := repository.ChequeoEmailExisteUsuario(claims.Email)
		if errr != nil {
			return claims, false, string(""), errr
		}
		if encontrado == 1 {
			Email = claims.Email
			IDUsuario = claims.ID
			encontradoBool = true
		}
		return claims, encontradoBool, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
