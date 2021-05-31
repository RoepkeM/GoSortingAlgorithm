package ordenacion

import "math/rand"

//algoritmo de Quicksort
//Esta funcion recive una lista de elemntos tipo int y regresa la lista ordenada y el nombre del algoritmo
func Quicksort(lista []int) ([]int,string){

	//Si la lista tiene solo 2 variables regresamos, esto nos sirve para limitar la recursividad y que no sea infinito
	if len(lista) < 2 {
		return lista,"Quick sort"
	}

	left, right := 0, len(lista)-1 //Ponemos el valor del tamaño de la lista en una variable y una variable iniciada a 0

	pivot := rand.Int() % len(lista)//Designamos el elemento pivote

	lista[pivot], lista[right] = lista[right], lista[pivot] //Cambiamos de posición al pivote con el elemento de la derecha

	for i, _ := range lista { //Recorremos la lista que tenemos
		if lista[i] < lista[right] {//Comparamos cual elemento es menos o mayo y lo intercambiamos
			lista[left], lista[i] = lista[i], lista[left]
			left++
		}
	}

	lista[left], lista[right] = lista[right], lista[left]

	//Lamamos recursivamente este mismo elemento para hacer el mismo procedimiento de ordenar el elemento pivote hasta
	//que solo haya 2 elementos en cada lista
	Quicksort(lista[:left])
	Quicksort(lista[left+1:])

	return lista,"Quick_sort" //regresamos la lista de elementos ordenada con el nombre del algoritmo
}
