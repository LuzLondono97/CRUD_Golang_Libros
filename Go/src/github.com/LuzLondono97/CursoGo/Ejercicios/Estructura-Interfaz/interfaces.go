package main

import "fmt"

type User interface {
	Permisos() int //1-5
	Nombre() string
}

//Usuario Admin
type Admin struct {
	nombre string
}

func (admin Admin) Permisos() int {
	return 5
}

func (admin Admin) Nombre() string {
	return admin.nombre
}

//Usuario Editor
type Editor struct {
	nombre string
}

func (editor Editor) Permisos() int {
	return 3
}

func (editor Editor) Nombre() string {
	return editor.nombre
}

//Metodo para validar el usuario
func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " Tiene permisos de administrador"
	}
	return user.Nombre() + " No tiene permisos"
}

func main() {

	fmt.Println("Paso 1")
	//Usuario Admin
	admin := Admin{"Uriel"}
	fmt.Println(auth(admin))

	//Usuario Editor
	editor := Editor{"Mary"}
	fmt.Println(auth(editor))

	fmt.Println("Paso 2")
	usuarios := []User{Admin{"Andrea"}, Editor{"Daniel"}}
	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
}
