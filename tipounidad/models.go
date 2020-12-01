package tipounidad

/*TipoUnidad permite transformar la informacion de la BD a un json segun sugun la esquema de esta estructura*/
type TipoUnidad struct {
	ID          int                `json:"id"`
	Nombre      string             `json:"nombre,omitempty"`
	Descripcion string             `json:"descripcion,omitempty"`
	UnidadAcad  []*UnidadAcademica `json:"unidadesAcademicas,omitempty"`
}

/*UnidadAcademica permite mostar la unidad academica*/
type UnidadAcademica struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}
