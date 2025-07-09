package main

import "fmt"

func main() {
	fmt.Println("Bienvenidos a mi calculadora en Golang")
	var operador string
	fmt.Println("Escoge un operador por favor: +, -, *, /")
	fmt.Scan(&operador)

	var numero1, numero2 int
	fmt.Println("Ingresa tus numeros")
	fmt.Scan(&numero1, &numero2)

	switch operador {
	case "+":
		fmt.Println(numero1 + numero2)
	case "-":
		fmt.Println(numero1 - numero2)
	case "*":
		fmt.Println(numero1 * numero2)
	case "/":
		fmt.Println(numero1 / numero2)
	default:
		fmt.Println("Operador ingresado invalido")
	}
}
