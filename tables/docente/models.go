package docente

import "github.com/API_REST_BDII_LP2_MYSQL/tables/persona"

/*Docente nos permite usar para comvertirlo en JSON*/
type Docente struct {
	ID      int             `json:"id"`
	Persona *persona.Person `json:"persona"`
}
