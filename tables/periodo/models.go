package periodo

/*Periodo permite mapearse en formato JSON*/
type Periodo struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	FechaIni string `json:"fechaInicio"`
	FechaFin string `json:"fechaFin"`
}
