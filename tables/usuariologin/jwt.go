package usuariologin

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT Permite Crear el JWT*/
func GeneroJWT(t *Usuario) (string, error) {
	miClave := []byte("secret_token_e_learning") //Creando clave privada
	claims := jwt.MapClaims{}
	claims["email"] = t.UsuarioEmail
	claims["nombre"] = t.UsuarioNombre
	claims["rol"] = t.Rol
	claims["_id"] = t.UsuarioID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims) //El goritmo para encriptar "header"
	tokenStr, err := token.SignedString(miClave)                //Firmando con el string de firma
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
