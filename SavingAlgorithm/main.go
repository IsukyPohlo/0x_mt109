package main

import (
	"log"
	"savingAlgo/readFile"
	"savingAlgo/savings"
)

func main() {

	//
	const file = "./Input/10_cm.json"

	//Leer entrada
	routesIn, err := readFile.ReadFileJson(file)
	if err != nil {
		log.Fatal(err)
	}

	//Calcular Savings y Ordenar
	saveList := savings.CalculateSaves((&routesIn))

	//Grafos
	bestRoute := savings.CalculateRoute(saveList, len(routesIn.Matrix))

	//Ruta
	savings.PrintRoute(&bestRoute)

	//Ruta con distancias
	savings.PrintRutaDist(&bestRoute, &routesIn)

}
