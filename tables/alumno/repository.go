package alumno

import (
	"database/sql"

	"github.com/API_REST_BDII_LP2_MYSQL/tables/persona"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	AgregarPersonaAlumno(param *addAlumnoRequest) (int, int, error)
	ExistePersonaAlumnoPorID(param *getAlumnoByIDRequest) (int, error)
	BuscarPersonaPorID(personaID int) (*persona.Person, error)
	ObtenerTodoPersonaAlumno() ([]*Alumno, error)
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

func (repo *repository) AgregarPersonaAlumno(param *addAlumnoRequest) (int, int, error) {
	const queryStr = `INSERT INTO ALUMNO (PERSONA_ID) VALUES (?)`
	result, err := repo.db.Exec(queryStr, param.ID)
	if err != nil {
		return 0, 0, err
	}
	rowAffected, err := result.RowsAffected()
	return param.ID, int(rowAffected), err
}

func (repo *repository) BuscarPersonaPorID(personaID int) (*persona.Person, error) {
	queryStr := `SELECT PERSONA_ID, NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, 
	DNI, FECHA_NACIMIENTO FROM PERSONA WHERE PERSONA_ID = ? AND ESTADO_ELIMINADO = 1`
	result := repo.db.QueryRow(queryStr, personaID)
	var fechaUint []uint8
	persona := &persona.Person{}
	err := result.Scan(&persona.ID, &persona.Nombre, &persona.ApellidoPaterno,
		&persona.ApellidoMaterno, &persona.Genero, &persona.DNI, &fechaUint)
	if err != nil {
		return nil, err
	}
	// fec, _ := helper.ConvStrADate(string(fechaUint))
	persona.FechaNacimiento = string(fechaUint)
	return persona, nil
}

func (repo *repository) ExistePersonaAlumnoPorID(param *getAlumnoByIDRequest) (int, error) {
	const queryStr = `SELECT COUNT(*) FROM ALUMNO WHERE PERSONA_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)
	var cant int = 0
	err := result.Scan(&cant)
	return cant, err
}

func (repo *repository) ObtenerTodoPersonaAlumno() ([]*Alumno, error) {
	const queryStr = `SELECT PERSONA_ID FROM ALUMNO`
	results, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var alumnos []*Alumno
	for results.Next() {
		alumno := &Alumno{}
		err := results.Scan(&alumno.ID)
		if err != nil {
			return nil, err
		}
		alumno.Persona, err = repo.BuscarPersonaPorID(alumno.ID)
		if err != nil {
			return nil, err
		}
		alumnos = append(alumnos, alumno)
	}
	return alumnos, err
}
