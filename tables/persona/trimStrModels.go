package persona

import (
	"strings"
)

func trimStrAddPersonRequest(params *addPersonRequest) *addPersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPat = strings.TrimSpace(params.ApellidoPat)
	params.ApellidoMat = strings.TrimSpace(params.ApellidoMat)
	params.Genero = strings.ToUpper(strings.TrimSpace(params.Genero))
	params.Dni = strings.TrimSpace(params.Dni)
	params.FechaNac = strings.TrimSpace(params.FechaNac)
	return params
}

func trimStrUpdatePersonRequest(params *updatePersonRequest) *updatePersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPaterno = strings.TrimSpace(params.ApellidoPaterno)
	params.ApellidoMaterno = strings.TrimSpace(params.ApellidoMaterno)
	params.Genero = strings.ToUpper(strings.TrimSpace(params.Genero))
	params.Dni = strings.TrimSpace(params.Dni)
	params.FechaNacimiento = strings.TrimSpace(params.FechaNacimiento)
	return params
}
