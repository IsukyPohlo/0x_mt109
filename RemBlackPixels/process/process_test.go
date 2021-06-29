package process

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {

	const T = true
	const O = false

	/*

			//Test conv a bools

			convIn := [][]uint8{
				{1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 1, 1, 1},
			}

			convExp := [][]bool{
				{T, T, T, T, T},
				{T, O, T, T, T},
				{T, T, T, T, T},
				{T, T, T, O, T},
				{T, T, T, T, T},
			}

			convOut := boolCon(&convIn)

			for i, pxRow := range convOut {
				for j, pxBool := range pxRow {
					if pxBool != convExp[i][j] {
						t.Errorf("error en la conversion en: %d-%d", i, j)
					}
				}
			}

			//Generar slice de procesos

			procIn := [][]uint8{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1},
			}

			procExp := [][]bool{
				{O, O, O, O, O, O, O},
				{O, O, O, O, O, O, O},
				{O, O, O, O, O, O, O},
				{O, O, O, O, O, O, O},
				{O, O, O, O, O, O, O},
			}

			procOut, procLen, procWid := genProcSlice(&procIn)

			for i, procRow := range procOut {
				for j, procBool := range procRow {
					if procBool != procExp[i][j] {
						t.Errorf("error en la generacion en: %d-%d", i, j)
					}
				}
			}

			if (len(procOut[0])) != procLen {
				t.Errorf("error en la dimension en y")
			}

			if (len(procOut)) != procWid {
				t.Errorf("error en la dimension en x")
			}


		//Prueba de busqueda
		fmt.Println("Prueba de busqueda")
		fmt.Println()

		var testImage Image

		searchIn := [][]uint8{
			{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0},
			{1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0},
			{0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
		}

		testImage.pixelImage = boolCon(&searchIn)
		testImage.processed, testImage.length, testImage.width = genProcSlice(&searchIn)

		fmt.Println("Prueba busqueda")
		cleanImage(&testImage)

		cleanPixels(&testImage)

		for _, row := range testImage.processed {
			for _, pixel := range row {
				if pixel {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
		}

		fmt.Println()

		for _, row := range testImage.pixelImage {
			for _, pixel := range row {
				if pixel {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
		}

	*/

	//Prueba completa

	fmt.Println("Prueba completa")
	fmt.Println()

	completeIn := [][]uint8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
	}

	completeExp := [][]uint8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
	}

	completeOut := ImageProcess(&completeIn)

	fmt.Println(" Entrada :")

	for _, row := range completeIn {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}

	//Mostrar salida

	fmt.Println()
	fmt.Println("Esperado")

	for _, row := range completeExp {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}

	//Mostrar salida

	fmt.Println()
	fmt.Println("Salida")

	for _, row := range completeOut {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}

	fmt.Println("comp")

	for i := 0; i < len(completeOut); i++ {
		for j := 0; j < len(completeOut[0]); j++ {
			//fmt.Print(completeOut[i][j])
			//fmt.Print(completeExp[i][j])
			if completeOut[i][j] != completeExp[i][j] {
				t.Errorf("Resultado inesperado en %d,%d", i, j)
			}
		}
	}

	fmt.Println()

}
