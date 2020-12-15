package periodo

import (
	"database/sql"
	"errors"
	"fmt"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarPeriodo(params *addPeriodoRequest) (int, int, error)
	ObtenerPeriodoPorID(param *getPeridioByIDRequest) (*Periodo, error)
	ActualizarPeriodo(params *updatePeriodoRequest) (int, error)
	ObtenerTodosLosPeriodos() ([]*Periodo, error)
	EliminarPeriodoPorID(param *deletePeriodoByIDRequest) (int, error)
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

func (repo *repository) RegistrarPeriodo(params *addPeriodoRequest) (int, int, error) {
	const queryStr = `INSERT INTO PERIODO (NOMBRE, FECHA_INI, FECHA_FIN) VALUES (?, ?, ?)`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.FechaIni, params.FechaFin)
	if err != nil {
		return -1, -1, errors.New(fmt.Sprint("Error al registrar en la BD ", err))
	}
	peridoID, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}
	return int(peridoID), int(rowAffected), err
}

func (repo *repository) ObtenerPeriodoPorID(param *getPeridioByIDRequest) (*Periodo, error) {
	const queryStr = `SELECT PERIODO_ID, NOMBRE, FECHA_INI, FECHA_FIN FROM PERIODO WHERE PERIODO_ID = ? AND ESTADO_ELIMINADO = 1`
	result := repo.db.QueryRow(queryStr, param.ID)
	periodo := &Periodo{}
	err := result.Scan(&periodo.ID, &periodo.Nombre, &periodo.FechaIni, &periodo.FechaFin)
	return periodo, err
}

func (repo *repository) ActualizarPeriodo(params *updatePeriodoRequest) (int, error) {
	const queryStr = `UPDATE PERIODO SET NOMBRE = ?, FECHA_INI = ?, FECHA_FIN = ? WHERE PERIODO_ID = ? AND ESTADO_ELIMINADO = 1`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.FechaIni, params.FechaFin, params.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodosLosPeriodos() ([]*Periodo, error) {
	const queryStr = `SELECT PERIODO_ID, NOMBRE, FECHA_INI, FECHA_FIN FROM PERIODO WHERE ESTADO_ELIMINADO = 1`
	result, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var periodos []*Periodo
	for result.Next() {
		periodo := &Periodo{}
		err := result.Scan(&periodo.ID, &periodo.Nombre, &periodo.FechaIni, &periodo.FechaFin)
		if err != nil {
			return nil, err
		}
		periodos = append(periodos, periodo)
	}
	return periodos, err
}

func (repo *repository) EliminarPeriodoPorID(param *deletePeriodoByIDRequest) (int, error) {
	const queryStr = `UPDATE PERIODO SET ESTADO_ELIMINADO = 0 WHERE PERIODO_ID = ? AND ESTADO_ELIMINADO = 1`
	result, err := repo.db.Exec(queryStr, param.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}
