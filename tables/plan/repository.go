package plan

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarPlan(params *addPlanRequest) (int, int, error)
	ObtenerPlanPorID(param *getPlanByIDRequest) (*Plan, error)
	ActualizarPlan(params *updatePlanRequest) (int, error)
	ObtenerTodosLosPlanes() ([]*Plan, error)
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

func (repo *repository) RegistrarPlan(params *addPlanRequest) (int, int, error) {
	const queryStr = `INSERT INTO PLAN_E(JERARQUIA_ID, NOMBRE, DESCRIPCION) VALUES(?, ?, ?)`
	result, err := repo.db.Exec(queryStr, params.JerarquiaID, params.Nombre, params.Descripcion)
	if err != nil {
		return -1, -1, err
	}
	planID, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(planID), int(rowAffected), err
}

func (repo *repository) ObtenerPlanPorID(param *getPlanByIDRequest) (*Plan, error) {
	const queryStr = `SELECT PLAN_ID, JERARQUIA_ID, NOMBRE, DESCRIPCION FROM VW_PLAN WHERE PLAN_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)
	plan := &Plan{}
	err := result.Scan(&plan.ID, &plan.JerarquiaID, &plan.Nombre, &plan.Descripcion)
	return plan, err
}

func (repo *repository) ActualizarPlan(params *updatePlanRequest) (int, error) {
	const queryStr = `UPDATE PLAN_E SET NOMBRE = ?, DESCRIPCION = ? WHERE PLAN_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Descripcion, params.ID)
	if err != nil {
		return -1, nil
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodosLosPlanes() ([]*Plan, error) {
	const queryStr = `SELECT PLAN_ID, JERARQUIA_ID, NOMBRE, DESCRIPCION FROM VW_PLAN`
	result, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var planes []*Plan
	for result.Next() {
		plan := &Plan{}
		err := result.Scan(&plan.ID, &plan.JerarquiaID, &plan.Nombre, &plan.Descripcion)
		if err != nil {
			return nil, err
		}
		planes = append(planes, plan)
	}
	return planes, err
}
