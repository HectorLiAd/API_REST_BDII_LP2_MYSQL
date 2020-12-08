package tiporecurso

/*TipoRecurso permite mostrar en formato json*/
type TipoRecurso struct {
	ID                 int    `json:"id"`
	Nombre             string `json:"nombre"`
	EstadoCalificativo int    `json:"estadoCalificativo"`
	BloquearRecurso    int    `json:"bloquearRecurso"`
}
