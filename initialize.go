package main

func initialize() []CastleIcon{

	castleone := newCastle("Stratholm")
	castleone.setFriendly()
	castleOneIcon := newCastleIcon(*castleone, true, 10, 10)

	groupSample := NewGroup(newOrk())
	groupSample.AddToGroup(newPikeman())
	groupSample.AddToGroup(newPeasant())
	groupSample.AddToGroup(newGoblin())



	castletwo := newCastle("Stromgard")
	castletwo.setGroup(groupSample)
	castleTwoIcon := newCastleIcon(*castletwo, false, 7, 2)

	var casSlice []CastleIcon
	casSlice = append(casSlice, *castleOneIcon)
	casSlice = append(casSlice, *castleTwoIcon)

	return casSlice
}
