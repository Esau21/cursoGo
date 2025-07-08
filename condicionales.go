package main

import "fmt"

func main() {
	var edad int = 66

	if edad >= 65 {
		fmt.Println("La persona es jubilada")
	} else if edad >= 18 {
		fmt.Println("La persona es mayor edad")
	} else {
		fmt.Println("La persona es menor de edad")
	}

	println("Fin del programa")
}
