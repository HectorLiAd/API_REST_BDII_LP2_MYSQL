package unidadacademica

import (
	"database/sql"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (int, int, error)
	ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository permite crear el repositoriio*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}
func (repo *repository) AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (int, int, error) {
	const queryStr = `INSERT INTO UNIDAD_ACAD(TU_ID, NOMBRE) VALUES(? , ?)`
	result, err := repo.db.Exec(queryStr, params.TipoUnidadID, params.Nombre)
	if err != nil {
		return 0, 0, err
	}
	idRowInsert, err := result.LastInsertId()
	rowAffected, err := result.RowsAffected()
	return int(idRowInsert), int(rowAffected), err
}

func (repo *repository) ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error) {
	const querySrt = `SELECT UNIDAD_ACAD_ID, TIPO_UNIDAD, NOMBRE FROM VW_UNIDAD_ACADEMICA WHERE UNIDAD_ACAD_ID = ?`
	result := repo.db.QueryRow(querySrt, param.ID)
	unidadAcademica := &UnidadAcademica{}
	err := result.Scan(&unidadAcademica.ID, &unidadAcademica.TipoUnidad, &unidadAcademica.Nombre)
	return unidadAcademica, err
}
