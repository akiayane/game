package main

import "fmt"

func main() {
	clear()
	myMap := newWorldmap()
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
