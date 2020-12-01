package unidadacademica

/*UnidadAcademica obtener la unidad academica*/
type UnidadAcademica struct {
	ID           int    `json:"id"`
	TipoUnidadID int    `json:"idTipoUnidad"`
	Nombre       string `json:"nombre"`
	// TipoUnidad   tipounidad.TipoUnidad `json:"TipoUnidad,omitempty"`
}
