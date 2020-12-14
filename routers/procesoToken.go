package routers

import (
	"errors"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/usuario"
	"github.com/API_REST_BDII_LP2_MYSQL/userlogin"
	jwt "github.com/dgrijalva/jwt-go"
)

func limpiarDatosUserLogin() {
	userlogin.UsuarioID = 0
	userlogin.UsuarioCorreo = ""
}

/*ProcesoToken Proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	limpiarDatosUserLogin()
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return nil, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	claims, isValidTkn, err := getClaimsToken(tk)
	if isValidTkn {
		databaseConnection := database.InitDB()
		defer databaseConnection.Close()
		var repository = usuario.NewRepository(databaseConnection)
		encontrado, errr := repository.ChequeoEmailExisteUsuario(claims.Email)
		cantPersons, estadoPerson, errr := repository.BuscarPersona(claims.ID)
		if cantPersons != 1 {
			return nil, false, "", errors.New("El usuario al cual desea acceder no existe")
		}
		if estadoPerson == 0 {
			return nil, false, "", errors.New("El usuario esta eliminado temporalmente")
		}
		if errr != nil {
			return claims, false, "", errr
		}
		if encontrado == 1 {
			userlogin.UsuarioID = claims.ID
			userlogin.UsuarioCorreo = claims.Email
			return claims, true, "HOOLA", nil
		}
		return claims, false, "HOOLA", nil
	}

	return nil, false, string(""), err
}

func getClaimsToken(tokenString string) (*models.Claim, bool, error) {
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_token_e_learning"), nil
	})
	if token.Valid {
		return claims, true, err
	}
	return claims, false, err
}
