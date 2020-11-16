package main

type Goblin struct {
	character
}

func newGoblin() Character {
	Char := &Goblin{ character : character{
		Name: "Goblin",
		Hp: 10,
		Damage: 5,
		Defence: 5,
		Speed: 6,
		CritChance: 10,
		Status: true,
		Gold: 50,
	},

	}
	return Char
}
