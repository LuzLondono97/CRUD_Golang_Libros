package main

import "fmt"

func main() {

	fmt.Println("Paso 1")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Paso 2")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Paso 3")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j > 10 {
			break
		}
	}
}
