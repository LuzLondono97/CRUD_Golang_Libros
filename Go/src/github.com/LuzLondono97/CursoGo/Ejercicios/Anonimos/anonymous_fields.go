package main

import "fmt"

type Human struct {
	name string
}

func (human Human) hablar() string {
	return "Bla bla bla  "
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (tutor Tutor) hablar() string {
	//return "Bienvenidos"
	return tutor.Human.hablar() + "Bienvenidos"
}

func main() {
	tutor := Tutor{Human{"Luz"}, Dummy{"Daniel"}}
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)

	fmt.Println(tutor.Human.hablar())
	//Toma el del tutor es lo mismo si fuera => fmt.Println(tutor.Tutor.hablar())
	fmt.Println(tutor.hablar())

}
