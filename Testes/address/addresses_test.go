package address_test // Usando um outro pacote dentro de um mesmo pacote.
// Não recomendado por não ser uma boa prática.

import (
	"ferramenta-testes/address" // Importando o pacote para ser usado no mesmo pacote.
	"testing"
)

type testCase struct {
	address        string
	expectedReturn string
}

func TestTypeOfAddresses(t *testing.T) {

	testCases := []testCase{
		{"Rua 12", "Rua"},
		{"Chacara 93", "Chacara"},
		{"Estrada", "Invalid type"},
		{"Avenida Floriano Peixoto", "Avenida"},
		{"Quadra 106", "Quadra"},
		{"", "Invalid type"},
	}

	for _, cenario := range testCases {
		receivedAddress := address.TypeOfAddresses(cenario.address)

		if receivedAddress != cenario.expectedReturn {
			t.Errorf("Endereço: %s, Esperado: %s, Recebido: %s", cenario.address, cenario.expectedReturn, receivedAddress)
		}

	}

}
