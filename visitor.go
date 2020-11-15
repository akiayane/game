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
		fmt.Println("1. Manage Group" +
			" 2. Buy Items" +
			" 3. Level Up Group" +
			". Type exit to leave from castle")

		var input string
		fmt.Scan(&input)

		if input == "exit" {break}

		switch input {
		case "1":
			for {
				vft.myGroup.readGroup()
				fmt.Println("1. Remove from group" +
					" 2. Add to group" +
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
						fmt.Println("1. Peasant 2. Ork 3. Goblin 4. Pikeman")

						var input string
						fmt.Scan(&input)

						if input == "back" {break}

						switch input {
						case "1":
							vft.myGroup.AddToGroup(newPeasant())
						case "2":
							vft.myGroup.AddToGroup(newOrk())
						case "3":
							vft.myGroup.AddToGroup(newGoblin())
						case "4":
							vft.myGroup.AddToGroup(newPikeman())
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
				fmt.Println("For some amount of money you can upgrade your group. " +
					"Type back to return to previous menu")
				fmt.Println("Gold structure should be fixed")

				fmt.Println("Level of group:", vft.myGroup.lvl)

				vft.myGroup.levelUpGroup()

				var input string
				fmt.Scan(&input)

				if input == "back" {break}
			}
			//levelup
		default:
			fmt.Println("Incorrect input")


		}
	}


}
