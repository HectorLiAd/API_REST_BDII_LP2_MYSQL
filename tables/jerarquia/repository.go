package jerarquia

import (
	"database/sql"
	"errors"
	"fmt"
)

/*Repository para llamar manilpular la BD*/
type Repository interface {
	RegistrarJerarquia(params *addJerarquiaRequest) (int, int, error)
	AgregarJerarquiaPadre(params *addJerarquiParentRequest) (int, error)
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
