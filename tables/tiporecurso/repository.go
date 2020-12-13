package tiporecurso

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarTipoRecurso(params *addTipoRecursoRequest) (int, int, error)
	ObtenerTodoTipoRecurso() ([]*TipoRecurso, error)
	ObtenerTipoRecursoPorID(param *getTipoRecursoByIDRequest) (*TipoRecurso, error)
	ActualizarTipoRecurso(params *updateTipoRecursoRequest) (int, error)
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

func (repo *repository) AgregarTipoRecurso(params *addTipoRecursoRequest) (int, int, error) {
	const queryStr = `INSERT INTO TIPO_RECURSO(NOMBRE, ESTADO_CALIFICATIVO, BLOQUEAR_RECURSO) VALUES (?, ?, ?)`
	resultInsert, err := repo.db.Exec(queryStr, params.Nombre,
		params.EstadoCalificativo, params.BloquearRecurso)
	if err != nil {
		return 0, 0, err
	}
	tipoRecursoID, err := resultInsert.LastInsertId()
	if err != nil {
		return 0, 0, nil
	}
	rowAffected, err := resultInsert.RowsAffected()
	return int(tipoRecursoID), int(rowAffected), nil
}

func (repo *repository) ObtenerTodoTipoRecurso() ([]*TipoRecurso, error) {
	const queryStr = `SELECT TIPO_R_ID, NOMBRE, ESTADO_CALIFICATIVO, BLOQUEAR_RECURSO FROM TIPO_RECURSO`
	result, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var tiposRecursos []*TipoRecurso
	for result.Next() {
		tipoRecurso := &TipoRecurso{}
		err := result.Scan(
			&tipoRecurso.ID,
			&tipoRecurso.Nombre,
			&tipoRecurso.EstadoCalificativo,
			&tipoRecurso.BloquearRecurso,
		)
		if err != nil {
			return nil, err
		}
		tiposRecursos = append(tiposRecursos, tipoRecurso)
	}
	return tiposRecursos, nil
}

func (repo *repository) ObtenerTipoRecursoPorID(param *getTipoRecursoByIDRequest) (*TipoRecurso, error) {
	const queryStrt = `SELECT TIPO_R_ID, NOMBRE, ESTADO_CALIFICATIVO, BLOQUEAR_RECURSO FROM TIPO_RECURSO WHERE TIPO_R_ID = ?`
	result := repo.db.QueryRow(queryStrt, param.ID)
	tipoRecurso := &TipoRecurso{}
	err := result.Scan(
		&tipoRecurso.ID,
		&tipoRecurso.Nombre,
		&tipoRecurso.EstadoCalificativo,
		&tipoRecurso.BloquearRecurso,
	)
	return tipoRecurso, err
}

func (repo *repository) ActualizarTipoRecurso(params *updateTipoRecursoRequest) (int, error) {
	const queryStr = `UPDATE TIPO_RECURSO SET NOMBRE = ?, ESTADO_CALIFICATIVO = ?, BLOQUEAR_RECURSO = ? WHERE TIPO_R_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.EstadoCalificativo, params.BloquearRecurso, params.ID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}
