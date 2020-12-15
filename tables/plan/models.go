package plan

/*Plan nos sirve para poder mostrar el JSON*/
type Plan struct {
	ID          int    `json:"id"`
	JerarquiaID int    `json:"jerarquiaId"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
