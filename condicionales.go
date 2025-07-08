package main

import "fmt"

func condicionales() {
	var edad int = 21
	var permiso bool = false

	if edad < 18 && permiso {
		fmt.Println("El menor de edad puede viajar solo")
	} else if edad < 18 && !permiso {
		fmt.Println("El menor de edad no puede viajar solo")
	} else {
		fmt.Println("La persona puede viajar solo")
	}

	println("Fin del programa")
}
