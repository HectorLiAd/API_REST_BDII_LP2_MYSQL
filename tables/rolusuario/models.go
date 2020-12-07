package rolusuario

/*RolUsuario struct para poder mapear a un formato JSON*/
type RolUsuario struct {
	ID       int    `json:"id"`
	Rol      string `json:"rol"`
	UserName string `json:"userName"`
}
