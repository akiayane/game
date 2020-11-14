package main

import (
	"fmt"
	"sort"
)

type group struct{
	cells []*character
	counter int
}

func  NewGroup(m *character) *group {
	Group:=&group{cells: []*character{m}}
	Group.counter=1
	return Group
}

func (g *group) KickFromGroup(c *character) {
	for  i :=0; i<len(g.cells); i++ {
		if g.cells[i]==c {
			g.cells[i]=nil
			break
		}
	}
}

func (g *group) AddToGroup(c *character) {
	if g.counter<4{
		g.counter++
		g.cells=append(g.cells,c)
	} else {
		fmt.Println("Group already full!")
	}
}

func (g *group) Count() int {
	return len(g.cells)
}

func (g *group) IsDead() bool {
	s := g.cells
	dedmembers:=0
	for  i :=0; i<len(s); i++ {
		if g.cells[i].Status==false {
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
		if g.cells[i]!=nil&&g.cells[i].Status!=false {
			fmt.Printf("%s[%d] ",s[i].getName(),i+1)
		}
	}
	fmt.Println()
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
	b:=newBattle()
	n:=g1.Count()+g2.Count()
	fmt.Println(n," total number of participants from both sides")
	//var tmp,indox int
	var fighters []*character
	for h:=0; h<g1.Count(); h++{
		fighters=append(fighters,g1.cells[h])
	}
	for h:=0; h<g2.Count(); h++{
		fighters=append(fighters,g2.cells[h])
		fighters[h].Enemy=true
	}
	var sorted []*character
	sorted=fighters
	sort.SliceStable(sorted, func(i,j int) bool {
		return sorted[i].Speed > sorted[j].Speed
	})
	end:=false
	for end!=true{
		fmt.Println("ROUND ",b.round)
		for g:=0; g<len(sorted); g++{
			if end==true {break} else {
				for u:=g; u<len(sorted); u++{
					if g1.IsDead()==true {
						fmt.Println("You lose")
						return false
					} else if g2.IsDead()==true {
						fmt.Println("You win")
						return true
					} else {
						if sorted[g].Status!=false {
							if sorted[g]==fighters[u]{
								if fighters[u].Enemy==true {
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
