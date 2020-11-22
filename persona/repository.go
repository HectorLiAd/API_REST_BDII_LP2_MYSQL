package persona

import (
	"database/sql"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	GetPersonByID(param *getPersonByIDRequest) (*Person, error)
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

func (repo *repository) GetPersonByID(param *getPersonByIDRequest) (*Person, error) {
	const sql = `SELECT * FROM PERSONA WHERE PERSONA_ID = ? AND ESTADO <> 0`
	row := repo.db.QueryRow(sql, param.PersonaID)
	var fechaNac []uint8
	persona := &Person{}
	err := row.Scan(
		&persona.ID,
		&persona.Nombre,
		&persona.ApellidoPaterno,
		&persona.ApellidoMaterno,
		&persona.Genero,
		&persona.Dni,
		&fechaNac,
		&persona.Estado,
	)
	// year, month, day := fecha_nac.Date()
	// fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)
	// persona.FechaNacimiento = fechaNac.Format("02/01/2006")
	persona.FechaNacimiento = string(fechaNac)
	return persona, err
}
