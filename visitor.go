package main

type visitor interface {
	visitCastle(*Castle, *group)
	//expansion example visitWildAnimals, visitWildSeller
}

type visitForFight struct {
	myGroup *group
	enemyGroup *group
}
func (vff *visitForFight) visitCastle(castle *Castle, group *group) {
	vff.myGroup = group
	vff.enemyGroup = castle.group
	vff.myGroup.BattleStart(vff.enemyGroup)
}

type visitForTrade struct {
	myGroup *group
	mainChar *character
}
func (vft *visitForTrade) visitCastle(castle *Castle, group *group) {
	vft.myGroup = group
	vft.mainChar = group.cells[0]
	//interface for trading
}
