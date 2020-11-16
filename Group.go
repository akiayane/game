package main

import (
	"fmt"
	"sort"
)

type group struct{
	cells []Character
	counter int
	lvl int
}

func NewGroup(m Character) *group {
	Group:=&group{cells: []Character{m}}
	Group.counter=1
	Group.lvl = 1
	return Group
}

func (g *group) KickFromGroup(c Character) {

	g.cells = removeFromGroup(g.cells, c)
	g.counter--

}

func (g *group) AddToGroup(c Character) {
	if g.counter<4{
		if g.cells[0].getName()==mainName {
			c.setFriend(true)
			c.setName("My " + c.getName())
		}
		g.counter++
		for i := 1; i<g.lvl; i++ {
			c.levelUp()
		}
		g.cells=append(g.cells,c)
	} else {
		fmt.Println("Group already full!")
	}
}


func (g *group) Count() int {
	return g.counter //len(g.cells)
}

func (g *group) IsDead() bool {
	s := g.cells
	dedmembers:=0
	for  i :=0; i<len(s); i++ {
		if g.cells[i].isStatus()==false {
			dedmembers++
		}
	}
	if dedmembers==len(g.cells){
		return true
	} else {
		return false
	}
}

func(g *group) readGroup() {
	s := g.cells
	for  i :=0; i<len(s); i++ {
		if g.cells[i]!=nil&&g.cells[i].isStatus()!=false {
			fmt.Printf("%s[%d] ",s[i].getName(),i+1)
		}
	}
	fmt.Println()
}

func(g *group) levelUpGroup() {
	for _, char := range g.cells {
		char.levelUp()
	}
	g.lvl++
}

func(g *group) enterCastle(c *Castle) {
	if c.friendly {
		fmt.Print(" Friendly Castle")
		visitForTrade := &visitForTrade{}
		c.accept(visitForTrade, g)
	} else {
		fmt.Println(" Enemy Castle")
		visitForFight := &visitForFight{}
		c.accept(visitForFight, g)
	}
}

func(g1 *group) BattleStart(g2 *group) bool {
	fmt.Println("BATTLE STARTS!")
	fmt.Println()

	fmt.Println("Enemy group:")
	for _,s:= range g2.cells {
		s.getStats()
	}

	b := newBattle()
	n := g1.Count() + g2.Count()
	fmt.Println(n, " total number of participants from both sides")
	//var tmp,indox int
	var fighters []Character
	for h := 0; h < g1.Count(); h++ {
		fighters = append(fighters, g1.cells[h])
	}

	for h := 0; h < g2.Count(); h++ {
		fighters = append(fighters, g2.cells[h])
		fighters[g1.Count()+h].setEnemy(true)
	}
	var sorted []Character
	sorted = fighters
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].getSpeed() > sorted[j].getSpeed()
	})
	end := false
	for end != true {
		fmt.Println("ROUND ", b.round)
		for g := 0; g < len(sorted); g++ {
			if end == true {
				break
			} else {
				for u := g; u < len(sorted); u++ {
					if g1.IsDead() == true {
						fmt.Println("You lose")
						return false
					} else if g2.IsDead() == true {
						fmt.Println("You win")
						return true
					} else {
						if sorted[g].isStatus() != false {
							if sorted[g] == fighters[u] {
								if fighters[u].isEnemy() != true {
									fighters[u].Attack(g2)
								} else {
									fighters[u].Attack(g1)
								}
							}
						}
					}
				}
			}
		}
		b.nextRound()
	}
	return false
}

func removeFromGroup(List []Character, char Character) []Character {
	ListLength := len(List)
	for i, item := range List {
		if char.getName() == item.getName() {
			List[ListLength-1], List[i] = List[i], List[ListLength-1]
			return List[:ListLength-1]
		}
	}
	return List
}
