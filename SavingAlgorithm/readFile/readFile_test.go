package readFile

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {

	rutas, err := ReadFileJson("./Input/10_cm.json")
	if err != nil {
		t.Errorf("Error al leer archivo")
	}

	//fmt.Println(rutas)

	for j, row := range rutas.Matrix {
		for i, val := range row {
			fmt.Print("[", i, j, "]", val.DistanceInMeters)
		}
		fmt.Println()
	}

}
