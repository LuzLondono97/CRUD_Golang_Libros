package main

import "fmt"

type User struct {
	edad               int
	nombres, apellidos string
}

func (usuario User) nombreCompleto() string {
	return usuario.nombres + " " + usuario.apellidos
}

func (usuario *User) setName(nombre string) {
	usuario.nombres = nombre
}

func main() {

	fmt.Println("Paso 1")
	uriel := User{nombres: "Uriel", apellidos: "Hernandez"}
	luz := User{23, "Luz Andrea", "Londoño"}
	fmt.Println(uriel)
	fmt.Println(luz)

	fmt.Println("Paso 2")
	punteroUser := new(User)
	punteroUser.nombres = "Mary"
	punteroUser.apellidos = "Lopez"
	//fmt.Println((*punteroUser).nombres)
	fmt.Println(punteroUser.nombreCompleto())

	fmt.Println("Paso 3")
	punteroUser.setName("Marcos")
	fmt.Println(punteroUser.nombres)
}
