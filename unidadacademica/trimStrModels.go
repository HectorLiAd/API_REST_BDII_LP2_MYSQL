package unidadacademica

import "strings"

func trimStrAddUnidadAcademicaRequest(params *addUnidadAcademicaRequest) *addUnidadAcademicaRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	return params
}
