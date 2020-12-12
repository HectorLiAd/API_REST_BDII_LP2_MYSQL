package rol

import (
	"database/sql"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	InsertarRol(param *addRolRequest) (int64, int64, error)
	ActualizarRol(params *updateRolRequest) (int, error)
	ObtenerRolByID(param *getRolByIDRequest) (*Rol, error)
	ObtenerTodosLosRoles() ([]*Rol, error)
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

func (re *repository) InsertarRol(param *addRolRequest) (int64, int64, error) {
	const queryStr = `INSERT INTO ROL (NOMBRE) VALUES (?)`
	result, err := re.db.Exec(queryStr, param.Nombre)
	if err != nil {
		return 0, 0, err
	}
	rolID, err := result.LastInsertId()
	if err != nil {
		return 0, 0, nil
	}
	rowAffected, err := result.RowsAffected()
	return rolID, rowAffected, err
}

func (re *repository) ActualizarRol(params *updateRolRequest) (int, error) {
	const queryStr = `UPDATE ROL SET NOMBRE = ? WHERE ROL_ID = ?`
	result, err := re.db.Exec(queryStr, params.Nombre, params.ID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (re *repository) ObtenerRolByID(param *getRolByIDRequest) (*Rol, error) {
	const queryStr = `SELECT ROL_ID, NOMBRE FROM ROL WHERE ROL_ID = ?`
	row := re.db.QueryRow(queryStr, param.ID)
	rol := &Rol{}
	err := row.Scan(&rol.ID, &rol.Nombre)
	return rol, err
}

func (re *repository) ObtenerTodosLosRoles() ([]*Rol, error) {
	const queryStr = `SELECT ROL_ID, NOMBRE FROM ROL`
	rows, errQ := re.db.Query(queryStr)
	if errQ != nil {
		return nil, errQ
	}
	var roles []*Rol
	for rows.Next() {
		rol := &Rol{}
		errS := rows.Scan(&rol.ID, &rol.Nombre)
		if errS != nil {
			return nil, errS
		}
		roles = append(roles, rol)
	}
	return roles, nil
}
