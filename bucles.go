package main

import "fmt"

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
		if fruta == "melon" {
			fmt.Println("Excelete as accedido a la fruta naranja")
			break
		}
		fmt.Println("Indica tu fruta de preferencia")
		fmt.Scan(&fruta)
	}
}
