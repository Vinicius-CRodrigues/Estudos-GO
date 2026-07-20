package atckgame

type Attacker interface {
	Attack() int
}

type Warrior struct {
	Name  string
	Power int
}

type Wizard struct {
	Name       string
	MagicPower int
}

func (War Warrior) Attack() int {
	Damage := 2 * War.Power

	return Damage
}

func (Wiz Wizard) Attack() int {
	WizDamage := 3 * Wiz.MagicPower
	return WizDamage
}
