package main

func initialize() []CastleIcon{

	castleone := newCastle("Stratholm")
	castleone.setFriendly()
	castleOneIcon := newCastleIcon(*castleone, true, 10, 10)

	peasant := newPeasant()
	peasant.Enemy = true
	goblin := newGoblin()
	goblin.Enemy = true
	ork := newOrk()
	ork.Enemy = true
	pikeman := newPikeman()
	pikeman.Enemy = true

	groupSample := NewGroup(peasant)
	groupSample.AddToGroup(goblin)
	groupSample.AddToGroup(ork)
	groupSample.AddToGroup(pikeman)



	castletwo := newCastle("Stromgard")
	castletwo.setGroup(groupSample)
	castleTwoIcon := newCastleIcon(*castletwo, false, 7, 2)

	var casSlice []CastleIcon
	casSlice = append(casSlice, *castleOneIcon)
	casSlice = append(casSlice, *castleTwoIcon)

	return casSlice
}
