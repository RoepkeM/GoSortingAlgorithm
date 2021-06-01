package ordenacion

//This will be the calling function of the MergeStart-sort algorithm
func MergeStart(lista []int) ([]int,string) {
	l := mergeSort(lista) //We call the mergeSort function, and save its response in a variable
	return l,"Merge-Sort"// We return the ordered list and the name of the algorithm
}

//MergeStart-sort algorithm part 1
//This function receives a list of elements of type int and returns a list of elements of type int
func mergeSort(list []int) []int{
	var num = len(list) // We save the list size

	//
	//If the list has only 1 element, it returns the list, this serves to limit the recursion until it has a single
	//element and does not try to continue until an error occurs
	if num == 1 {
		return list
	}

	middle := int(num / 2) //This variable is the half of the size of the given list

	//We create 2 lists to put one half of the list and the other separately
	var (
		left = make([]int, middle) //The size of this list is half the size of the list.
		right = make([]int, num-middle) // The size of this is the list minus half, different than above in case it is
										// an odd list
	)

	//We populate the created lists
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = list[i]
		} else {
			right[i-middle] = list[i]
		}
	}

	//Recursively we call a function that as parameters has the call of this same function twice
    //In itch call we put the left side of the list and the right side
	return merge(mergeSort(left), mergeSort(right))
}

//Merge-sort algorithm part 2
//This function asks for 2 lists of type int and returns 1 list of type int
func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right)) // We make a list variable that is the size of the 2 variables
												// in the parameters

	//We go through the lists and compare their elements and order in the result variable and then return the result
	//variable as an ordered list
	i := 0//counter
	for len(left) > 0 && len(right) > 0 { //We check in the loop that we do not go beyond the size of the lists
		if left[0] < right[0] {//We compare element in the initial position of the one list with the initial position of the other list
			result[i] = left[0]//If the item in the list on the left is May than the one on the right, we swap it
			left = left[1:]//The list on the left is now only the item on the right
		} else {// If the element on the left is not greater than the one on the right, we only move the elements in the order that they are to the result list
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	//We order the elements within the result list once they are already organized
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return //We return the variable already specified in the function parameters as result
}
