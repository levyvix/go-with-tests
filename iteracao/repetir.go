package iteracao

import "strings"

const (
	quantidadeRepeticoes = 5
)

func Repetir(caractere string) string {
	return strings.Repeat(caractere, quantidadeRepeticoes)
}
