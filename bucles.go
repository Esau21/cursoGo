package main

import (
	"fmt"
	"strings"
)

func main() {
	/* var suma int = 0
	for i := 0; i <= 100; i++ {

		if i%2 != 0 {
			suma = suma + i
		}
	}

	fmt.Println(suma) */
	/*
		miMapa := map[string]string{
			"EL SALVADOR": "SAN SALVADOR",
			"COLOMBIA":    "BOGOTA",
			"VENEZUELA":   "CARACAS",
			"MEXICO":      "CIUDAD DE MEXICO",
		}

		for k, v := range miMapa {
			fmt.Println("La captital de " + k + " es " + v)
		} */

	var fruta string = ""

	for {
		fmt.Println("Indica tu fruta esperada")
		fmt.Scan(&fruta)

		fruta = strings.ToLower(fruta)

		if fruta == "melon" {
			fmt.Println("Excelete as accedido a la fruta: " + fruta)
			break
		} else {
			fmt.Println("No es la fruta correcta")
		}
	}
}
