package main

type Peasant struct {
	character
}

func newPeasant() Character {
	Char := &Peasant{ character : character{
		Name: "Peasant",
		Hp: 10,
		Damage: 4,
		Defence: 4,
		Speed: 5,
		CritChance: 10,
		Status: true,
		Gold: 30,
	},

	}
	return Char
}
