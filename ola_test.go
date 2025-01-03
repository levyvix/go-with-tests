package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Levy")
	esperado := "Ola, Levy"
	if resultado != esperado {
		t.Errorf("esperado '%s', resultado '%s'", esperado, resultado)
	}
}
