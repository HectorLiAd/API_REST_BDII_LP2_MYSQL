package tipounidad

/*TipoUnidad permite transformar la informacion de la BD a un json segun sugun la esquema de esta estructura*/
type TipoUnidad struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}
