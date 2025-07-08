package main

import "fmt"

func main() {
	var fruta string
	fmt.Println("Ingresa valores alfa numericos")
	fmt.Scan(&fruta)

	switch fruta {
	case "manzana":
		fmt.Println("Has ingresado 'manzana'.")
	case "pera":
		fmt.Println("Has ingresado 'pera'.")
	case "banana":
		fmt.Println("Has ingresaso 'banaa'.")
	default:
		fmt.Println("No reconozco ese valor")
	}
}
