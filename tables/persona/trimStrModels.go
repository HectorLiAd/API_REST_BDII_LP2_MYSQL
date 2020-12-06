package persona

import (
	"strings"
)

func trimStrAddPersonRequest(params *addPersonRequest) *addPersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPaterno = strings.TrimSpace(params.ApellidoPaterno)
	params.ApellidoMaterno = strings.TrimSpace(params.ApellidoMaterno)
	params.Genero = strings.TrimSpace(params.Genero)
	params.Dni = strings.TrimSpace(params.Dni)
	params.FechaNacimiento = strings.TrimSpace(params.FechaNacimiento)
	return params
}

func trimStrUpdatePersonRequest(params *updatePersonRequest) *updatePersonRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	params.ApellidoPaterno = strings.TrimSpace(params.ApellidoPaterno)
	params.ApellidoMaterno = strings.TrimSpace(params.ApellidoMaterno)
	params.Genero = strings.TrimSpace(params.Genero)
	params.Dni = strings.TrimSpace(params.Dni)
	params.FechaNacimiento = strings.TrimSpace(params.FechaNacimiento)
	return params
}
