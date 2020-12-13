package tipounidad

import (
	"database/sql"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	crearTipoUnidad(params *addTipoUnidadRequest) (int, error)
	ObtenerTodosLosTiposDeUnidad() ([]*TipoUnidad, error)
	ObtenerUnidadAcademica(unidadAcadID int) ([]*UnidadAcademica, error)
	ObtenerTipoDeUnidadByID(param *getTipoUnidadByIDRequest) (*TipoUnidad, error)
	ActualizarTipoUnidad(params *updateTipoUnidadRequest) (int, error)
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

func (repo *repository) crearTipoUnidad(params *addTipoUnidadRequest) (int, error) {
	const querySrt = `INSERT INTO TIPO_UNIDAD(NOMBRE, DESCRIPCION) VALUES(?, ?)`
	result, err := repo.db.Exec(querySrt, params.Nombre, params.Descripcion)
	if err != nil {
		return 0, err
	}
	filasAfectadas, err := result.LastInsertId()
	return int(filasAfectadas), err
}

func (repo *repository) ObtenerTodosLosTiposDeUnidad() ([]*TipoUnidad, error) {
	const queryStr = `SELECT TIPO_U_ID, NOMBRE, DESCRIPCION FROM TIPO_UNIDAD`
	result, err := repo.db.Query(queryStr)
	var tipoUnidades []*TipoUnidad
	for result.Next() {
		tipoUnidad := &TipoUnidad{}
		err := result.Scan(
			&tipoUnidad.ID,
			&tipoUnidad.Nombre,
			&tipoUnidad.Descripcion,
		)
		if err != nil {
			return nil, err
		}
		unidadAcad, _ := repo.ObtenerUnidadAcademica(tipoUnidad.ID)
		tipoUnidad.UnidadAcad = unidadAcad
		tipoUnidades = append(tipoUnidades, tipoUnidad)
	}
	return tipoUnidades, err
}

func (repo *repository) ObtenerUnidadAcademica(unidadAcadID int) ([]*UnidadAcademica, error) {
	const queryStr = `SELECT UNIDAD_ACAD_ID, NOMBRE FROM UNIDAD_ACAD WHERE TIPO_U_ID = ?`
	result, err := repo.db.Query(queryStr, unidadAcadID)
	var unidadesAcad []*UnidadAcademica
	for result.Next() {
		unidadAcad := &UnidadAcademica{}
		err := result.Scan(
			&unidadAcad.ID,
			&unidadAcad.Nombre,
		)
		if err != nil {
			return nil, err
		}
		unidadesAcad = append(unidadesAcad, unidadAcad)
	}
	return unidadesAcad, err
}

func (repo *repository) ObtenerTipoDeUnidadByID(param *getTipoUnidadByIDRequest) (*TipoUnidad, error) {
	const queryStr = `SELECT TIPO_U_ID, NOMBRE, DESCRIPCION FROM TIPO_UNIDAD WHERE TIPO_U_ID = ?`
	row := repo.db.QueryRow(queryStr, param.ID)
	tipoUnidad := &TipoUnidad{}
	err := row.Scan(&tipoUnidad.ID, &tipoUnidad.Nombre, &tipoUnidad.Descripcion)
	if err != nil {
		return nil, err
	}
	unidadAcad, _ := repo.ObtenerUnidadAcademica(tipoUnidad.ID)
	tipoUnidad.UnidadAcad = unidadAcad
	return tipoUnidad, err
}

func (repo *repository) ActualizarTipoUnidad(params *updateTipoUnidadRequest) (int, error) {
	const queryStr = `UPDATE TIPO_UNIDAD SET NOMBRE = ?, DESCRIPCION = ? WHERE TIPO_U_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Descripcion, params.ID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowAffected), nil
}
