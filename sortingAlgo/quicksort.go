package ordenacion

import "math/rand"

//This function receives a list of elements type int and returns the ordered list and the name of the algorithm
func Quicksort(list []int) ([]int,string){

	//If the list has only 2 variables we return, this helps us to limit the cursivity
	if len(list) < 2 {
		return list,"Quick sort"
	}

	left, right := 0, len(list)-1 // We put the value of the list size in a variable and a variable started at 0

	pivot := rand.Int() % len(list) //We designate the pivot element

	list[pivot], list[right] = list[right], list[pivot] // We change the position of the pivot with the element on the right

	for i, _ := range list { //We go through the list we have
		if list[i] < list[right] { //We compare which element is less or greater and we exchange it
			list[left], list[i] = list[i], list[left]
			left++
		}
	}

	list[left], list[right] = list[right], list[left]

	//Recursively we call this function to do the same procedure of sorting the pivot element until there are only 2
	//elements in each list
	Quicksort(list[:left]) //We do the process above whit the right side of the list
	Quicksort(list[left+1:]) //We do the process above whit the left side of the list

	return list,"Quick_sort" //We return the list of elements ordered with the name of the algorithm
}
