package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //driver para postgres

)

// getConnection => obtiene la conexion a la BD
func GetConnection() *sql.DB {
	//postgres://user:password@localhost:puerto/name_bd ?sslmode=disable
	dsn := "postgres://postgres:12345678@localhost:5432/biblioteca?sslmode=disable"
	db, err := sql.Open("postgres", dsn) //bd y cadena de conexion
	if err != nil {
		log.Fatal(err)
	}
	return db
}
