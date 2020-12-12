package database

import (
	"database/sql"
	"fmt"
)

/*InitDB Permite hacer la conexion a la BD oracle*/
func InitDB() *sql.DB {
	// connectionString := "pingupingu:scadSC%f12&@tcp(localhost:3306)/bd-name"
	connectionString := "uc41fkhz0c5cbztf:SybsmMbVqUj6JniRcmB@tcp(b1w8f7hqdbnb3gy4js1l-mysql.services.clever-cloud.com:20436)/b1w8f7hqdbnb3gy4js1l"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
