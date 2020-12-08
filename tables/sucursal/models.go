package sucursal

/*Sucursal modelos que se mastrara en formato json*/
type Sucursal struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	Descripcion string `json:"descripcion"`
}
