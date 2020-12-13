package unidadacademica

import (
	"database/sql"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (int, int, error)
	ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error)
	ActualizarUnidadAcademicaByID(params *updateUnidadAcademicaRequest) (int, error)
	ObtenerTodasLasUnidadesAcademicas() ([]*UnidadAcademica, error)
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
	const queryStr = `INSERT INTO UNIDAD_ACAD(TIPO_U_ID, NOMBRE) VALUES(? , ?)`
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

func (repo *repository) ActualizarUnidadAcademicaByID(params *updateUnidadAcademicaRequest) (int, error) {
	const queryStr = `UPDATE UNIDAD_ACAD SET NOMBRE = ? WHERE UNIDAD_ACAD_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.ID)
	if err != nil {
		return 0, nil
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodasLasUnidadesAcademicas() ([]*UnidadAcademica, error) {
	const queryStr = `SELECT UNIDAD_ACAD_ID, TIPO_UNIDAD, NOMBRE FROM VW_UNIDAD_ACADEMICA`
	result, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var unidadesAcademicas []*UnidadAcademica
	for result.Next() {
		unidadAcademica := &UnidadAcademica{}
		errScan := result.Scan(&unidadAcademica.ID, &unidadAcademica.TipoUnidad, &unidadAcademica.Nombre)
		if errScan != nil {
			return nil, errScan
		}
		unidadesAcademicas = append(unidadesAcademicas, unidadAcademica)
	}
	return unidadesAcademicas, nil
}
