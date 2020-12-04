package database

import (
	"database/sql"
	"fmt"
)

/*InitDB Permite hacer la conexion a la BD oracle*/
func InitDB() *sql.DB {
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
