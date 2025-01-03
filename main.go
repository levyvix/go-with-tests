package main

import "fmt"

const (
	prefixoOlaPortugues = "Ola,"
	prefixoOlaEspanhol  = "Hola,"
	prefixoOlaFrances   = "Bonjour,"
	espanhol            = "espanhol"
	frances             = "frances"
)

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {

	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func Ola(name, idioma string) string {
	if name == "" {
		name = "Mundo"
	}

	return fmt.Sprintf("%s %s", prefixoDeSaudacao(idioma), name)

}
func main() {
	fmt.Println(Ola("Levy", ""))
}
