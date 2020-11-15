package main

func initialize() []*CastleIcon{

	myCastle := newCastle("Homeland")
	myCastle.setFriendly()
	myCastleIcon := newCastleIcon(*myCastle, true, 13, 6)

	firstEnemyCastleGroup := NewGroup(newGoblin())
	firstEnemyCastleGroup.AddToGroup(newGoblin())
	firstEnemyCastleGroup.AddToGroup(newPeasant())
	firstEnemyCastleGroup.AddToGroup(newPeasant())

	firstEnemyCastle := newCastle("Dark Dungeons")
	firstEnemyCastle.setGroup(firstEnemyCastleGroup)
	firstEnemyCastleIcon := newCastleIcon(*firstEnemyCastle, false, 64, 9)

	waterCastleGroup := NewGroup(newPikeman())
	waterCastleGroup.AddToGroup(newPikeman())
	waterCastleGroup.AddToGroup(newPeasant())
	waterCastleGroup.AddToGroup(newPeasant())

	waterCastle := newCastle("Aqua Heaven")
	waterCastle.setGroup(waterCastleGroup)
	waterCastleIcon := newCastleIcon(*waterCastle, false, 121, 4)

	hillCastleGroup := NewGroup(newOrk())
	hillCastleGroup.AddToGroup(newGoblin())
	hillCastleGroup.AddToGroup(newGoblin())
	hillCastleGroup.AddToGroup(newGoblin())

	hillCastle := newCastle("Pit of Death")
	hillCastle.setGroup(hillCastleGroup)
	hillCastleIcon := newCastleIcon(*hillCastle, true, 102, 17)

	mountainCastleGroup := NewGroup(newOrk())
	mountainCastleGroup.AddToGroup(newPikeman())
	mountainCastleGroup.AddToGroup(newPeasant())
	mountainCastleGroup.AddToGroup(newGoblin())

	mountainCastle := newCastle("Windy Peaks")
	mountainCastle.setGroup(mountainCastleGroup)
	mountainCastleIcon := newCastleIcon(*mountainCastle, true, 41, 27)

	farFarCastleGroup := NewGroup(newOrk())
	farFarCastleGroup.AddToGroup(newOrk())
	farFarCastleGroup.AddToGroup(newOrk())
	farFarCastleGroup.AddToGroup(newOrk())

	farFarCastle := newCastle("Gachi Club")
	farFarCastle.setGroup(farFarCastleGroup)
	farFarCastleIcon := newCastleIcon(*farFarCastle, false, 122, 30)

	var casSlice []*CastleIcon
	casSlice = append(casSlice, myCastleIcon)
	casSlice = append(casSlice, firstEnemyCastleIcon)
	casSlice = append(casSlice, waterCastleIcon)
	casSlice = append(casSlice, hillCastleIcon)
	casSlice = append(casSlice, mountainCastleIcon)
	casSlice = append(casSlice, farFarCastleIcon)

	return casSlice
}
