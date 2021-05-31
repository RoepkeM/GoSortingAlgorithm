package main

import (
	"fmt"
	"github.com/RoepkeM/GoSortinAlgorithm/sortingAlgo"
	"math/rand"
	"time"
)


//Initial function: in this function we will generate a list of random of any size that we desired
//then, we will call the algorithm that we want to order the
//list. Around this call is a timer to measure the execution time.
func main() {

	list := []int{}      //We generate an empty slice
	numElements := 10000 // We chose the amount of element in the slice
	for i := 0; i < numElements; i++ {
		list = append(list,rand.Intn(100000))
	}
	var name string

	//If you want, you could print the unsorted list
	//fmt.Printf("The unsorted list is: %v \n",list)

	t0 := time.Now()//Start of the timer

	list, name =ordenacion.Burbuja(list)
	//list,name=ordenacion.Seleccion(list)
	//list,name=ordenacion.Insercion(list)
	//list,name=ordenacion.InicioMerge(list)
	//list,name=ordenacion.Quicksort(list)

	t1 := time.Since(t0)//End of the timer

	//If you want, you could print the sorted list
	//fmt.Printf("La lista ordenada es: %v \n",list)

	//We print the results of the timer
	fmt.Printf("The execution time with the %s algorithm is %v with %d elements", name,t1, numElements)

}
