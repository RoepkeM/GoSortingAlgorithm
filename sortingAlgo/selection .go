package ordenacion


//This function receives a list of elements type int and returns the ordered list and the name of the algorithm
func Selection (list []int) ([]int,string){
	for i := 0; i < len(list)-1; i++ { //We go through the list

		for j:=i;j<len(list);j++ { //We go through the list to compare the element 'i' with the rest

			if list[j]< list[i] { //If an element is less than the one compared, these are exchanged
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list,"Selection " //Regresamos la list ordenada y el nombre del algoritmo
}
