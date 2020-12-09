package alumno

import "time"

/*Alumno permite que se muestre en formato json*/
type Alumno struct {
	ID      int      `json:"id"`
	Persona *Persona `json:"persona"`
}

/*Persona se mostrara en json*/
type Persona struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"nombre"`
	ApellidoPaterno string    `json:"apellidoPaterno"`
	ApellidoMaterno string    `json:"apellidoMaterno"`
	Genero          string    `json:"genero"`
	DNI             string    `json:"DNI"`
	FechaNacimiento time.Time `json:"fechaDeNacimiento"`
}
