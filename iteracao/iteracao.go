package iteracao

//func Repetir(caractere string) string {
//	var repeticoes string
//	for i := 0; i < 5; i++ {
//		repeticoes = repeticoes + caractere
//	}
//	return repeticoes
//}

const quantidadeRepeticoes = 5

func Repetir(caractere string) string  {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes ; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
