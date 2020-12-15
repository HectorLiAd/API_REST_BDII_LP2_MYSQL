package plancurso

import (
	"database/sql"

	"github.com/API_REST_BDII_LP2_MYSQL/tables/curso"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/plan"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarPlanCurso(params *addPlanCursoRequest) (int, int, error)
	ObtenerPlanCursoPorID(param *getPlanCursoByIDRequest) (*PlanCurso, error)
	ActualizarPlanCursoPorID(params *updatePlanCursoRequest) (int, error)
	ObtenerTodoPlanCurso() ([]*PlanCurso, error)
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

func (repo *repository) ObtenerPlanCursoPorID(param *getPlanCursoByIDRequest) (*PlanCurso, error) {
	const queryStr = `SELECT PLAN_CURSO_ID, CURSO_ID, PLAN_ID, CICLO FROM PLAN_CURSO WHERE PLAN_CURSO_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)
	planCurso := &PlanCurso{}
	err := result.Scan(&planCurso.ID, &planCurso.CursoID, &planCurso.PlanID, &planCurso.Ciclo)
	curso, err := repo.obtenerCursoPorID(planCurso.CursoID)
	if err != nil {
		return nil, err
	}
	plan, err := repo.obtenerPlanPorID(planCurso.PlanID)
	if err != nil {
		return nil, err
	}
	planCurso.Curso = curso
	planCurso.Plan = plan
	planCurso.PlanID = 0
	planCurso.CursoID = 0
	return planCurso, err
}

func (repo *repository) obtenerCursoPorID(ID int) (*curso.Curso, error) {
	const queryStr = `SELECT CURSO_ID, NOMBRE, DETALLE FROM CURSO WHERE CURSO_ID = ? AND ESTADO_ELIMINADO = 1`
	result := repo.db.QueryRow(queryStr, ID)
	curso := &curso.Curso{}
	err := result.Scan(&curso.ID, &curso.Nombre, &curso.Descripcion)
	return curso, err
}

func (repo *repository) obtenerPlanPorID(ID int) (*plan.Plan, error) {
	const queryStr = `SELECT PLAN_ID, JERARQUIA_ID, NOMBRE, DESCRIPCION FROM VW_PLAN WHERE PLAN_ID = ?`
	result := repo.db.QueryRow(queryStr, ID)
	plan := &plan.Plan{}
	err := result.Scan(&plan.ID, &plan.JerarquiaID, &plan.Nombre, &plan.Descripcion)
	return plan, err

}

func (repo *repository) ActualizarPlanCursoPorID(params *updatePlanCursoRequest) (int, error) {
	const queryStr = `UPDATE PLAN_CURSO SET CURSO_ID = ?, PLAN_ID = ?, CICLO = ? WHERE PLAN_CURSO_ID = ?`
	result, err := repo.db.Exec(queryStr, params.CursoID, params.PlanID, params.Ciclo, params.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodoPlanCurso() ([]*PlanCurso, error) {
	const queryStr = `SELECT PLAN_CURSO_ID, CURSO_ID, PLAN_ID, CICLO FROM PLAN_CURSO`
	results, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var planCursos []*PlanCurso
	for results.Next() {
		planCurso := &PlanCurso{}
		err := results.Scan(&planCurso.ID, &planCurso.CursoID, &planCurso.PlanID, &planCurso.Ciclo)
		if err != nil {
			return nil, err
		}
		curso, err := repo.obtenerCursoPorID(planCurso.CursoID)
		if err != nil {
			return nil, err
		}
		plan, err := repo.obtenerPlanPorID(planCurso.PlanID)
		if err != nil {
			return nil, err
		}
		planCurso.Curso = curso
		planCurso.Plan = plan
		planCurso.PlanID = 0
		planCurso.CursoID = 0
		planCursos = append(planCursos, planCurso)
	}
	return planCursos, err
}
