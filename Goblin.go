package main

type Goblin struct {
	Character
}

func newGoblin() *character {
	Char := &character{
		Name: "Goblin",
		Hp: 10,
		Damage: 5,
		Defence: 5,
		Speed: 6,
		Status: true,
	}
	return Char
}
