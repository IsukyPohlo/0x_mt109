package process

import (
	"fmt"
)

// ---- CONST -----
var T = true
var O = false

// ------ STRUCTS -----

type Image struct {
	pixelImage [][]bool
	processed  [][]bool
	length     int
	width      int
}

type WayStrt struct {
	curr Coord
}

type Coord struct {
	X, Y int
}

// ---- FUNCIONES -----

// Procesamiento de la imagen
func ImageProcess(preImage *[][]uint8) [][]uint8 {

	var temImage Image

	//convertir a bools

	//Generar slice de procesados
	temImage.processed, temImage.length, temImage.width = genProcSlice(preImage)

	//Recepcion
	temImage.pixelImage = boolCon(preImage)

	//Generar mascara
	cleanImage(&temImage)

	//Quitar pixeles sobrantes
	cleanPixels(&temImage)

	return uIntCon(&temImage.pixelImage)

}

// - Convertir a bools?

func boolCon(preImage *[][]uint8) [][]bool {
	fmt.Println("Conversion")

	var rowBool []bool
	var allBool [][]bool

	//Conversion a bool

	rowBool = nil

	for _, row := range *preImage {

		rowBool = nil

		for _, pixel := range row {

			if pixel == 1 {
				rowBool = append(rowBool, true)
			} else {
				rowBool = append(rowBool, false)
			}

		}
		allBool = append(allBool, rowBool)

	}

	return allBool
}

// - Array de procesamiento:   [0 = Sin procesar, 1 = Procesado]
func genProcSlice(preImage *[][]uint8) ([][]bool, int, int) {

	width := len(*preImage)
	length := len((*preImage)[0])

	procSlice := make([][]bool, width)
	for i := range procSlice {
		procSlice[i] = make([]bool, length)
	}

	return procSlice, length, width
}

// - Buscar pixeles del borde, evitar esquinas

// 		- Ejecutar Busqueda en los bordes
func cleanImage(i *Image) {

	w := WayStrt{
		curr: Coord{0, 0},
	}

	oX := 0
	oY := 0

	//Determinar 4 bordes

	(*i).width = len((*i).pixelImage)
	(*i).length = len((*i).pixelImage[0])

	//Borde Arriba:
	//fmt.Println("ARRIBA")

	for oX = 1; oX < (*i).length-1; oX++ {

		w.curr.Y = 0
		w.curr.X = oX

		//fmt.Println("(Y:", w.curr.Y, "X: ", w.curr.X, ")")

		searchColl(w, i)

	}

	//Borde Abajo:
	//fmt.Println("ABAJO")
	for oX = 1; oX < (*i).length-1; oX++ {

		w.curr.X = oX
		w.curr.Y = (*i).width - 1

		//Procesar orillas
		//(*i).processed[w.curr.Y][w.curr.X] = true
		//fmt.Println("Bor :", w.curr.Y, w.curr.X)

		searchColl(w, i)
	}

	//Borde Izquierda
	//fmt.Println("IZQ")

	for oY = 1; oY < (*i).width-1; oY++ {

		w.curr.Y = oY
		w.curr.X = 0

		//Procesar orillas
		//(*i).processed[w.curr.Y][w.curr.X] = true
		//fmt.Println("Bor :", w.curr.Y, w.curr.X)

		searchColl(w, i)
	}

	//Borde Derecha
	//fmt.Println("DER")
	for oY = 1; oY < (*i).width-1; oY++ {

		w.curr.Y = oY
		w.curr.X = (*i).length - 1

		searchColl(w, i)
	}

	//Aceptar las esquinas
	(*i).processed[0][0] = true
	(*i).processed[0][(*i).length-1] = true
	(*i).processed[(*i).width-1][0] = true
	(*i).processed[(*i).width-1][(*i).length-1] = true

}

