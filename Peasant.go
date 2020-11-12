package main

type Peasant struct {
	Character
}

func newPeasant() *character {
	Char := &character{
		Name: "Peasant",
		Hp: 10,
		Damage: 4,
		Defence: 4,
		Speed: 5,
		Status: true,
	}
	return Char
}
