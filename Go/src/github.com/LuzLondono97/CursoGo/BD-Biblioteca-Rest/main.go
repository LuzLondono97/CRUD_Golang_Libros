package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetLibrosEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	libros, err := LibrosConsultar()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(respuesta).Encode(libros)
}

func CreateLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var newlibro Libro
	_ = json.NewDecoder(req.Body).Decode(&newlibro)
	idInt, error := strconv.ParseInt(params["id"], 10, 64)
	if error != nil {
		fmt.Println("Hubo un error en el parseInt")
		log.Fatal(error)
	}
	newlibro.Codigo = idInt

	err := LibroCrear(newlibro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado Exitosamente!")
}

func UpdateLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	libros, err := LibrosConsultar()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range libros {
		idInt, _ := strconv.ParseInt(params["id"], 10, 64)
		if item.Codigo == idInt {
			if err := LibroActualizar(item); err != nil {
				fmt.Println(err)
			}
			break
		}
	}
	fmt.Println("Modificado Correctamente!")
	json.NewEncoder(respuesta).Encode(libros)
}

func DeleteLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	libros, err := LibrosConsultar()
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range libros {
		idInt, _ := strconv.ParseInt(params["id"], 10, 64)
		if item.Codigo == idInt {
			if err := LibroBorrar(idInt); err != nil {
				fmt.Println(err)
			}
			break
		}
	}
	json.NewEncoder(respuesta).Encode(libros)
}

func main() {
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/libros", GetLibrosEndpoint).Methods("GET")
	router.HandleFunc("/libro/{id}", CreateLibroEndpoint).Methods("POST")
	router.HandleFunc("/libro/{id}", UpdateLibroEndpoint).Methods("PUT")
	router.HandleFunc("/libro/{id}", DeleteLibroEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router)) //En caso de que ocurra un error
}
