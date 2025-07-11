package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculadora() {
	fmt.Println("Bienvenidos a mi calculadora")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Por favor escoge un operador para poder hacer la operacion matematica")
	fmt.Println("Operadores validos, +, -, *, /")
	fmt.Println("Si  quieres salir del programa solo presiona la tecla enter")

	operadorEnter, _ := reader.ReadString('\n')
	operador := strings.TrimSpace(operadorEnter)

	if operador == "" {
		fmt.Println("Saliendo del programa regresa pronto...")
		return
	}

	var numero1, numero2 int
	fmt.Println("Inserta dos numeros para poder reaizar la operacion" + operador)
	fmt.Scan(&numero1, &numero2)

	switch operador {
	case "+":
		suma := numero1 + numero2
		fmt.Printf("EL RESULTADO DE ESTA SUMA ES: %d\n", suma)
	case "-":
		resta := numero1 - numero2
		fmt.Printf("EL RESULTADO DE ESTE RESTA ES: %d\n", resta)

	case "*":
		multiplicacion := numero1 * numero2
		fmt.Printf("EL RESULTADO DE ESTA MULTIPLICACION ES: %d\n", multiplicacion)
	case "/":
		if numero2 != 0 {
			divison := float64(numero1) / float64(numero2)
			fmt.Printf("EL RESULTADO DE ESTA DIVISION ES: %.2f\n", divison)
		} else {
			fmt.Println("Error no es divisible entre 0")
		}
	default:
		fmt.Println("Error el operador que elegiste no es valido")
	}

}
