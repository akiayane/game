package main

type Ork struct {
	Character
}

func newOrk() *character {
	Char := &character{
		Name: "Ork",
		Hp: 20,
		Damage: 8,
		Defence: 6,
		Speed: 4,
		Status: true,
	}
	return Char
}
