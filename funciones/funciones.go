package main

import "fmt"

/* Funciones para Golang */

func presentarResulados(f func(float64, float64) float64,
	a float64, b float64) {
	/* fmt.Println("E resultado de ", a, " y ", b, "es ", suma(a, b)) */
	fmt.Println("Para a = ", a, " y b = ", b, " La suma resultante es ", f(a, b))
}

func presentarResuladosResta(f func(float64, float64) float64,
	a float64, b float64) {
	fmt.Println("Para a = ", a, " y b = ", b, " La resta resultante es ", f(a, b))
}

func main() {

	suma := func(a float64, b float64) float64 {
		return a + b
	}

	resta := func(a float64, b float64) float64 {
		return a - b
	}

	presentarResulados(suma, 7.20, 6.90)
	presentarResuladosResta(resta, 34, 20)
}
