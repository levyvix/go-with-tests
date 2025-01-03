package iteracao

const (
	quantidadeRepeticoes = 5
)

func Repetir(caractere string) (resultado string) {
	for range quantidadeRepeticoes {
		resultado += caractere
	}

	return resultado
}
