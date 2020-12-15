package plancurso

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarPlanCurso(params *addPlanCursoRequest) (int, int, error)
	//ObtenerPlanCursoPorID(param *getPlanCursoByIDRequest) error
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

func (repo *repository) AgregarPlanCurso(params *addPlanCursoRequest) (int, int, error) {
	const queryStr = `INSERT INTO PLAN_CURSO (CURSO_ID, PLAN_ID, CICLO) VALUES(?, ?, ?)`
	result, err := repo.db.Exec(queryStr, params.CursoID, params.PlanID, params.Ciclo)
	if err != nil {
		return -1, -1, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}
	planCursoID, err := result.LastInsertId()

	return int(planCursoID), int(rowAffected), err
}