// + BUSQUEDA
// - Ciclo: Hay colision?
// 		- Detectar colision en X
//			-Menos donde ya estuve
// 		- Detectar colision en Y
//			-Menos donde ya estuve
//		- Marcar como procesado
// - Coincidencias? repetir ciclo en pixel de colision
func searchColl(w WayStrt, i *Image) {

	var nw WayStrt

	nw.curr.X = 0
	nw.curr.Y = 0

	//fmt.Println()
	//fmt.Println("[Y:", w.curr.Y, ",X:", w.curr.X, "]", (*i).processed[w.curr.Y][w.curr.X], "C: ", (*i).pixelImage[w.curr.Y][w.curr.X])
	//fmt.Println()

	//no esta procesado?
	if !(*i).processed[w.curr.Y][w.curr.X] {

		//Es negro?
		if (*i).pixelImage[w.curr.Y][w.curr.X] {
			//Calcula3
			(*i).processed[w.curr.Y][w.curr.X] = true

			//Calcular el siguiente paso

			//PRIMER PASO

			if w.curr.Y == 0 { //arriba
				//fmt.Println("PRIMER PASO ARRIBA")
				w.curr.Y++
				searchColl(w, i)
			} else if w.curr.Y == (*i).width-1 { //abajo
				//fmt.Println("PRIMER PASO ABAJO")
				w.curr.Y--
				searchColl(w, i)
			} else if w.curr.X == 0 { //izquierda
				//fmt.Println("PRIMER PASO IZQUIERDA")
				w.curr.X++
				searchColl(w, i)
			} else if w.curr.X == (*i).length-1 { //derecha
				//fmt.Println("PRIMER PASO DERECHA")
				w.curr.X--
				searchColl(w, i)
			} else {
				//Cap derecha
				if w.curr.X < (*i).length-1 && (*i).pixelImage[w.curr.Y][w.curr.X+1] && !(*i).processed[w.curr.Y][w.curr.X+1] {
					//Hacia derecha?
					nw.curr.X = w.curr.X + 1
					nw.curr.Y = w.curr.Y
					//fmt.Println("der", nw.curr.Y, nw.curr.X)
					searchColl(nw, i)
				}

				//Cap izquierda
				if w.curr.X > 1 && (*i).pixelImage[w.curr.Y][w.curr.X-1] && !(*i).processed[w.curr.Y][w.curr.X-1] {
					//Hacia izquierda?
					nw.curr.X = w.curr.X - 1
					nw.curr.Y = w.curr.Y
					//fmt.Println("izq", nw.curr.Y, nw.curr.X)
					searchColl(nw, i)
				}

				//Cap abajo
				if w.curr.Y < (*i).width-1 && (*i).pixelImage[w.curr.Y+1][w.curr.X] && !(*i).processed[w.curr.Y+1][w.curr.X] {

					nw.curr.X = w.curr.X
					nw.curr.Y = w.curr.Y + 1
					//fmt.Println("aba ", nw.curr.Y, nw.curr.X)
					searchColl(nw, i)

				}
				//Cap arriba
				if w.curr.Y > 1 && (*i).pixelImage[w.curr.Y-1][w.curr.X] && !(*i).processed[w.curr.Y-1][w.curr.X] {
					//Hacia arriba?
					nw.curr.X = w.curr.X
					nw.curr.Y = w.curr.Y - 1
					//fmt.Println("arr", nw.curr.Y, nw.curr.X)
					searchColl(nw, i)

				}

			}

			//Se acabo la ruta
			return

		} else {

			//no es negro, pero si es orilla se marca igual

			if w.curr.Y == 0 || w.curr.Y == (*i).width-1 || w.curr.X == 0 || w.curr.X == (*i).length-1 {
				//fmt.Println("Orilla ")
				(*i).processed[w.curr.Y][w.curr.X] = true
			}
			//se acaba
			return
		}

	} else {
		//fmt.Println("X")
		//se acaba again
		return
	}

}

func cleanPixels(im *Image) {

	for i, row := range (*im).processed {
		for j, proBool := range row {
			if !proBool {

				(*im).pixelImage[i][j] = false

			}
		}
	}

}

// - Convertir a bools?

func uIntCon(preImage *[][]bool) [][]uint8 {
	fmt.Println("Deconversion")

	var rowInt []uint8
	var allInt [][]uint8

	//Conversion a bool

	rowInt = nil

	for _, row := range *preImage {

		rowInt = nil

		for _, pixel := range row {

			if pixel {
				rowInt = append(rowInt, 1)
			} else {
				rowInt = append(rowInt, 0)
			}

		}
		allInt = append(allInt, rowInt)

	}

	return allInt
}
