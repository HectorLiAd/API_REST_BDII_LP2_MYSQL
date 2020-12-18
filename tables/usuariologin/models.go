package usuariologin

/*Usuario para poder mostrar al usuario en el response*/
type Usuario struct {
	UsuarioID       int       `json:"usuario_id"`
	UsuarioNombre   string    `json:"nombre"`
	UsuarioEmail    string    `json:"email"`
	UsuarioPassword string    `json:"password,omitempty"`
	UsuarioAvatar   string    `json:"avatar,omitempty"`
	Rol             []*string `json:"rol"`
}

// type Rol struct {
// }

/*RespuestaLogin token que será generado al iniciar sesión*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
