package main

import (
	"fmt"
	"sort"
)

type Battle struct {
	round int
}

type battle interface {
	nextRound()
}

func (b *Battle) nextRound()  {
	b.round++
}

func newBattle() *Battle {
	return &Battle{1}
}

func (g1 *group) BattleStart(g2 *group) bool {
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