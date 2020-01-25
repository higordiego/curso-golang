package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte %s) - índices: esperado (%d) <> enconttado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		atual := strings.Index(teste.texto, teste.parte)
		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
