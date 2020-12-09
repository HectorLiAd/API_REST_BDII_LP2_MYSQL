package helper

import (
	"time"
)

/*ConvStrADate convierte string a date*/
func ConvStrADate(date string) (time.Time, error) {
	dateTimeOne, errS := time.Parse("2006/01/02", date)
	dateTimeTwo, errE := time.Parse("2006-01-02", date)
	if errS != nil {
		return dateTimeOne, errS
	}

	if errE != nil {
		return dateTimeTwo, errE
	}
	return dateTimeTwo, nil
}
