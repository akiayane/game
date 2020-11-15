package main

type MainCharacter struct {
	character
	Level int
	Location int
	gold int
}

func (m *MainCharacter) levelUp() {
	m.Level++
	m.Hp+=10
	m.Damage+=2
	m.Defence+=2
	m.Speed+=1
}

func newMainCharacter(name string) Character {
	MainCharacter := &MainCharacter{ character : character{
		Name: name,
		Level: 1,
		Hp: 10,
		Damage: 5,
		Defence: 5,
		CritChance: 10,
		Speed: 5,
		Status: true,
		Friend: true,
		Gold: 0,
	},
	}
	return MainCharacter
}