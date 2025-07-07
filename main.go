package main

import "fmt"

func main() {

	var miNombre string = "Edgar Esau Zelaya Moran"
	var apellidoPersona = "Ayala"
	tercerNombre := "hector"
	fmt.Println("Hola mundo mi nombre es :" + miNombre)
	fmt.Println("Mi apelido es: " + apellidoPersona)
	fmt.Println("Este es mi otro nombre desde Go :" + tercerNombre)

	/* Parte numerica para Go */

	var numero1 = 10
	var numero2 = 10
	edad := 100000000

	var sumarDesdeGo = numero1 + numero2

	fmt.Println(sumarDesdeGo)
	fmt.Println(edad)

	/* arreglos en go */
	/* si se pasa del limite que le especificamos en el array dara error por wl fuerte prototipado del leneguaje Go */
	var listafrutas = [5]string{"Pera", "Naranja", "Manzana", "Guineo", "Sandia"}
	fmt.Println(listafrutas[2])

	listPaises := []string{"Peru", "Chile", "Mexico"}
	fmt.Println(listPaises)
	//listPaises[0] = "El salvador"
	//fmt.Println(listPaises)

	listPaises = append(listPaises, "El salvador")

	fmt.Println(listPaises)

	listaPaises2 := listPaises[1:3]

	fmt.Println(listaPaises2)

	listaPaises3 := listPaises[2:]

	fmt.Println(listaPaises3)
}
