package main

import (
	"fmt"
	"os"
)

type visitor interface {
	visitCastle(*Castle, *group)
	//expansion example visitWildAnimals, visitWildSeller
}

type visitForFight struct {
	myGroup *group
	enemyGroup *group
}
func (vff *visitForFight) visitCastle(castle *Castle, gp *group) {
	var groupHpCopy []int

	vff.myGroup = gp
	vff.enemyGroup = castle.group

	for _, s := range vff.myGroup.cells {
		groupHpCopy = append(groupHpCopy, s.getHp())
	}

	result := vff.myGroup.BattleStart(vff.enemyGroup)

	if result {

		for i, s := range vff.myGroup.cells {
			s.setHp(groupHpCopy[i])
		}

		castle.setFriendly()

		for  {
			fmt.Println("Type exit to leave from castle")


			var input string
			fmt.Scan(&input)

			if input == "exit" {break}
		}
	} else {
		fmt.Println("Game Over")
		os.Exit(0)
	}
}

type visitForTrade struct {
	myGroup *group
	mainChar Character
}
func (vft *visitForTrade) visitCastle(castle *Castle, group *group) {
	vft.myGroup = group
	vft.mainChar = group.cells[0]
	//interface for trading

	for {
		fmt.Println()
		fmt.Println("1.Manage Group" +
			" 2.Buy Items" +
			" 3.Level Up Group" +
			". Type exit to leave from castle")

		var input string
		fmt.Scan(&input)

		if input == "exit" {break}

		switch input {
		case "1":
			for {
				vft.myGroup.readGroup()
				fmt.Println("1.Remove from group" +
					" 2.Add to group" +
					". Type back to return to previous menu")

				var input string
				fmt.Scan(&input)

				if input == "back" {break}

				switch input {
				case "1":
					for {
						fmt.Println("Whom you are going to remove? Type back to return to previous menu")
						vft.myGroup.readGroup()

						var input string
						fmt.Scan(&input)

						if input == "back" {break}

						switch input {
						case "1":
							fmt.Println("Main Character can not be removed")
						case "2":
							if len(vft.myGroup.cells) > 1 {
								vft.myGroup.KickFromGroup(vft.myGroup.cells[1])
							} else {
								fmt.Println("Incorrect input")
							}
						case "3":
							if len(vft.myGroup.cells) > 2 {
								vft.myGroup.KickFromGroup(vft.myGroup.cells[2])
							} else {
								fmt.Println("Incorrect input")
							}
						case "4":
							if len(vft.myGroup.cells) > 3 {
								vft.myGroup.KickFromGroup(vft.myGroup.cells[3])
							} else {
								fmt.Println("Incorrect input")
							}
						default:
							fmt.Println("Incorrect input")
						}
					}
				case "2":
					for {
						fmt.Println("Whom you are going to add? Type back to return to previous menu")

						tmpPeasant := newPeasant()
						tmpOrk := newOrk()
						tmpGoblin := newGoblin()
						tmpPikeman := newPikeman()

						fmt.Println()

						tmpPeasant.getStats()
						tmpOrk.getStats()
						tmpGoblin.getStats()
						tmpPikeman.getStats()
						fmt.Println("1.Peasant 2.Ork 3.Goblin 4.Pikeman")
						fmt.Println("Gold:", vft.mainChar.getGold())
						fmt.Print("Current group: ")
						vft.myGroup.readGroup()


						var input string
						fmt.Scan(&input)

						if input == "back" {
							clear()
							break
						}

						switch input {
						case "1":
							if vft.mainChar.getGold() >= tmpPeasant.getGold() {
								if group.counter < 4 {
									vft.mainChar.setGold(vft.mainChar.getGold() - tmpPeasant.getGold())
								}
								vft.myGroup.AddToGroup(newPeasant())
							} else {
								fmt.Println("Not Enough Gold")
							}
						case "2":
							if vft.mainChar.getGold() >= tmpOrk.getGold() {
								if group.counter < 4 {
									vft.mainChar.setGold(vft.mainChar.getGold() - tmpOrk.getGold())
								}
								vft.myGroup.AddToGroup(newOrk())
							} else {
								fmt.Println("Not Enough Gold")
							}
						case "3":
							if vft.mainChar.getGold() >= tmpGoblin.getGold() {
								if group.counter < 4 {
									vft.mainChar.setGold(vft.mainChar.getGold() - tmpGoblin.getGold())
								}
								vft.myGroup.AddToGroup(newGoblin())
							} else {
								fmt.Println("Not Enough Gold")
							}
						case "4":
							if vft.mainChar.getGold() >= tmpPikeman.getGold() {
								if group.counter < 4 {
									vft.mainChar.setGold(vft.mainChar.getGold() - tmpPikeman.getGold())
								}
								vft.myGroup.AddToGroup(newPikeman())
							} else {
								fmt.Println("Not Enough Gold")
							}
						default:
							fmt.Println("Incorrect input")
						}
					}
				}
			}
			// manage group
		case "2":
			for {
				fmt.Println("Buy Item Should Be fixed")

				var input string
				fmt.Scan(&input)

				if input == "back" {break}
			}
			// bay item
		case "3":
			for {
				fmt.Println("For some amount of money you can upgrade your group.")

				lvlUpGold := vft.myGroup.lvl*50

				fmt.Println("Current level of group:", vft.myGroup.lvl, "Gold:", vft.mainChar.getGold())
				fmt.Println(lvlUpGold, "gold required for level up")

				fmt.Println("Do you want upgrade your group? 1.Yes 2.No")

				var input string
				fmt.Scan(&input)

				if input == "back" {break}

				switch input {
				case "1":
					if vft.mainChar.getGold() >= lvlUpGold {
						vft.mainChar.setGold(vft.mainChar.getGold() - lvlUpGold)
						vft.myGroup.levelUpGroup()
					} else {
						fmt.Println("Not Enough Gold")
					}
				case "2":
					input = "back"
				default:
					fmt.Println("Incorrect input")
				}

				if input == "back" {break}
			}
			//levelup
		default:
			fmt.Println("Incorrect input")


		}
	}


}
