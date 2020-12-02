package helper

import (
	"regexp"
)

/*ValidarEmailStr permite validar el Email de tipo string, un tipo hector.limauya@upeu.edu.pe */
func ValidarEmailStr(email string) bool {
	if len(re.FindStringIndex(email)) > 0 {
		return true
	}
	return false
}

/*ValidarPasswordStr permite validar el password de tipo str, valida caracteres mayusculas y minusculas números y caracteres especiales*/
func ValidarPasswordStr(str string) bool {
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidarDniStr permite validar el DNI de tipo string, solo admite de tipo numerico entero */
func ValidarDniStr(str string) bool {
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidarDateStr permite validar la fecha de tipo string, solo almite los siguientes formatos de fechas YYYY/MM/DD o YYYY-MM-DD*/
func ValidarDateStr(str string) bool {
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}
