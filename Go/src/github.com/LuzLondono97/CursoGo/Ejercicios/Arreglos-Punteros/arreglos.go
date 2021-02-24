package main

import "fmt"

func main() {
	//var arreglo [3]int
	//arreglo := [3]int{1, 2, 3}

	arreglo := [3]int{1, 2}

	arreglo[2] = 20

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}

	var matriz [3][3]int
	matriz[0][1] = 10
	fmt.Println(matriz)
}
