package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*Claim Estructra para procesar el JWT*/
type Claim struct {
	Email string `json:"email"`
	ID    int    `json:"_id"`
	Name  string `json:"nombre"`
	jwt.StandardClaims
}
