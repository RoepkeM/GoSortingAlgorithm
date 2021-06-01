package ordenacion

//This function receives a list of elements type int and returns the ordered list and the name of the algorithm
func Bubble(list []int) ([]int,string){
	isOrdered := false //we set a bool variable to false, this means that the list is not ordered

	for !isOrdered {  // we start a loop that will not end until the value is not(!) true

		isOrdered = true //if the list is ordered this variable will remain in its true value and in the next iteration it will exit the loop

		for i := 0; i < len(list)-1; i++ { //we start a foor loop to go through the list, this goes to the second to last position of the list

			if list[i]> list[i+1] { //If element i of the list is greater than element i + 1 these are exchanged
				list[i], list[i+1] = list[i+1], list[i] //We exchange the elements
				isOrdered = false                    // the list is not ordered, we return to the beginning of the loop
			}
		}
	}

	return list, "Bubble" //we return the result of the ordered list and its name
}




