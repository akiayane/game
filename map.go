package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
	"os/exec"
)

type HeroIcon struct {
	currentX int
	currentY int
	previousX int
	previousY int
	icon string
	speed int
	movesleft int

	group *group
	gold int
}

type CastleIcon struct{
	locX int
	locY int
	doorsX int
	doorsY int
	icon string

	castle *Castle
}

func newCastleIcon(castle Castle, upwards bool, x int, y int) *CastleIcon{
	var icon string
	if castle.friendly{
		icon = "H"
	}else{
		icon = "X"
	}
	if upwards{
		return  &CastleIcon{x, y, x, y-1, icon, &castle}
	}else{
		return  &CastleIcon{x, y, x, y+1, icon, &castle}
	}
}

type worldmap struct {
	x int
	y int
	myMap [140][35]string
	HeroIcon HeroIcon
	empty string
	casSlice []CastleIcon
	days int
	weeks int
}

func (w *worldmap) checkdoors(){
	for i := 0; i<len(w.casSlice); i++{
		if w.myMap[w.casSlice[i].doorsX][w.casSlice[i].doorsY] == w.HeroIcon.icon{
			clear()
			for{
				x := strconv.Itoa(w.casSlice[i].doorsX)
				y := strconv.Itoa(w.casSlice[i].doorsY)
				fmt.Printf("Current position: %s, %s", x, y )
				//w.HeroIcon.group.cells[0].gold = w.HeroIcon.gold
				w.HeroIcon.group.enterCastle(w.casSlice[i].castle)

					clear()
					if w.casSlice[i].locY - w.casSlice[i].doorsY == 1{
						w.moveUp()
						break
					}else{
						w.moveDown()
						break
					}

			}
			}
	}
}

func newWorldmap() *worldmap{
	var myMap = [140][35]string{}
	mapp := readLines("txt.txt")
	for i:=0; i<35; i++{
		for j:= 0; j<140; j++{

			myMap[j][i] = string(mapp[i][j])
		}
	}
	Alibek:= newMainCharacter("Alibek")
	Alibek.setGold(200)
	mainName = Alibek.getName()
	Group := NewGroup(Alibek)
	//Group.AddToGroup(newGoblin())
	//Group.AddToGroup(newOrk())
	//Group.AddToGroup(newOrk())
	Hero := HeroIcon{10,4,10,4, "@", Alibek.getSpeed()+3, Alibek.getSpeed()+4,Group, 100}

	return &worldmap{140, 35, myMap,  Hero, ".", initialize(), 0,0}
}

func (hi *HeroIcon) update(state int) {

}

func (w *worldmap) updStats(){
	if w.HeroIcon.movesleft == 0{
		w.nextday()
		w.HeroIcon.movesleft = w.HeroIcon.speed
	}
	w.HeroIcon.movesleft--
	movesLeft := strconv.Itoa(w.HeroIcon.movesleft)
	speed:= strconv.Itoa(w.HeroIcon.speed)
	day := strconv.Itoa(w.days)
	weeks := strconv.Itoa(w.weeks)
	gold := strconv.Itoa(w.HeroIcon.group.cells[0].getGold())
	lvl := strconv.Itoa(w.HeroIcon.group.lvl)
	fmt.Printf("Moves left: %s | Your speed: %s | Days: %s | Weeks : %s | Your Gold : %s | Your level : %s", movesLeft,speed,day,weeks,gold,lvl)
	fmt.Print("\n")
}

func (w *worldmap) nextday(){
	w.days++
	if w.days%7==0{
		w.weeks++
	}
	dailyGold := 0
	for i:=0; i<len(w.casSlice); i++{
		if w.casSlice[i].castle.friendly{
			dailyGold += w.casSlice[i].castle.income
		}else{
			w.casSlice[i].castle.currentGold = w.casSlice[i].castle.currentGold + w.casSlice[i].castle.income
		}
	}
	w.HeroIcon.group.cells[0].setGold(w.HeroIcon.group.cells[0].getGold() + dailyGold)
}


func (w *worldmap) fillMap(){
	var i,j int
	for  i = 0; i < w.x; i++ {
		for j = 0; j < w.y; j++ {
			w.myMap[i][j] = w.empty
		}
	}
}

func (w *worldmap) printMap(){
	w.updStats()
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

func (w *worldmap) cage(){

	for i:= 0; i<w.x; i++{
		for j:= 0; j<w.y; j++{
			switch w.myMap[i][j] {
			case "Q":
				w.myMap[i][j] = "╔"
			case "W":
				w.myMap[i][j] = "╗"
			case "E":
				w.myMap[i][j] = "╚"
			case "R":
				w.myMap[i][j] = "╝"
			case "Y":
				w.myMap[i][j] = "║"
			case "T":
				w.myMap[i][j] = "═"
			}
		}
	}
	w.printMap()
}

func (w *worldmap) getCoor(){
	for i:= 0; i<w.x; i++{
		for j:= 0; j<w.y; j++{
			if w.myMap[i][j] == "H" || w.myMap[i][j] == "X"{
				fmt.Println(i, j)
			}
		}
	}
}

func (w *worldmap) moveRight(){
	if w.myMap[w.HeroIcon.currentX+1][w.HeroIcon.currentY] != w.empty{
		w.printMap()
	}else if w.HeroIcon.currentX + 1 == w.x-1 {
		w.printMap()
	}else{
		w.HeroIcon.previousX = w.HeroIcon.currentX
		w.HeroIcon.previousY = w.HeroIcon.currentY
		w.HeroIcon.currentX = w.HeroIcon.currentX + 1
		w.myMap[w.HeroIcon.previousX][w.HeroIcon.previousY] = w.empty
		w.printMap()
	}
}

func (w *worldmap) moveLeft(){
	if w.myMap[w.HeroIcon.currentX-1][w.HeroIcon.currentY] != w.empty{
		w.printMap()
	}else if w.HeroIcon.currentX - 1 == 0 {
		w.printMap()
	}else {
		w.HeroIcon.previousX = w.HeroIcon.currentX
		w.HeroIcon.previousY = w.HeroIcon.currentY
		w.HeroIcon.currentX = w.HeroIcon.currentX - 1
		w.myMap[w.HeroIcon.previousX][w.HeroIcon.previousY] =w.empty
		w.printMap()
	}
}

func (w *worldmap) moveUp(){
	if w.myMap[w.HeroIcon.currentX][w.HeroIcon.currentY-1] != w.empty{
		w.printMap()
	}else if w.HeroIcon.currentY - 1 == 0 {
		w.printMap()
	}else {
		w.HeroIcon.previousX = w.HeroIcon.currentX
		w.HeroIcon.previousY = w.HeroIcon.currentY
		w.HeroIcon.currentY = w.HeroIcon.currentY - 1
		w.myMap[w.HeroIcon.previousX][w.HeroIcon.previousY] = w.empty
		w.printMap()
	}
}

func (w *worldmap) moveDown(){
	if w.myMap[w.HeroIcon.currentX][w.HeroIcon.currentY+1] != w.empty{
		w.printMap()
	}else if w.HeroIcon.currentY + 1 == w.y-1 {
		w.printMap()
	}else {
		w.HeroIcon.previousX = w.HeroIcon.currentX
		w.HeroIcon.previousY = w.HeroIcon.currentY
		w.HeroIcon.currentY = w.HeroIcon.currentY + 1
		w.myMap[w.HeroIcon.previousX][w.HeroIcon.previousY] = w.empty
		w.printMap()
	}
}

func clear(){
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLines(path string) ([]string) {
	file, _ := os.Open(path)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}




