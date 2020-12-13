package unidadacademica

/*UnidadAcademica obtener la unidad academica*/
type UnidadAcademica struct {
	ID         int    `json:"id"`
	TipoUnidad string `json:"ipoUnidad,omitempty"`
	Nombre     string `json:"nombre"`
}
