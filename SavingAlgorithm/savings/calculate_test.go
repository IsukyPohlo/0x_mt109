package savings

import (
	"savingAlgo/readFile"
	"testing"
)

func TestCalc(t *testing.T) {

	//Obtener entrada de rutas
	rutas, err := readFile.ReadFileJson("../readFile/Input/1000_cm.json")
	if err != nil {
		t.Error("Error al abrir archivo")
	}

	//Calcula y ordena (desc) Saves
	saveList := CalculateSaves((&rutas))

	//fmt.Println("Lista ordenada")
	//fmt.Println(saveList)

	//Grafos
	bestRoute := CalculateRoute(saveList, len(rutas.Matrix))

	//Ruta
	//PrintRoute(&bestRoute)

	//Ruta con distancias
	PrintRutaDist(&bestRoute, &rutas)

	//fmt.Print((rutas.Matrix[prin][rutas[prin]].DistanceInMeters)

}
