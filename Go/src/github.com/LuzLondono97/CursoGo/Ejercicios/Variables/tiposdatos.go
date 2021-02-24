package main

import (
	"fmt"
	"strconv"

)

func main() {
	//Int to String
	//edad := 23
	//edadtext := strconv.Itoa(edad)
	//fmt.Println("Mi edad es: " + edadtext)

	//String to Int
	edadStr := "23"
	edadInt, _ := strconv.Atoi(edadStr)
	fmt.Println(edadInt + 10)

}
