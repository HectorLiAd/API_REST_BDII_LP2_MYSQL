package unidadacademica

/*UnidadAcademica obtener la unidad academica*/
type UnidadAcademica struct {
	ID         int    `json:"id"`
	TipoUnidad string `json:"TipoUnidad,omitempty"`
	Nombre     string `json:"nombre"`
}
