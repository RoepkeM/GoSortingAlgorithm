package ordenacion

//algoritmo de Merge-Sort
//Este será la función de llamado del algoritmo de Merge-sort
func InicioMerge(lista []int) ([]int,string) {
	l := mergeSort(lista) //llamamos a la función  mergeSort que tiene el algoritmo y guardamos su respuesta en un variable
	return l,"Merge-Sort"//regresamos la lista ordenada  y el nombre del algoritmo
}

//parte 1 del algoritmo Merge-sort
//Esta función recibe un lista de elementos de tipo int y regresa una lista de elementos de tipo int
func mergeSort(lista []int) []int{
	var num = len(lista) // Guardamos el tamaño de la lista

	//Si la lista tiene 1 solo elemento regresa la lista, esto sirve para delimitar la recursividad hasta que tenga un
	//solo elemento y no intente seguir en un loop infinito
	if num == 1 {
		return lista
	}

	middle := int(num / 2) //En una variable ponemos el valor numérico del tamaño de la lista dividido en 2

	//Creamos 2 listas para poner la una mitad de la lista y la otra en por separado
	var (
		left = make([]int, middle) //tamaño de esta lista es la mitad del tamño de la lista.
		right = make([]int, num-middle) //el tamaño de esta es la lista  menos la mitad, diferente a la de arriba en
		                                // caso de que sea una lista impar
	)

	//Poblamos las listas creadas
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = lista[i]
		} else {
			right[i-middle] = lista[i]
		}
	}

	//Recursivamente llamamos a una función que como parámetros  tiene el llamado de esta misma función dos veces
	//Ya que la función merge regresa una lista de int que es lo mismo que esta función tiene que regresar
	return merge(mergeSort(left), mergeSort(right))
}

//parte 2 del algoritmo Merge-sort
//Esta función pide 2 listas y regresa una lista
func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right)) // Hacemos una variable resultado que es de tamaño de las 2 variables en
	                                             // los parámetros

	//Recorremos las listas y comparamos sus elementos y ordenamos en la variable resultado para después devolver la
	//variable resultado como una lista ordenada
	i := 0//variable contadora
	for len(left) > 0 && len(right) > 0 { //Comprobamos en el loop que no nos salgamos del tamaño de las listas
		if left[0] < right[0] {//Comparamos elemento en la posición inicial de la una lista con el inicial de la otra lista
			result[i] = left[0]//Si el elemento de la lista de la izquierda es mayo que el de la derecha lo intercambiamos
			left = left[1:]//La lista de la izquierda ahora solo es el elemento de la recha
		} else {//Si no es mayor el elemento de la izquierda al de la derecha solo movemos los elementos en el orden que están a la lista resultado
			result[i] = right[0]
			right = right[1:]
		}
		i++//sumamos uno al contador
	}

	//ordenamos los elementos dentro de la variable resultado una vez ya organizados
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return //Regresamos la variable ya especificada en los parámetros de la función
}
