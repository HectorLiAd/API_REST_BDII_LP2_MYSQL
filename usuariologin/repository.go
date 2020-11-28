package usuariologin

import (
	"database/sql"
)

//Repository Tendremos un metodos en la interface para implementar en una estructura
type Repository interface {
	ChequeoExisteUsuario(email *string) (*Usuario, int, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository creara el repositorio*/
func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) ChequeoExisteUsuario(email *string) (*Usuario, int, error) {
	contCorreo := 2
	usuario := &Usuario{}
	const queryStr = `SELECT PERSONA_ID, USER_NAME, EMAIL, CLAVE, AVATAR from USUARIO WHERE EMAIL = ? AND ESTADO_ELIMINADO = 1`
	row := repo.db.QueryRow(queryStr, email)
	row.Scan(&usuario.UsuarioID, &usuario.UsuarioNombre,
		&usuario.UsuarioEmail, &usuario.UsuarioPassword,
		&usuario.UsuarioAvatar)

	const queryStrCont = `SELECT COUNT(EMAIL) FROM USUARIO WHERE EMAIL = ?`
	rowCont := repo.db.QueryRow(queryStrCont, email)
	rowCont.Scan(&contCorreo)
	return usuario, contCorreo, nil
}
