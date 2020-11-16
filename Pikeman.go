package main

type Pikeman struct {
	character
}

func newPikeman() Character {
	Char := &Pikeman{ character : character{
		Name: "Pikeman",
		Hp: 20,
		Damage: 10,
		Defence: 5,
		Speed: 5,
		CritChance: 8,
		Status: true,
		Gold: 125,
	},

	}
	return Char
}
