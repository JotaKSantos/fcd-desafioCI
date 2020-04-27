package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	x := soma(2, 3)

	if x != 5 {
		t.Errorf("Erro ao efetuar a soma")
	}
}
