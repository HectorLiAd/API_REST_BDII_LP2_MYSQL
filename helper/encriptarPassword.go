package helper

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword permite cifrar el pwd*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //La eficiencia se relaciona con el costo pero toma tiempo y demora desencriptar
	byte, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(byte), err
}
