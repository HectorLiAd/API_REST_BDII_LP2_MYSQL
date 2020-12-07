package rolusuario

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarRolUsuario(params *addRolUsuarioRequest) (int, int, error)
	ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository crear el nuevo repositorio y retorna con la BD conectada*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) AgregarRolUsuario(params *addRolUsuarioRequest) (int, int, error) {
	const queryStr = `INSERT INTO ROL_USUARIO (RU_ID, PERSONA_ID) VALUES(?, ?)`
	result, errE := repo.db.Exec(queryStr, params.RolID, params.PersonaID)
	if errE != nil {
		return 0, 0, errE
	}
	rolUsuarioID, errLI := result.LastInsertId()
	if errLI != nil {
		return 0, 0, errLI
	}
	rowAffected, errRA := result.RowsAffected()
	return int(rolUsuarioID), int(rowAffected), errRA
}

func (repo *repository) ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error) {
	const queryStr = `SELECT * FROM VW_ROL_USUARIO WHERE RO_US_ID = ?`
	rowRolUsuario := repo.db.QueryRow(queryStr, param.ID)
	rolUsuario := &RolUsuario{}
	err := rowRolUsuario.Scan(&rolUsuario.ID, &rolUsuario.Rol, &rolUsuario.UserName)
	return rolUsuario, err
}