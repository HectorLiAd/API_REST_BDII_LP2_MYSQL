package alumno

import (
	"github.com/API_REST_BDII_LP2_MYSQL/tables/persona"
)

/*Alumno permite que se muestre en formato json*/
type Alumno struct {
	ID      int             `json:"id"`
	Persona *persona.Person `json:"persona"`
}

/*Persona se mostrara en json*/
// type Persona struct {
// 	ID              int       `json:"id"`
// 	Nombre          string    `json:"nombre"`
// 	ApellidoPaterno string    `json:"apellidoPaterno"`
// 	ApellidoMaterno string    `json:"apellidoMaterno"`
// 	Genero          string    `json:"genero"`
// 	DNI             string    `json:"DNI"`
// 	FechaNacimiento time.Time `json:"fechaDeNacimiento"`
// }
