package usuario

import (
	"database/sql"
	"errors"
	"fmt"
)

/*Repository nos sirve para poder realizar consultas a la BDs*/
type Repository interface {
	ChequeoUsuarioCreado(personaID int) (int, error)
	ChequeoEmailExisteUsuario(email string) (int, error)
	RegistrarUsuario(params *registerUserRequest) (int, error)
	BuscarPersona(param int) (int, int, error)
	BuscarImagenUsuario(params *obtenerAvatarRequest) (*obtenerAvatarRequest, error)
	SubirImagenUsuario(param *subirAvartarRequest, usuaioID int) (int, error)
	ObtenerTodosLosUsuarios() ([]*Usuario, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository permitirá crear el repositorio y retornar asi misma*/
func NewRepository(pdb *sql.DB) Repository {
	return &repository{
		db: pdb,
	}
}

func (repo *repository) ChequeoEmailExisteUsuario(email string) (int, error) {
	contCorreo := 1
	const queryStrCont = `SELECT COUNT(*) FROM USUARIO WHERE EMAIL = ? AND ESTADO_ELIMINADO = 1`
	rowCont := repo.db.QueryRow(queryStrCont, email)
	rowCont.Scan(&contCorreo)
	return contCorreo, nil
}

func (repo *repository) ChequeoUsuarioCreado(personaID int) (int, error) {
	contCorreo := -1
	const queryStrCont = `SELECT COUNT(*) FROM USUARIO WHERE PERSONA_ID = ?`
	rowCont := repo.db.QueryRow(queryStrCont, personaID)
	err := rowCont.Scan(&contCorreo)
	return contCorreo, err
}

func (repo *repository) RegistrarUsuario(params *registerUserRequest) (int, error) {
	const queryStr = `INSERT INTO USUARIO(PERSONA_ID, USER_NAME, EMAIL, CLAVE)
	VALUES(?, ?, ?, ?)`

	result, err := repo.db.Exec(queryStr, params.PersonaID, params.UserName,
		params.Email, params.Password)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

/*BuscarPersonaExistente buscamos si la persona existe en la BD*/
func (repo *repository) BuscarPersona(param int) (int, int, error) {
	contResult := 0
	estadoPersona := -1
	const queryStr = `SELECT COUNT(*), ESTADO_ELIMINADO FROM PERSONA WHERE PERSONA_ID = ?`
	result := repo.db.QueryRow(queryStr, param)
	err := result.Scan(&contResult, &estadoPersona)
	return contResult, estadoPersona, err
}

func (repo *repository) SubirImagenUsuario(param *subirAvartarRequest, usuarioID int) (int, error) {
	const queryStr = `UPDATE USUARIO SET AVATAR = ? WHERE PERSONA_ID = ?`
	result, err := repo.db.Exec(queryStr, param.File, usuarioID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) BuscarImagenUsuario(params *obtenerAvatarRequest) (*obtenerAvatarRequest, error) {
	const queryStr = `SELECT AVATAR FROM USUARIO WHERE PERSONA_ID = ?`
	result := repo.db.QueryRow(queryStr, params.ID)
	avatarUsuario := &obtenerAvatarRequest{}
	err := result.Scan(&avatarUsuario.File)
	avatarUsuario.ID = params.ID
	return avatarUsuario, err
}

func (repo *repository) ObtenerTodosLosUsuarios() ([]*Usuario, error) {
	const queryStr = `SELECT PERSONA_ID, USER_NAME, EMAIL, DNI, ESTADO_PERSONA, ESTADO_USUARIO FROM VW_USUARIO`
	results, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error al consultar la BD ", err))
	}
	var usuarios []*Usuario
	for results.Next() {
		usuario := &Usuario{}
		err = results.Scan(&usuario.ID, &usuario.NombreUsuario, &usuario.Email, &usuario.DNI, &usuario.EstadoPersona, &usuario.EstadoUsuario)
		if err != nil {
			return nil, errors.New(fmt.Sprint("Error al escanear las registros usuarios ", err))
		}
		rol, err := repo.ObtenerRolUsuario(usuario.ID)
		if err != nil {
			return nil, err
		}
		usuario.Rol = rol
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repo *repository) ObtenerRolUsuario(personaID int) ([]*string, error) {
	const queryStr = `SELECT ROL FROM VW_ROL_USUARIO WHERE PERSONA_ID = ?`
	results, err := repo.db.Query(queryStr, personaID)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error la hacer la consulta en la BD ", err))
	}
	var roles []*string
	for results.Next() {
		var rol *string
		err := results.Scan(&rol)
		if err != nil {
			return nil, errors.New(fmt.Sprint("Error al escanear los registros rol ", err))
		}
		roles = append(roles, rol)
	}
	return roles, nil
}
