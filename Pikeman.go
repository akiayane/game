package main

type Pikeman struct {
	Character
}

func newPikeman() *character {
	Char := &character{
		Name: "Pikeman",
		Hp: 20,
		Damage: 10,
		Defence: 5,
		Speed: 5,
	}
	return Char
}
