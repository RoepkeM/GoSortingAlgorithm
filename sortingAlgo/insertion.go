package ordenacion


//This function receives a list of elements type int and returns the ordered list and the name of the algorithm
func Insertion(list []int) ([]int,string){

	for i := 0; i <=len(list)-1 ; i++ {
		j:=i //This variable will set the margin between the sorter and unsorted list
		for j>0 { //This loop will start in the end of the sorted list
			if list[j] < list[j-1] { //We compare the ordered part with the incoming item
				list[j], list[j-1]= list[j-1], list[j] //If it is bigger we exchange it
			}
			j-= 1//The counter is reduced
		}
	}
	return list,"Isertion" // We return the list and the name of the algorithm
}

