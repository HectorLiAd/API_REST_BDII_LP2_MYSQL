package persona

import (
	"strings"
)

func trimStrAddPersonRequest(params *addPersonRequest) *addPersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPat = strings.TrimSpace(params.ApellidoPat)
	params.ApellidoMat = strings.TrimSpace(params.ApellidoMat)
	params.Genero = strings.ToUpper(strings.TrimSpace(params.Genero))
	params.DNI = strings.TrimSpace(params.DNI)
	params.FechaNac = strings.TrimSpace(params.FechaNac)
	return params
}

func trimStrUpdatePersonRequest(params *updatePersonRequest) *updatePersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPat = strings.TrimSpace(params.ApellidoPat)
	params.ApellidoMat = strings.TrimSpace(params.ApellidoMat)
	params.Genero = strings.ToUpper(strings.TrimSpace(params.Genero))
	params.DNI = strings.TrimSpace(params.DNI)
	params.FechaNac = strings.TrimSpace(params.FechaNac)
	return params
}
