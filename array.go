package main

type Array[T any] []T

type GenericMapHandler[T any] func(v T) T

type GenericForEachHandler[T any] func(v T)

// La méthode forEach() permet d'exécuter une fonction donnée sur chaque élément du tableau.
func (array *Array[T]) ForEach(handler GenericForEachHandler[T]) {
	for _, v := range *array {
		handler(v)
	}
}

// La méthode map() crée un nouveau tableau avec les résultats de l'appel d'une fonction fournie sur chaque élément du tableau appelant.
func (array *Array[T]) Map(handler GenericMapHandler[T]) (newArray Array[T]) {
	for _, v := range *array {
		newArray = append(newArray, handler(v))
	}
	return newArray
}

// Here we are forced to create two method instead of one with type constraint because typeconstraint doesn't work that method

// La méthode ConcatArray() est utilisée afin de fusionner le tableau original avec le tableau en argument en les concaténant.
func (array *Array[T]) ConcatArray(value Array[T]) {
	*array = append(*array, value...)
}

// La méthode ConcatArrays() est utilisée afin de fusionner plusieurs tableaux avec le tableau d'origine en les concaténant.
func (array *Array[T]) ConcatArrays(arrays ...Array[T]) {
	for _, a := range arrays {
		*array = append(*array, a...)
	}
}

// La méthode ConcatValues() est utilisée afin de fusionné plusieurs élements avec le tableau original en les concaténant
func (array *Array[T]) ConcatValues(values ...T) {
	*array = append(*array, values...)
}

func (array *Array[T]) Pop() {
	newArray := *array

	if len(*array) > 0 {
		*array = newArray[1:]
	}
}

// La méthode filter() crée et retourne un nouveau tableau contenant tous les éléments du tableau d'origine qui remplissent une condition déterminée par la fonction callback
func (array *Array[T]) Filter(filterHandler func(value T) bool) (newArray Array[T]) {
	for _, a := range *array {
		if filterHandler(a) {
			newArray = append(newArray, a)
		}
	}
	return newArray
}

// La méthode FilterInternal est similaire à la fonction Filter sauf qu'elle modifie directement le tableau d'origine au lieu de renvoyer un nouveau tableau
func (array *Array[T]) FilterInternal(filterHandler func(value T) bool) {
	var newArray Array[T]
	for _, a := range *array {
		if filterHandler(a) {
			newArray = append(newArray, a)
		}
	}
	*array = newArray
}

// La méthode findIndex() renvoie l'indice du premier élément du tableau qui satisfait une condition donnée par une fonction. Si la fonction renvoie faux pour tous les éléments du tableau, le résultat vaut -1.
func (array *Array[T]) FindIndex(findHandler func(value T) bool) int {
	for idx, v := range *array {
		if findHandler(v) {
			return idx
		}
	}
	return -1
}
