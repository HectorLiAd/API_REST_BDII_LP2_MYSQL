package rol

import (
	"strings"
)

func trimStrAddRolRequest(param *addRolRequest) *addRolRequest {
	param.Nombre = strings.TrimSpace(param.Nombre)
	return param
}

func trimStrUpdateRolRequest(params *updateRolRequest) *updateRolRequest {
	params.Nombre = strings.TrimSpace(params.Nombre)
	return params
}
