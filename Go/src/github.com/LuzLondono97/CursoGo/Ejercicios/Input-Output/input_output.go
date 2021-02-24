package main

import (
	"bufio"
	"fmt"
	"os"

)

func main() {
	/*var edad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es %d\n", edad)*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}
