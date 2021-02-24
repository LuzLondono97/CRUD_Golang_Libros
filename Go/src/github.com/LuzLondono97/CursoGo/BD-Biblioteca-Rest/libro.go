package main

import (
	"errors"
)

//Estructura de Libro
type Libro struct {
	Codigo int64 `json:"id"`
	Titulo  string `json:"title"`
	Autor   string
	Editorial   string
}

//Consultar todos los libros
func LibrosConsultar() (libros []Libro) err error {
	sqlSelect := `SELECT codigo, titulo, autor, editorial
					FROM libros`

	db := GetConnection()
	defer db.Close()

	rows, err := db.Query(sqlSelect)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		l := Libro{}
		err = rows.Scan(
				&l.Codigo,
				&l.Titulo,
				&l.Autor,
				&l.Editorial )
		if err != nil {
			return nil, nil
		}

		libros = append(libros, l)
	}

	return libros, nil
}

//Crear Libro
func LibroCrear(l Libro) error {
	sqlInsert := `INSERT INTO
					libros (titulo, autor, editorial)
					VALUES ($1, $2, $3)`

	//Permitir nil de un string
	//nullNombre := sql.NullString{}

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Validar nil de un string
	/*if l.Titulo == "" {
		nullNombre.Valid = false
	} else {
		nullNombre.Valid = true
		nullNombre.String = l.Titulo
	}*/

	result, err := stmt.Exec(l.Titulo, l.Autor, l.Editorial)
	if err != nil {
		return err
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba 1 fila afectada")
	}

	return nil
}

//Modificar Libro
func LibroActualizar(l Libro) error {
	sqlUpdate := `UPDATE libros 
					SET titulo = $1, autor = $2, editorial = $3
					WHERE codigo = $4`

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(sqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(l.Titulo, l.Autor, l.Editorial, l.Codigo)
	if err != nil {
		return err
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Se esperaba 1 fila afectada")
	}

	return nil
}

//Eliminar registro del Libro
func LibroBorrar(id int64) error {
	sqlDelete := `DELETE FROM libros WHERE codigo = $1`

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}
