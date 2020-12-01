package tipounidad

import (
	"strings"
)

func trimStrAddTipoUnidadRequest(params *addTipoUnidadRequest) *addTipoUnidadRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.Descripcion = strings.TrimSpace(params.Descripcion)
	return params
}
