package arrays

//func Soma(numeros [5]int) int {
//	soma := 0
//	for _, numeros := range numeros {
//		soma += numeros
//
//	}
//	return soma
//}

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
