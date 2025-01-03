package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Expected %s but got %s", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for range b.N {
		Repetir("a")
	}
}
