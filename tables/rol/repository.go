package rol

import (
	"database/sql"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	InsertarRol(param *addRolRequest) (int64, int64, error)
	ActualizarRol(params *updateRolRequest) (int, error)
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
	const queryStr = `UPDATE ROL SET NOMBRE = ? WHERE RU_ID = ?`
	result, err := re.db.Exec(queryStr, params.Nombre, params.ID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}
