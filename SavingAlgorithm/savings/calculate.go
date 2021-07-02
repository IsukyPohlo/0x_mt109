package savings

import (
	"fmt"
	"savingAlgo/readFile"
	"sort"
	"sync"
)

//Structs

type Saves struct {
	i, j int
	s    float32
}

type Graph struct {
	next     int
	distance int
}

//SORT
type BySaves []Saves

func (a BySaves) Len() int           { return len(a) }
func (a BySaves) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySaves) Less(i, j int) bool { return a[i].s > a[j].s }

//IMPORTANTE, Matrix[j][i]
//			  		[X][Y]

func CalculateSaves(rutas *readFile.Routes) []Saves {

	size := len(rutas.Matrix)

	//Slice de savings

	var savings []Saves

	var wg sync.WaitGroup

	savs := make(chan Saves)

	cont := 1

	for i := 1; i < size; i++ {

		for j := 1; j < size; j++ {

			if i != j {

				cont++

				//Mandar workers
				wg.Add(1)
				go SaveWorker(i, j, savs, rutas, &wg)

				//fmt.Println("S[", i, j, "] = d[", i, "0 ] (", di0, ") + d[ 0", j, "] (", d0j, ") - d[", i, j, "] (", dij, ") =", save.s)

			}
		}
	}

	//Obtener resultados

	go func() {

		for newSav := range savs {
			//fmt.Println("saving")
			savings = append(savings, newSav)
		}

	}()

	//Esperar a los workers

	wg.Wait()

	//Ordenar
	sort.Sort(BySaves(savings))

	//fmt.Println("Retornado")
	return savings
}

//SaveWorker

func SaveWorker(i int, j int, savs chan Saves, rutas *readFile.Routes, wg *sync.WaitGroup) {

	var save Saves

	var di0 = rutas.Matrix[0][i].DistanceInMeters
	var d0j = rutas.Matrix[j][0].DistanceInMeters
	var dij = rutas.Matrix[j][i].DistanceInMeters

	save.s = di0 + d0j - dij
	save.i = i
	save.j = j

	//fmt.Println("- ", i, j, di0, d0j, dij, save.s)

	savs <- save

	defer wg.Done()
}

//Calcular la ruta mas optima
func CalculateRoute(savings []Saves, size int) map[int]int {

	//fmt.Println("size: ", size)

	//Grafo con solo un vertice posible
	route := make(map[int]int)

	//Armar ruta mas optima

	//Agregar desde el arriba
	for i, vert := range savings {

		//Terminar al llenar el mapa
		if len(route) == size+1 {
			break
		}

		from := vert.i
		to := vert.j

		//Primer, ciclo, agregar directamente
		if i == 0 {
			route[0] = from
			route[from] = to
			route[to] = 0
			continue
		}

		_, okFrom := route[from]
		_, okTo := route[to]

		//fmt.Println(len(route))
		//fmt.Println(from, to)
		//printRoute(&route)

		//Si no existe ninguno (?)
		if !okFrom && !okTo {
			//fmt.Println("NO EXISTE NINGUNO")
			continue
		}

		//Si ambos existen, nada
		if okFrom && okTo {
			//fmt.Println("X")
			continue

		} else {

			//Si existe FROM
			if okFrom {

				//Agregar TO -> FROM->
				route[to] = route[from]

				//FROM -> TO
				route[from] = to

			} else

			//Si existe TO
			{

				//Buscar prevTO
				for key, val := range route {
					if val == to {
						//PrevTO -> FROM
						route[key] = from
					}
				}
				//FROM->tO
				route[from] = to
			}
			//fmt.Println("Added")
		}
	}

	return route
}

func PrintRoute(route *map[int]int) {

	prin := 0

	fmt.Print("(", prin, ") ")
	fmt.Print("(", (*route)[prin], ") ")

	prin = (*route)[prin]

	for {
		if prin == 0 {
			break
		}
		fmt.Print("(", (*route)[prin], ") ")
		prin = (*route)[prin]
	}

	fmt.Println()
}

func PrintRutaDist(route *map[int]int, dist *readFile.Routes) {

	prin := 0

	fmt.Print("(", prin, ") ")
	fmt.Print((*dist).Matrix[prin][(*route)[prin]].DistanceInMeters)

	for {
		if (*route)[prin] == 0 {
			break
		}
		fmt.Print(" (", (*route)[prin], ") ")
		prin = (*route)[prin]
		fmt.Print((*dist).Matrix[prin][(*route)[prin]].DistanceInMeters)
	}

	fmt.Println()
}
