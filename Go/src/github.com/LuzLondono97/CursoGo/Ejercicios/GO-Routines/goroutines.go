package main

import (
	"fmt"
	"strings"
	"time"

)

func main() {
	go miNombreLentooo("andrea")
	fmt.Println("Queeee Aburridoooo")
	var wait string
	fmt.Scanln(&wait)
}

func miNombreLentooo(name string) {
	letras := strings.Split(name, "") //cuando el segundo parametro es vacio particiona toda la cadena

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //Retraza el proceso 1 segundo
		fmt.Println(letra)
	}
}
