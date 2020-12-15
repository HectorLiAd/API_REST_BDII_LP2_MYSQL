package curso

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarCurso(params *addCursoRequest) (int, int, error)
	ObtenerCursoPorID(param *getCursoByIDRequest) (*Curso, error)
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

func (repo *repository) RegistrarCurso(params *addCursoRequest) (int, int, error) {
	const queryStr = `INSERT INTO CURSO(NOMBRE, DETALLE) VALUES(?, ?)`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Detalle)
	if err != nil {
		return -1, -1, err
	}
	cursoID, err := result.LastInsertId()
	if err != nil {
		return -1, -1, nil
	}
	rowAffected, err := result.RowsAffected()
	return int(cursoID), int(rowAffected), err
}

func (repo *repository) ObtenerCursoPorID(param *getCursoByIDRequest) (*Curso, error) {
	const queryStr = `SELECT CURSO_ID, NOMBRE, DETALLE FROM CURSO WHERE CURSO_ID = ? AND ESTADO_ELIMINADO = 1`
	result := repo.db.QueryRow(queryStr, param.ID)
	curso := &Curso{}
	err := result.Scan(&curso.ID, &curso.Nombre, &curso.Descripcion)
	return curso, err
}
