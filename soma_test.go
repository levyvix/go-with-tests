package main

import "testing"

func TestSoma(t *testing.T) {
	t.Run("colecao de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15
		if resultado != esperado {
			t.Errorf("esperado %d, resultado %d", esperado, resultado)
		}
	})

	t.Run("colecao de qualquer tamahno", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6
		if resultado != esperado {
			t.Errorf("esperado %d, resultado %d", esperado, resultado)
		}
	})
}
