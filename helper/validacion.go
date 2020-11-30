package helper

import (
	"regexp"
)

/*ValidarEmailStr permite validar el Email de tipo string, un tipo hector.limauya@upeu.edu.pe */
func ValidarEmailStr(email string) bool {
	var re = regexp.MustCompile(`^(?:[a-zA-Z0-9])([-_0-9a-zA-Z]+(\.[-_0-9a-zA-Z]+)*|^\"([\001-\010\013\014\016-\037!#-\[\]-\177]|\\[\001-011\013\014\016-\177])*\")@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}\.?$`)
	if len(re.FindStringIndex(email)) > 0 {
		return true
	}
	return false
}

/*ValidarPasswordStr permite validar el password de tipo str, valida caracteres mayusculas y minusculas números y caracteres especiales*/
func ValidarPasswordStr(str string) bool {
	var re = regexp.MustCompile(`[^A-Za-z0-9]`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidarDniStr permite validar el DNI de tipo string, solo admite de tipo numerico entero */
func ValidarDniStr(str string) bool {
	var re = regexp.MustCompile(`^[0-9]+$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidarDateStr permite validar la fecha de tipo string, solo almite los siguientes formatos de fechas YYYY/MM/DD o YYYY-MM-DD*/
func ValidarDateStr(str string) bool {
	var re = regexp.MustCompile(`^\d{4}[\-\/\s]?((((0[13578])|(1[02]))[\-\/\s]?(([0-2][0-9])|(3[01])))|(((0[469])|(11))[\-\/\s]?(([0-2][0-9])|(30)))|(02[\-\/\s]?[0-2][0-9]))$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}
