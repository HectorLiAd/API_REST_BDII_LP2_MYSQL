package usuario

import (
	"database/sql"

	"github.com/API_REST_BDII_LP2_MYSQL/models"
)

/*Repository nos sirve para poder realizar consultas a la BDs*/
type Repository interface {
	ChequeoUsuarioCreado(personaID int) (int, error)
	ChequeoEmailExisteUsuario(email string) (int, error)
	RegistrarUsuario(params *registerUserRequest) (*models.ResultOperacion, error)
	BuscarPersona(param int) (int, int, error)
	SubirImagenUsuario(param *subirAvartarRequest, usuaioID int) (int, error)
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

func (repo *repository) RegistrarUsuario(params *registerUserRequest) (*models.ResultOperacion, error) {
	const queryStr = `INSERT INTO USUARIO(PERSONA_ID, USER_NAME, EMAIL, CLAVE, AVATAR)
	VALUES(?, ?, ?, ?, ?)`

	result, err := repo.db.Exec(queryStr, params.PersonaID, params.UserName,
		params.Email, params.Password, params.Avatar)
	id, er := result.LastInsertId()
	if er != nil {
		return nil, er
	}
	rowAffected, err := result.RowsAffected()
	return &models.ResultOperacion{
		Name:        "Usuario " + params.UserName + " registrado correctamente",
		Codigo:      int(id),
		RowAffected: int(rowAffected),
	}, err
}

/*BuscarPersonaExistente buscamos si la persona existe en la BD*/
func (repo *repository) BuscarPersona(param int) (int, int, error) {
	contResult := 0
	estadoPersona := -1
	const queryStr = `SELECT COUNT(*), ESTADO FROM PERSONA WHERE PERSONA_ID = ?`
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
