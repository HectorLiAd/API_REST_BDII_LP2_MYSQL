package cargaacademica

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarCargaAcad(params *addCargaAcademica) (int, int, error)
	ObtenerCargaAcadPorID(param *getCargaAcadByID) (*CargaAcademica, error)
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

func (repo *repository) RegistrarCargaAcad(params *addCargaAcademica) (int, int, error) {
	const querStr = `INSERT INTO CARGA_ACADEMICA(JERARQUIA_ID, PERSONA_ID, PERIODO_ID, PLAN_CURSO_ID) VALUES (?, ?, ?, ?)`
	result, err := repo.db.Exec(querStr, params.JerarquiaID, params.PersonaID, params.PeriodoID, params.PlanCursoID)
	if err != nil {
		return -1, -1, nil
	}
	cargaAcademicaID, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(cargaAcademicaID), int(rowAffected), err
}

func (repo *repository) ObtenerCargaAcadPorID(param *getCargaAcadByID) (*CargaAcademica, error) {
	const queryStr = `SELECT C_ACAD_ID, JERARQUIA_ID, PERSONA_ID, PERIODO_ID, PLAN_CURSO_ID, FORMATO_ESTADO FROM CARGA_ACADEMICA WHERE C_ACAD_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)

	cargaAcad := &CargaAcademica{}
	err := result.Scan(&cargaAcad.ID, &cargaAcad.JerarquiaID, &cargaAcad.PersonaID, &cargaAcad.PeriodoID, &cargaAcad.PlanCursoID, &cargaAcad.FormatoEstado)
	if err != nil {
		return nil, err
	}
	return cargaAcad, err
}
