package usuariologin

import (
	"database/sql"
)

//Repository Tendremos un metodos en la interface para implementar en una estructura
type Repository interface {
	ChequeoExisteUsuario(email *string) (*Usuario, int, error)
	// ChequeoExisteUsuarioPersona(params *passwordResetRequest) (*Usuario, int, error)
	// ActualizarPasswordUsuario(params *Usuario) (int, error)
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
	err := row.Scan(&usuario.UsuarioID, &usuario.UsuarioNombre,
		&usuario.UsuarioEmail, &usuario.UsuarioPassword,
		&usuario.UsuarioAvatar)
	if err != nil {
		return nil, 0, err
	}
	const queryStrCont = `SELECT COUNT(EMAIL) FROM USUARIO WHERE EMAIL = ?`
	rowCont := repo.db.QueryRow(queryStrCont, email)
	errr := rowCont.Scan(&contCorreo)
	return usuario, contCorreo, errr
}

/*
func (repo *repository) ChequeoExisteUsuarioPersona(params *passwordResetRequest) (*Usuario, int, error) {
	cantResult := 0
	usuario := &Usuario{}
	const queryStr = `SELECT COUNT(U.PERSONA_ID), U.PERSONA_ID, U.USER_NAME, U.EMAIL, U.AVATAR
	 FROM USUARIO U INNER JOIN PERSONA P ON U.PERSONA_ID = P.PERSONA_ID WHERE
	 P.NOMBRE = ? AND
	 P.APELLIDO_P = ? AND
	 P.APELLIDO_M = ? AND
	 P.FECHA_NACIMIENTO = ? AND
	 U.EMAIL = ? AND
	 P.ESTADO = 1 AND
	 U.ESTADO_ELIMINADO = 1
	 `
	result := repo.db.QueryRow(queryStr, params.NombrePersonal,
		params.ApellidoPaterno, params.ApellidoMaterno, params.FechaNacimiento,
		params.Email)
	err := result.Scan(&cantResult, &usuario.UsuarioID, &usuario.UsuarioNombre,
		&usuario.UsuarioEmail, &usuario.UsuarioAvatar)
	return usuario, cantResult, err
}

func (repo *repository) ActualizarPasswordUsuario(params *Usuario) (int, error) {
	const queryStr = `
UPDATE USUARIO SET CLAVE = ? WHERE PERSONA_ID = ? AND EMAIL = ?`
	result, err := repo.db.Exec(queryStr, params.UsuarioPassword,
		params.UsuarioID, params.UsuarioEmail)
	if err != nil {
		return 0, err
	}
	cantAfectados, err := result.RowsAffected()
	return int(cantAfectados), err
}
*/
