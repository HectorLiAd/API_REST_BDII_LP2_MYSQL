package docente

import (
	"database/sql"

	"github.com/API_REST_BDII_LP2_MYSQL/tables/persona"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarDocente(param *docenteRequest) (int, error)
	ObtenerDocentePorID(param *docenteRequest) (*Docente, error)
	ObtenerPersonaPorID(ID int) (*persona.Person, error)
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

func (repo *repository) RegistrarDocente(param *docenteRequest) (int, error) {
	const queryStr = `INSERT INTO DOCENTE (PERSONA_ID) VALUES(?)`
	result, err := repo.db.Exec(queryStr, param.ID)
	if err != nil {
		return -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerDocentePorID(param *docenteRequest) (*Docente, error) {
	const queryStr = `SELECT PERSONA_ID FROM DOCENTE WHERE PERSONA_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)
	docente := &Docente{}
	err := result.Scan(&docente.ID)
	return docente, err

}

func (repo *repository) ObtenerPersonaPorID(ID int) (*persona.Person, error) {
	const queryStr = `SELECT PERSONA_ID, NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, 
	DNI, FECHA_NACIMIENTO FROM PERSONA WHERE PERSONA_ID = ? AND ESTADO_ELIMINADO = 1`
	result := repo.db.QueryRow(queryStr, ID)
	persona := &persona.Person{}
	err := result.Scan(&persona.ID, &persona.Nombre, &persona.ApellidoPaterno, &persona.ApellidoMaterno, &persona.Genero, &persona.DNI, &persona.FechaNacimiento)
	return persona, err
}
