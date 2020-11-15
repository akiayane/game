package main

import "fmt"

type gameFacade struct {
	gameMap *worldmap
}

func (gf *gameFacade) createMainChar(name string) Character {
	mainChar := newMainCharacter(name)

	return mainChar
}

func (gf *gameFacade) createGroup(char Character) *group {
	gp := NewGroup(char)

	return gp
}

func (gf *gameFacade) getCastleSlice() []*Castle {
	var result []*Castle

	for _,s:= range gf.gameMap.casSlice {
		result = append(result, s.castle)
	}

	return result
}

func (gf *gameFacade) getGroup() *group {
	return gf.gameMap.HeroIcon.group
}

func (gf *gameFacade) getMainChar() Character {
	return gf.gameMap.HeroIcon.group.cells[0]
}

func (gf *gameFacade) printMap() {

	w := &worldmap{}
	if gf.gameMap == nil {
		w = newWorldmap(newMainCharacter("Sample"), NewGroup(newMainCharacter("Sample")))
		gf.gameMap = w
	}

	w = gf.gameMap
	var i,j int
	w.myMap[w.HeroIcon.currentX][w.HeroIcon.currentY] = w.HeroIcon.icon
	for i:= 0; i< len(w.casSlice); i++{
		if w.casSlice[i].castle.friendly {
			w.casSlice[i].icon = "H"
		}
		w.myMap[w.casSlice[i].locX][w.casSlice[i].locY] = w.casSlice[i].icon
	}
	for j = 0; j<w.y; j++{
		for  i = 0; i < w.x; i++ {
			fmt.Print(w.myMap[i][j])
		}
		fmt.Print("\n")
	}
}

func (gf *gameFacade) printCastleCoordinates() {
	gf.gameMap.getCoor()
}

func (gf *gameFacade) startGame() {
	fmt.Println("Welcome to the lands of the Northern Daugras!")
	fmt.Println("First of all give the name to your main character")

	var input string
	fmt.Scan(&input)

	mainChar := gf.createMainChar(input)
	mainChar.setGold(200)
	mainName = mainChar.getName()
	mainGroup := gf.createGroup(mainChar)

	clear()
	myMap := newWorldmap(mainChar, mainGroup)
	myMap.cage()

	for{
		fmt.Println("Use w,s,a,d to move. Type exit to leave!")
		myMap.checkdoors()
		var dir string
		fmt.Scan(&dir)
		fmt.Print("\n")
		if dir == "exit"{
			clear()
			break
		}
		if dir == "w"{
			clear()
			myMap.moveUp()
		}
		if dir == "s"{
			clear()
			myMap.moveDown()
		}
		if dir == "d"{
			clear()
			myMap.moveRight()
		}
		if dir == "a"{
			clear()
			myMap.moveLeft()
		}
	}
}
