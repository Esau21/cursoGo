package main

import "fmt"

func mapas() {
	miMapa := map[string]string{
		"EL SALVADOR": "SAN SALVADOR",
		"COLOMBIA":    "BOGOTA",
		"VENEZUELA":   "CARACAS",
		"MEXICO":      "CIUDAD DE MEXICO",
	}

	fmt.Println("Mi mapa de paises es, ", miMapa)

	fmt.Println("La capital de venezuela es: ", miMapa["VENEZUELA"])

	miMapa["ARGENTINA"] = "BUENOS AIRES"

	fmt.Println("La capital de argentina es:", miMapa["ARGENTINA"])

	delete(miMapa, "ARGENTINA")

	fmt.Println("El mapa de paises es: ", miMapa)

}
