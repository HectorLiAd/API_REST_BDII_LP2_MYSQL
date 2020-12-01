package tipounidad

import (
	"database/sql"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	crearTipoUnidad(params *addTipoUnidadRequest) (int, error)
	ObtenerTodosLosTiposDeUnidad() ([]*TipoUnidad, error)
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
	const queryStr = `SELECT TU_ID, NOMBRE, DESCRIPCION FROM TIPO_UNIDAD`
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
		tipoUnidades = append(tipoUnidades, tipoUnidad)
	}
	return tipoUnidades, err
}
