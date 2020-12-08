package database

import (
	"database/sql"
	"fmt"
)

/*InitDB Permite hacer la conexion a la BD oracle*/
func InitDB() *sql.DB {
	// connectionString := "pingupingu:scadSC%f12&@tcp(localhost:3306)/bd-name"
	connectionString := "u9wqag4h88nf6eng:9k7s3xgPDtGXPRYbWajd@tcp(b9i80o4lzlbmwrxkujkq-mysql.services.clever-cloud.com:3306)/b9i80o4lzlbmwrxkujkq"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
