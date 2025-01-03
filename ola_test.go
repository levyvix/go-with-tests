package main

import (
	"testing"
)

func TestOla(t *testing.T) {
	verificaMensagem := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("esperado '%s', resultado '%s'", esperado, resultado)
		}
	}
	t.Run("diz ola para as pessoas", func(t *testing.T) {
		resultado := Ola("Levy", "")
		esperado := "Ola, Levy"
		verificaMensagem(t, resultado, esperado)
	})

	t.Run("diz 'Ola, Mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Ola, Mundo"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("levy", "espanhol")
		esperado := "Hola, levy"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("em farance", func(t *testing.T) {
		resultadp := Ola("levy", "frances")
		esperado := "Bonjour, levy"

		verificaMensagem(t, resultadp, esperado)
	})

}
