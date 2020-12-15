package curso

import "database/sql"

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarCurso(params *addCursoRequest) (int, int, error)
	ObtenerCursoPorID(param *getCursoByIDRequest) (*Curso, error)
	ActualizatCursoPorID(params *updateCursoByIDRequest) (int, error)
	ObtenerTodosLosCursos() ([]*Curso, error)
	SubirFondoCurso(param *updateImagenCursoByIDRequest) (int, error)
	ObtenerFondoCurso(params *getFondoCursoByIDRequest) (*getFondoCursoByIDRequest, error)
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

func (repo *repository) ActualizatCursoPorID(params *updateCursoByIDRequest) (int, error) {
	const queryStr = `UPDATE CURSO SET NOMBRE = ?, DETALLE = ? WHERE CURSO_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Detalle, params.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodosLosCursos() ([]*Curso, error) {
	const queryStr = `SELECT CURSO_ID, NOMBRE, DETALLE FROM CURSO`
	results, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var cursos []*Curso
	for results.Next() {
		curso := &Curso{}
		err := results.Scan(&curso.ID, &curso.Nombre, &curso.Descripcion)
		if err != nil {
			return nil, err
		}
		cursos = append(cursos, curso)
	}
	return cursos, err
}

func (repo *repository) SubirFondoCurso(param *updateImagenCursoByIDRequest) (int, error) {
	const queryStr = `UPDATE CURSO SET FONDO_IMG = ? WHERE CURSO_ID = ?`
	result, err := repo.db.Exec(queryStr, param.File, param.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerFondoCurso(params *getFondoCursoByIDRequest) (*getFondoCursoByIDRequest, error) {
	const querStr = `SELECT FONDO_IMG FROM CURSO WHERE CURSO_ID = ?`
	result := repo.db.QueryRow(querStr, params.ID)
	err := result.Scan(&params.File)
	return params, err
}
