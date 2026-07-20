package atckgame

import "testing"

func TestAttackCharacter(t *testing.T) {
	t.Run("Warrior", func(t *testing.T) {
		warrior := Warrior{Name: "Olaf", Power: 20}

		powerExpected := 40

		powerReceived := warrior.Attack()

		if powerExpected != powerReceived {
			t.Errorf("poder esperado: %d. Poder recebido: %d", powerExpected, powerReceived)
		}
	})

	t.Run("Wizard", func(t *testing.T) {
		wizard := Wizard{Name: "Olaf", MagicPower: 20}

		powerExpected := 60

		powerReceived := wizard.Attack()

		if powerExpected != powerReceived {
			t.Errorf("poder esperado: %d. Poder recebido: %d", powerExpected, powerReceived)
		}
	})
}
