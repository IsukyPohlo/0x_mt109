package main

import (
	"RemBlackPixels/process"
	"fmt"
)

func main() {

	//Vars

	pixelInput := [][]uint8{
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	//Obtener dimensiones
	length := len(pixelInput)
	width := len(pixelInput[0])
	fmt.Println("dimensiones: ", length, " - ", width)

	//Mostrar Input

	for _, row := range pixelInput {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//Entrada de la imagen

	pixelOutput := process.ImageProcess(&pixelInput)

	//Mostrar salida

	fmt.Println()
	fmt.Println("Resultado final: ")

	for _, row := range pixelOutput {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}

}
