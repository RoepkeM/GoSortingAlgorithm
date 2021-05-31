package ordenacion

//Algoritmo de Selleccion

//Esta funcion recive una lista de elemntos tipo int y regresa la lista ordenada y el nombre del algoritmo
func Seleccion(lista []int) ([]int,string){
	for i := 0; i < len(lista)-1; i++ { //Iniciamos un foor loop que recorre la lista

		for j:=i;j<len(lista);j++ { //Recorremos la lista para comparar el elemento 'i' con el resto

			if lista[j]<lista[i] { // Si un elemento es menor que el comparado, estos se intercambian
				lista[i],lista[j] = lista[j],lista[i]
			}
		}
	}
	return lista,"Selleccion"//Regresamos la lista ordenada y el nombre del algoritmo
}
