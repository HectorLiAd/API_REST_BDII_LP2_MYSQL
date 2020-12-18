package cargaacademica

/*CargaAcademica nos sirve para que actue como JSON por el mapeo*/
type CargaAcademica struct {
	ID            int `json:"id"`
	JerarquiaID   int `json:"jerarquiaId"`
	PersonaID     int `json:"pesonaId"`
	PeriodoID     int `json:"periodoId"`
	PlanCursoID   int `json:"planCursoId"`
	FormatoEstado int `json:"formatoEstado"`
}
