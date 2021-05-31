package ordenacion

//algoritmo de insercion
//Esta funcion recive una lista de elemntos tipo int y regresa la lista ordenada y el nombre del algoritmo
func Insercion(lista []int) ([]int,string){

	//Creamos un loop para recorrer la lista
	for i := 0; i <=len(lista)-1 ; i++ {
		j:=i //Esta variable es del mimo taño que la del índice del loop
		for j>0 { //Hacemos otro loop del tamaño del índice del anterior loop y sale cuando estemos en elemento 0
			if lista[j] < lista[j-1] { //Comparamos la parte ordenada con el elemento entrante
				lista[j],lista[j-1]=lista[j-1],lista[j]//Si es mayo lo intercambiamos
			}
			j-= 1//El contador se reduce
		}
	}
	return lista,"Insercion" // regresamos la lista y el nombre del algoritmo
}

