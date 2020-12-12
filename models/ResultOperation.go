package models

/*ResultOperacion obtener el resultado de una operacion*/
type ResultOperacion struct {
	Name        string `json:"mensaje"`
	Codigo      int    `json:"codigo"`
	RowAffected int    `json:"filasAfectadas"`
}
