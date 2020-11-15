package main

func initialize() []CastleIcon{

	myCastle := newCastle("Stratholm")
	myCastle.setFriendly()
	myCastleIcon := newCastleIcon(*myCastle, true, 13, 6)

	groupSample := NewGroup(newOrk())
	groupSample.AddToGroup(newPikeman())
	groupSample.AddToGroup(newPeasant())
	groupSample.AddToGroup(newGoblin())



	firstEnemyCastle := newCastle("Stromgard")
	firstEnemyCastle.setGroup(groupSample)
	firstEnemyCastleIcon := newCastleIcon(*firstEnemyCastle, false, 64, 9)

	waterCastle := newCastle("Stromgard")
	//castletree.setGroup(groupSample)
	waterCastleIcon := newCastleIcon(*waterCastle, false, 121, 4)

	hillCastle := newCastle("Stromgard")
	//castlefour.setGroup(groupSample)
	hillCastleIcon := newCastleIcon(*hillCastle, true, 102, 17)

	mountainCastle := newCastle("Stromgard")
	//castlefour.setGroup(groupSample)
	mountainCastleIcon := newCastleIcon(*mountainCastle, true, 41, 27)

	farFarCastle := newCastle("Stromgard")
	//castlefour.setGroup(groupSample)
	farFarCastleIcon := newCastleIcon(*farFarCastle, false, 122, 30)

	var casSlice []CastleIcon
	casSlice = append(casSlice, *myCastleIcon)
	casSlice = append(casSlice, *firstEnemyCastleIcon)
	casSlice = append(casSlice, *waterCastleIcon)
	casSlice = append(casSlice, *hillCastleIcon)
	casSlice = append(casSlice, *mountainCastleIcon)
	casSlice = append(casSlice, *farFarCastleIcon)

	return casSlice
}
