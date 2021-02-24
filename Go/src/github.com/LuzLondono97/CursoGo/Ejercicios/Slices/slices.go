package main

import "fmt"

func main() {
	fmt.Println("Paso 1")
	//var arregloSlice []int
	//arregloSlice := []int{1, 2, 3, 4}
	arregloSlice := []int{1, 2, 3, 4, 5}

	if arregloSlice == nil {
		fmt.Println("Esta vacio")
	} else {
		fmt.Println(arregloSlice)
	}

	fmt.Println(len(arregloSlice))
	//puntero al arreglo
	//longitud del arreglo
	//capacidad

	fmt.Println("Paso 2")
	arreglo := [3]int{1, 2, 3}
	slice := arreglo[:3]
	//slice := arreglo[1:3]
	//slice := arreglo[1:]
	fmt.Println(slice)
}
