package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetWriteDBConexion() *sql.DB {

	serverName := os.Getenv("DBCATALOGO_SERVER_WRITE")
	username := os.Getenv("DBCATALOGO_USER_WRITE")
	password := os.Getenv("DBCATALOGO_PASSWORD_WRITE")
	dbName := os.Getenv("DBCATALOGO_NAME_WRITE")

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=1",
		username, password, serverName, dbName)

	db, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db

}
