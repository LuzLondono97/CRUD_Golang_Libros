package main

import "fmt"

func main() {
	/*
		== igual a
		!= diferente
		< menor que
		> mayor que
		<= menor o igual que
		>= mayor o igual que
		&& AND
		|| OR
	*/

	x := 10
	y := 10

	if x > y {
		fmt.Printf("%d es mayor que %d \n", x, y)
	} else if y > x {
		fmt.Printf("%d es mayor que %d \n", y, x)
	} else {
		fmt.Println("Son Iguales")
	}

}
