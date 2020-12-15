package curso

/*Curso sirve para hacer el mapeo a JSON*/
type Curso struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
