package usuario

/*Usuario para poder mostrar al usuario en el response*/
type Usuario struct {
	ID            int       `json:"usuario_id"`
	NombreUsuario string    `json:"nombreuUsuario"`
	Email         string    `json:"email"`
	DNI           string    `json:"dni"`
	EstadoPersona string    `json:"estadoPersona"`
	EstadoUsuario string    `json:"estadoUsuario"`
	Rol           []*string `json:"rol"`
}
