package routers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/API_REST_BDII_LP2_MYSQL/database"
	"github.com/API_REST_BDII_LP2_MYSQL/models"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/usuario"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email Valor usado de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken Proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return nil, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	claims, isValidTkn, err := getClaimsToken(tk)

	if isValidTkn {
		encontradoBool := false

		databaseConnection := database.InitDB()
		defer databaseConnection.Close()
		var repository = usuario.NewRepository(databaseConnection)
		encontrado, errr := repository.ChequeoEmailExisteUsuario(claims.Email)
		if errr != nil {
			return claims, false, string(""), errr
		}
		if encontrado == 1 {
			encontradoBool = true
			// fmt.Println("El usuario ingresado si existe")
		}
		return claims, encontradoBool, "HOOLA", nil
	}

	return claims, false, string(""), err
}

func getClaimsToken(tokenString string) (*models.Claim, bool, error) {
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_token_e_learning"), nil
	})
	/*QUIERO MOSTRAR LOS DATOS SI SON VALIDOS
	Y DESPUES ENVIARLO A LA BASE DE DATOS PARA COMPROBAR
	Y LUEGO QUIERO QUE RETORNE MI MODELO CLAIMS*/
	// fmt.Println(claims.Email) POSIBLEMENTE PUEDA QUE NECESITE PARA VER QUE DATOS ME LLEGA DEL TOKEN
	// fmt.Println(claims.ID) POSIBLEMENTE PUEDA QUE NECESITE PARA VER QUE DATOS ME LLEGA DEL TOKEN
	// fmt.Println(claims.Name) POSIBLEMENTE PUEDA QUE NECESITE PARA VER QUE DATOS ME LLEGA DEL TOKEN
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	if token.Valid {
		fmt.Println(claims)
		// fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		return claims, true, err
	}
	return claims, false, err
}
