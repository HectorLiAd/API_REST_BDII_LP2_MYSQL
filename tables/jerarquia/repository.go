package jerarquia

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/API_REST_BDII_LP2_MYSQL/tables/sucursal"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/unidadacademica"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarJerarquia(params *addJerarquiaRequest) (int, int, error)
	AgregarJerarquiaPadre(params *addJerarquiParentRequest) (int, error)
	ObtenerJerarquiaPorID(param *getJerarquiaByIDRequest) (*Jerarquia, error)
	ObtenerUnidadAcademicaPorID(ID int) (*unidadacademica.UnidadAcademica, error)
	ObtenerSucursalPorID(ID int) (*sucursal.Sucursal, error)
	ObtenerJerarquiaIDsHijos(ID int) ([]int, error)
	TotalJerarquiaHijas(ID int) (int, error)
	ObtenerTodasLasJerarquias() ([]int, error)
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

func (repo *repository) RegistrarJerarquia(params *addJerarquiaRequest) (int, int, error) {
	const queryStr = `INSERT INTO JERARQUIA(UNIDAD_ACAD_ID, SUCURSAL_ID) VALUES(?, ?)`
	result, err := repo.db.Exec(queryStr, params.UnidadAcadID, params.SucursalID)
	if err != nil {
		return -1, -1, err
	}
	jerarquiaID, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	rowAffected, err := result.RowsAffected()
	return int(jerarquiaID), int(rowAffected), err
}

func (repo *repository) AgregarJerarquiaPadre(params *addJerarquiParentRequest) (int, error) {
	const queryConst = `UPDATE JERARQUIA SET PARENT_JERARQUIA_ID = ? WHERE JERARQUIA_ID = ?`
	result, err := repo.db.Exec(queryConst, params.JerarquiaParentID, params.JerarquiaID)
	if err != nil {
		return -1, errors.New(fmt.Sprint("Error al registrar en la BD ", err))
	}
	rowAffected, err := result.RowsAffected()
	return int(rowAffected), err
}

func (repo *repository) ObtenerJerarquiaPorID(param *getJerarquiaByIDRequest) (*Jerarquia, error) {
	const queryStr = `SELECT * FROM VW_JERARQUIA WHERE JERARQUIA_ID = ?`
	jerarquia := &Jerarquia{}
	result := repo.db.QueryRow(queryStr, param.ID)
	err := result.Scan(&jerarquia.ID, &jerarquia.UnidadacademicaID, &jerarquia.SucursalID, &jerarquia.JerarquiaID)
	return jerarquia, err
}

func (repo *repository) ObtenerSucursalPorID(ID int) (*sucursal.Sucursal, error) {
	const queryStr = `SELECT SUCURSAL_ID, NOMBRE, DIRECCION, DESCRIPCION FROM SUCURSAL WHERE SUCURSAL_ID = ?`
	result := repo.db.QueryRow(queryStr, ID)
	sucursal := &sucursal.Sucursal{}
	err := result.Scan(&sucursal.ID, &sucursal.Nombre,
		&sucursal.Direccion, &sucursal.Descripcion)
	return sucursal, err
}

func (repo *repository) ObtenerUnidadAcademicaPorID(ID int) (*unidadacademica.UnidadAcademica, error) {
	const querySrt = `SELECT UNIDAD_ACAD_ID, TIPO_UNIDAD, NOMBRE FROM VW_UNIDAD_ACADEMICA WHERE UNIDAD_ACAD_ID = ?`
	result := repo.db.QueryRow(querySrt, ID)
	unidadAcademica := &unidadacademica.UnidadAcademica{}
	err := result.Scan(&unidadAcademica.ID, &unidadAcademica.TipoUnidad, &unidadAcademica.Nombre)
	return unidadAcademica, err
}

func (repo *repository) ObtenerJerarquiaIDsHijos(ID int) ([]int, error) {
	const queryStr = `SELECT JERARQUIA_ID FROM VW_JERARQUIA WHERE PARENT_JERARQUIA_ID = ?`
	result, err := repo.db.Query(queryStr, ID)
	if err != nil {
		return nil, err
	}
	var jerarquiaHijosIDs []int
	for result.Next() {
		var jerarquiaID int
		err := result.Scan(&jerarquiaID)
		if err != nil {
			return nil, err
		}
		jerarquiaHijosIDs = append(jerarquiaHijosIDs, jerarquiaID)
	}
	return jerarquiaHijosIDs, nil
}

func (repo *repository) TotalJerarquiaHijas(ID int) (int, error) {
	const queryStr = `SELECT COUNT(*) FROM VW_JERARQUIA WHERE PARENT_JERARQUIA_ID = ?`
	result := repo.db.QueryRow(queryStr, ID)
	var totalJerarHijas int
	err := result.Scan(&totalJerarHijas)
	// fmt.Println(fmt.Sprint(ID, " tiene ", totalJerarHijas))
	return totalJerarHijas, err
}

func (repo *repository) ObtenerTodasLasJerarquias() ([]int, error) {
	const queryStr = `SELECT JERARQUIA_ID FROM VW_JERARQUIA`
	result, err := repo.db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	var jerarquiasIDs []int
	for result.Next() {
		var jerarquiaID int
		err := result.Scan(&jerarquiaID)
		if err != nil {
			return nil, err
		}
		jerarquiasIDs = append(jerarquiasIDs, jerarquiaID)
	}
	return jerarquiasIDs, nil
}
