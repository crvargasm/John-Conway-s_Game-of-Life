package main

/**
Regla #1: Una célula muerta con exactamente 3 células vecinas vivas "nace" (es decir, al turno siguiente estará viva).
Regla #2: Una célula viva con 2 o 3 células vecinas vivas sigue viva, en otro caso muere (por "soledad" o "superpoblación").
*/

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var grilla [][]int

const cuadrilla int = 21
const columnas int = cuadrilla
const filas int = cuadrilla

/*La lectura de archivos requiere comprobar la mayoría de las llamadas en busca de errores.
Esta función simplifica las comprobaciones de error.*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func draw(matrix [columnas][filas]int) {
	fmt.Println()
	for i := 0; i < ((columnas)*2)+1; i++ {
		fmt.Print("▁", "")
	}
	fmt.Println()
	for i := 0; i < columnas; i++ {
		fmt.Print("", " ")
		for j := 0; j < filas; j++ {
			if matrix[i][j] == 1 {
				fmt.Print("■", " ")
			} else {
				fmt.Print(" ", " ")
			}
		}
		fmt.Println("▏")
	}
	for i := 0; i < ((columnas)*2)+1; i++ {
		fmt.Print("▔", "")
	}
	fmt.Println("")
}

func contarVecinos(matrix [columnas][filas]int, x int, y int) int {
	var sum int = 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {

			var col int = (x + i + columnas) % columnas
			var fil int = (y + j + filas) % filas

			sum += matrix[col][fil]
		}
	}
	sum -= matrix[x][y]
	return sum
}

/*Generamos un tablero Aleatorio con un patron aleatorio con Seed basada en la hora actual*/
func tableroAleatorio(grilla [columnas][filas]int) [columnas][filas]int {
	rand.Seed(time.Now().UTC().UnixNano()) //Se genera la semilla a partir de la hora exacta en Nanosegundos
	for i := 0; i < columnas; i++ {        //Recorremos cada casilla para asignar un valor
		for j := 0; j < filas; j++ {
			grilla[i][j] = rand.Intn(2) //Asignamos a esa posicion un valor aleatorio entre 0 o 1
		}
	}
	return grilla
}

func main() {

	var grilla [columnas][filas]int

	//Iteración para Menú Inicial
	a := 0
	for a != 1 {
		var w int = 0
		fmt.Println("Bienvenido a The John Conway's Game of Life")
		fmt.Println("Para Iniciar Ingrese el Modo que desea Visualizar, seguido de la tecla 'Enter':")
		fmt.Println("Nota: Para avanzar entre repeticiones presione la tecla Enter")
		fmt.Println("      Para salir del aplicativo ingrese el número 1 seguido de la tecla Enter")
		fmt.Println("1-. Aleatorio		2-. Matriz Cargada en Carpeta		3-. Salir")
		fmt.Scanln(&w)

		switch w {
		case 1: //Generamos un Tablero Aleatorio
			a = 1
			grilla = tableroAleatorio(grilla)
			break
		case 2: //
			//Cargar Matriz
			a = 1
			dat, err := ioutil.ReadFile("tablero.txt")
			check(err)
			//lastBit := columnas * 2 //último bit: columnas*2 [21*2 = 42] esto por el espacio en blanco entre ellos
			//fmt.Print(string(dat[0:42]))	//1a Fila; 0:lastBit
			//fmt.Print(string(dat[0:43]))	//Bit 43 = salto de línea; (lastBit)+1
			//fmt.Print(string(dat[0 : lastBit*2])) //44 linea 2; ((lastBit*n)-lastBit)+2 :bitInício línea n; lastBit*n=ultimo bit de la fila n
			for i := 0; i < cuadrilla; i++ { //Linea i / fila i   ----
				for j := 0; j < cuadrilla; j++ { //Columna j / hilera j |
					a, err := strconv.Atoi(string(dat[(43*i)+((j)*2)]))
					if err != nil {
						fmt.Println(err)
						os.Exit(2)
					}
					grilla[i][j] = a
				}
			}
			break
		case 3: //Salimos del Sistema
			os.Exit(0)
			break
		default:
			fmt.Println("----------------------------------------------------------")
			fmt.Println("Upps, has ingresado mal, intenta nuevamente.")
			fmt.Println("----------------------------------------------------------")
			fmt.Println()
			break
		}
	}

	q := 0
	for q != 1 {
		draw(grilla)
		fmt.Scanln(&q)
		fmt.Println()
		var next [columnas][filas]int //La siguiente grilla para declararle las reglas de juego

		for i := 0; i < columnas; i++ {
			for j := 0; j < filas; j++ {

				var state = grilla[i][j]
				var vecinos int = contarVecinos(grilla, i, j)

				if state == 0 && vecinos == 3 { //Regla #1
					next[i][j] = 1
				} else if state == 1 && (vecinos < 2 || vecinos > 3) { //Regla #2
					next[i][j] = 0
				} else {
					next[i][j] = state
				}

			}
		}
		grilla = next
	}

}
