package unidadacademica

import (
	"database/sql"

	"github.com/API_REST_BDII_LP2_MYSQL/tipounidad"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (int, error)
	ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error)
	ObtenerTipoUnidadByID(idTipoUnidad int) (*tipounidad.TipoUnidad, error)
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
func (repo *repository) AgregarUnidadAcademica(params *addUnidadAcademicaRequest) (int, error) {
	const queryStr = `INSERT INTO UNIDAD_ACAD(TU_ID, NOMBRE) VALUES(? , ?)`
	result, err := repo.db.Exec(queryStr, params.TipoUnidadID, params.Nombre)
	if err != nil {
		return 0, err
	}
	idRowInsert, err := result.LastInsertId()
	return int(idRowInsert), err
}

func (repo *repository) ObtenerUnidadAcademicaByID(param *idUnidadAcademicaRequest) (*UnidadAcademica, error) {
	const querySrt = `SELECT UNIDAD_ACAD_ID, TU_ID, NOMBRE FROM UNIDAD_ACAD WHERE UNIDAD_ACAD_ID = ? AND ESTADO_ELIMINADO = 1;`
	result := repo.db.QueryRow(querySrt, param.ID)
	unidadAcademica := &UnidadAcademica{}
	err := result.Scan(&unidadAcademica.ID, &unidadAcademica.TipoUnidadID, &unidadAcademica.Nombre)
	return unidadAcademica, err
}

func (repo *repository) ObtenerTipoUnidadByID(idTipoUnidad int) (*tipounidad.TipoUnidad, error) {
	const queryStr = `SELECT TU_ID, NOMBRE, DESCRIPCION FROM TIPO_UNIDAD WHERE TU_ID = ?`
	result := repo.db.QueryRow(queryStr, idTipoUnidad)
	tipoUnidad := &tipounidad.TipoUnidad{}
	err := result.Scan(&tipoUnidad.ID, &tipoUnidad.Nombre, &tipoUnidad.Descripcion)
	return tipoUnidad, err
}
