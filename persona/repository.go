package persona

import (
	"database/sql"
	"fmt"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	GetPersonByID(param *getPersonByIDRequest) (*Person, error)
	GetPersons(params *getPersonsRequest) ([]*Person, error)
	GetTotalPersons() (int, error)
	InsertPerson(params *addPersonRequest) (int, error)
	UpdatePerson(params *updatePersonRequest) (int, error)
	DeletePerson(param *deletePersonRequest) (int, error)
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
	const queryStr = `SELECT PERSONA_ID, NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, DNI, FECHA_NACIMIENTO FROM PERSONA WHERE PERSONA_ID = ? AND ESTADO <> 0`
	row := repo.db.QueryRow(queryStr, param.PersonaID)
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
	)
	// year, month, day := fecha_nac.Date()
	// fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)
	// persona.FechaNacimiento = fechaNac.Format("02/01/2006")
	persona.FechaNacimiento = string(fechaNac)
	return persona, err
}

func (repo *repository) GetPersons(params *getPersonsRequest) ([]*Person, error) {
	const sql = `
	SELECT PERSONA_ID, NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, DNI, FECHA_NACIMIENTO
	FROM PERSONA WHERE ESTADO <> 0 limit ? offset ?`
	result, err := repo.db.Query(sql, params.Limit, params.Offset)
	var fechaNac []uint8
	if err != nil {
		return nil, nil
	}

	var persons []*Person
	for result.Next() {
		persona := &Person{}
		err := result.Scan(
			&persona.ID,
			&persona.Nombre,
			&persona.ApellidoPaterno,
			&persona.ApellidoMaterno,
			&persona.Genero,
			&persona.Dni,
			&fechaNac,
		)
		if err != nil {
			return nil, err
		}
		persona.FechaNacimiento = string(fechaNac)
		persons = append(persons, persona)
	}
	return persons, err
}

func (repo *repository) GetTotalPersons() (int, error) {
	const queryStr = `SELECT COUNT(PERSONA_ID) FROM PERSONA WHERE ESTADO <> 0`
	var total int
	row := repo.db.QueryRow(queryStr)

	err := row.Scan(&total)
	return total, err
}

func (repo *repository) InsertPerson(params *addPersonRequest) (int, error) {
	const queryStr = `INSERT INTO PERSONA (NOMBRE, APELLIDO_P, APELLIDO_M, GENERO, DNI, FECHA_NACIMIENTO)
						VALUES (?, ?, ?, ?, ?, ?);`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.ApellidoPaterno,
		params.ApellidoMaterno, params.Genero, params.Dni,
		params.FechaNacimiento)
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return int(id), err
}

func (repo *repository) UpdatePerson(params *updatePersonRequest) (int, error) {
	const queryStr = `
		UPDATE PERSONA SET 
		NOMBRE = ?, 
		APELLIDO_P = ?, 
		APELLIDO_M = ?, 
		GENERO = ?, 
		DNI = ?, 
		FECHA_NACIMIENTO = ? 
		WHERE ESTADO = 1 AND PERSONA_ID = ?;
	`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.ApellidoPaterno,
		params.ApellidoMaterno, params.Genero, params.Dni,
		params.FechaNacimiento, params.ID)
	rowAfected, errr := result.RowsAffected()
	if errr != nil {
		return 0, errr
	}
	fmt.Println(rowAfected)
	return int(rowAfected), err
}

func (repo *repository) DeletePerson(param *deletePersonRequest) (int, error) {
	const query = `UPDATE PERSONA SET ESTADO = 0 WHERE PERSONA_ID = ? AND ESTADO <> 0`
	result, err := repo.db.Exec(query, param.PersonaID)
	rowAfected, _ := result.RowsAffected()
	return int(rowAfected), err
}
