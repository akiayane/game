package main

type Ork struct {
	character
}

func newOrk() Character {
	Char := &Ork{ character : character{
		Name: "Ork",
		Hp: 20,
		Damage: 8,
		Defence: 6,
		Speed: 4,
		Status: true,
		Gold: 100,
	},

	}
	return Char
}
