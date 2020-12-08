package sucursal

import (
	"database/sql"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	InsertarSucursal(params *addSucursalRequest) (int, int, error)
	ActualizarSucursal(params *updateSucursalRequest) (int, error)
	ObtenerTodoSucursal() ([]*Sucursal, error)
	ObtenerSucursalPorID(param *getSucursalByIDRequest) (*Sucursal, error)
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

func (repo *repository) InsertarSucursal(params *addSucursalRequest) (int, int, error) {
	queryStr := `INSERT INTO SUCURSAL (NOMBRE, DIRECCION, DESCRIPCION) VALUES (?, ?, ?)`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Direccion, params.Descripcion)
	if err != nil {
		return 0, 0, err
	}
	idSucursal, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(idSucursal), int(rowAffected), err
}

func (repo *repository) ActualizarSucursal(params *updateSucursalRequest) (int, error) {
	const queryStr = `UPDATE SUCURSAL SET NOMBRE = ?, DIRECCION = ?, DESCRIPCION = ? WHERE SUCURSAL_ID = ?`
	result, err := repo.db.Exec(queryStr, params.Nombre, params.Direccion, params.Descripcion, params.ID)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerTodoSucursal() ([]*Sucursal, error) {
	const queryStr = `SELECT SUCURSAL_ID, NOMBRE, DIRECCION, DESCRIPCION FROM SUCURSAL`
	resultRow, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var sucursales []*Sucursal
	for resultRow.Next() {
		sucursal := &Sucursal{}
		err := resultRow.Scan(&sucursal.ID, &sucursal.Nombre, &sucursal.Direccion, &sucursal.Descripcion)
		if err != nil {
			return nil, err
		}
		sucursales = append(sucursales, sucursal)
	}
	return sucursales, nil
}

func (repo *repository) ObtenerSucursalPorID(param *getSucursalByIDRequest) (*Sucursal, error) {
	const queryStr = `SELECT SUCURSAL_ID, NOMBRE, DIRECCION, DESCRIPCION FROM SUCURSAL WHERE SUCURSAL_ID = ?`
	result := repo.db.QueryRow(queryStr, param.ID)
	sucursal := &Sucursal{}
	err := result.Scan(&sucursal.ID, &sucursal.Nombre,
		&sucursal.Direccion, &sucursal.Descripcion)
	return sucursal, err
}
