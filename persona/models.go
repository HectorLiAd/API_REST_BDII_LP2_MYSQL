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
}

/*PersonList mostrar la lista de personas*/
type PersonList struct {
	Data         []*Person `json:"data"`
	TotalRecords int       `json:"totalRecords"`
}

/*StatusPerson muestra la fila que se inserto */
type StatusPerson struct {
	PersonaID string `json:"PersonaID"`
}
