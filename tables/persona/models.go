package persona

/*Person sirve para mostrar en el response*/
type Person struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellidoPaterno"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	Genero          string `json:"genero"`
	DNI             string `json:"DNI"`
	FechaNacimiento string `json:"fechaDeNacimiento"`
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
