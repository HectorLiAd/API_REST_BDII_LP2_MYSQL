package persona

/*Person sirve para mostrar en el response*/
type Person struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre_personal"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Genero          string `json:"Genero"`
	Dni             string `json:"dni"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Estado          int    `json:"estado"`
}
