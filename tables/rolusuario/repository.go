package rolusuario

import (
	"database/sql"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarRolUsuario(params *addRolUsuarioRequest) (int, int, error)
	ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error)
	ObtenerTodosRolUsuario() ([]*RolUsuario, error)
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
	const queryStr = `INSERT INTO ROL_USUARIO (ROL_ID, PERSONA_ID) VALUES(?, ?)`
	result, err := repo.db.Exec(queryStr, params.RolID, params.PersonaID)
	if err != nil {
		return 0, 0, err
	}
	rolUsuarioID, errLI := result.LastInsertId()
	if errLI != nil {
		return 0, 0, errLI
	}
	rowAffected, errRA := result.RowsAffected()
	return int(rolUsuarioID), int(rowAffected), errRA
}

func (repo *repository) ObtenerRolUsuarioPorID(param *getRolUsuarioByIDRequest) (*RolUsuario, error) {
	const queryStr = `SELECT ROL_US_ID,ROL,USER_NAME FROM VW_ROL_USUARIO WHERE ROL_US_ID = ?`
	rowRolUsuario := repo.db.QueryRow(queryStr, param.ID)
	rolUsuario := &RolUsuario{}
	err := rowRolUsuario.Scan(&rolUsuario.ID, &rolUsuario.Rol, &rolUsuario.UserName)
	return rolUsuario, err
}

func (repo *repository) ObtenerTodosRolUsuario() ([]*RolUsuario, error) {
	const queryStr = `SELECT ROL_US_ID,ROL,USER_NAME FROM VW_ROL_USUARIO`
	rowsRolUsuario, errQ := repo.db.Query(queryStr)
	if errQ != nil {
		return nil, errQ
	}
	var rolesUsuarios []*RolUsuario
	for rowsRolUsuario.Next() {
		rolUsuario := &RolUsuario{}
		errScan := rowsRolUsuario.Scan(&rolUsuario.ID, &rolUsuario.Rol, &rolUsuario.UserName)
		if errScan != nil {
			return nil, errScan
		}
		rolesUsuarios = append(rolesUsuarios, rolUsuario)
	}
	return rolesUsuarios, nil
}
