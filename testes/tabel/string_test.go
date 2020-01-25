package strings

import "testing"
import "strings"

const msgIndex = "%s (parte %s) - índices: esperado (%d) <> enconttado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"opa", "opa", -1},
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.esperado, teste.parte)
		if atual != teste.esperado {
			t.Errof(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
