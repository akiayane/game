package main

import "fmt"

type MainCharacter struct {
	character
	Level int
	Location int
}

func (m *MainCharacter) levelUp() {
	m.Level++
	m.Hp+=10
	m.Damage+=2
	m.Defence+=2
	m.Speed+=1
	fmt.Println("Level Up !")
}

func newMainCharacter(name string) *character {
	MainCharacter := &character{
		Name: name,
		Level: 1,
		Hp: 10,
		Damage: 5,
		Defence: 5,
		CritChance: 10,
		Speed: 5,
		Status: true,
	}
	return MainCharacter
}