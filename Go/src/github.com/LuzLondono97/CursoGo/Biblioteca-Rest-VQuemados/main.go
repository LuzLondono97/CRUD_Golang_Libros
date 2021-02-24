package main //SERVICIO REST con valores quemados

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Libro struct {
	IDlibro string `json:"id,omitempty"`
	Nombre  string
	Autor   string
}

var libros []Libro

func GetLibrosEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	json.NewEncoder(respuesta).Encode(libros)
}

func GetLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range libros {
		if item.IDlibro == params["id"] {
			json.NewEncoder(respuesta).Encode(item)
			return
		}
	}
	json.NewEncoder(respuesta).Encode(&Libro{})
}

func CreateLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var newlibro Libro
	_ = json.NewDecoder(req.Body).Decode(&newlibro)
	newlibro.IDlibro = params["id"]
	libros = append(libros, newlibro)
	json.NewEncoder(respuesta).Encode(libros)
}

func DeleteLibroEndpoint(respuesta http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range libros {
		if item.IDlibro == params["id"] {
			libros = append(libros[:index], libros[index+1:]...)
			break
		}
	}
	json.NewEncoder(respuesta).Encode(libros)
}

func main() {
	router := mux.NewRouter()

	libros = append(libros, Libro{IDlibro: "1", Nombre: "Cien años de soledad", Autor: "Miguel García Marquéz"})
	libros = append(libros, Libro{IDlibro: "2", Nombre: "El perfume", Autor: "Patrick Süskind"})

	//endpoints
	router.HandleFunc("/libros", GetLibrosEndpoint).Methods("GET")
	router.HandleFunc("/libro/{id}", GetLibroEndpoint).Methods("GET")
	router.HandleFunc("/libro/{id}", CreateLibroEndpoint).Methods("POST")
	router.HandleFunc("/libro/{id}", DeleteLibroEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router)) //En caso de que ocurra un error
}
