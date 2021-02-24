package main

import "fmt"

func main() {

	fmt.Println("paso 1")
	slice := make([]int, 3, 5) //
	slice = append(slice, 15)  //añade un elemento al arreglo, si rebasara la capacidad la aumentaria
	fmt.Println(slice)         //puntero del arreglo
	fmt.Println(len(slice))    //tamaño del arreglo
	fmt.Println(cap(slice))    //capacidad

	fmt.Println("paso 2")
	//Funcion Copy; copia el minimo de elementos del arreglo fuente
	//copy(destino,fuente)
	arreglo := []int{1, 2, 3, 4}
	copia := make([]int, len(arreglo), cap(arreglo)*2)

	copy(copia, arreglo)

	fmt.Println(arreglo)
	fmt.Println(copia)
}
