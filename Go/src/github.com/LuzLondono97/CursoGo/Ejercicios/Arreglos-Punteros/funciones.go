package main

import "fmt"

// permiten ser modificada su lógica respetando la cantidad de parametros y los return

// las funciones también son tipos de datos
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//sumar
	fmt.Printf("Suma 5 + 7 = %d \n", Calculo(5, 7))
	//resta
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Suma 6 + 4 = %d \n", Calculo(6, 4))

	Operaciones()

	/*clousures*/
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla)

	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

/* CLOSURES: Para proteger e isolar el código o variables
pueden acceder a variables fuera de la función sin ser variables globales*/

// esta función retorna una función
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		//res := numero * secuencia
		return secuencia
	}
}
