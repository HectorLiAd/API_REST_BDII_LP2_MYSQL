package jerarquia

import (
	"github.com/API_REST_BDII_LP2_MYSQL/tables/sucursal"
	"github.com/API_REST_BDII_LP2_MYSQL/tables/unidadacademica"
)

/*Jerarquia estructura la cual nos sirve para que se comporte como un json o que haga el mapeo*/
type Jerarquia struct {
	ID                   int                              `json:"id"`
	UnidadAcademica      *unidadacademica.UnidadAcademica `json:"unidadAcademica"`
	UnidadacademicaID    int                              `json:"unidadAcademicaID,omitempty"`
	Sucursal             *sucursal.Sucursal               `json:"sucursal"`
	SucursalID           int                              `json:"sucursalID,omitempty"`
	Jerarquia            []*Jerarquia                     `json:"jerarquiaHija,omitempty"`
	JerarquiaID          int                              `json:"jerarquiID,omitempty"`
	TotaJerarquiaslHijas int                              `json:"totalJerarquiaHijas"`
}
