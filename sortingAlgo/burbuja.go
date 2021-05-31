package ordenacion


//Algoritmo de Burbuja

//Esta funcion recive una lista de elemntos tipo int y regresa la lista ordenada y el nombre del algoritmo
func Burbuja(lista []int) ([]int,string){
	estaOrdenand := false //iniciamos una variable bool a falso, esto significa que la lista no esta ordenada

	for !estaOrdenand {  // iniciamos un loop que no terminara hasta que el valor no(!) sea falso

		estaOrdenand = true //si la lista esta ordenada esta variable permanecerá en su valor de verdadero y en la siguiente repetición saldra del loop

		for i := 0; i < len(lista)-1; i++ {//iniciamos un foor loop para recorrer por la lista, este recorre hasta la penúltima posición de la lista

			if lista[i]>lista[i+1] {// si el elemento i de la lista es mayor que el elemento i+1 estos se intercambia
				lista[i], lista[i+1] = lista[i+1],lista[i] //intercambiamos los elementos
				estaOrdenand = false // la lista no esta ordenada, regresamos al inicio.
			}
		}
	}

	return lista, "Borbuja" // regresamos el resultado de la lista ordenada y su nombre
}




