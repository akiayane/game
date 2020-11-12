package main

import "fmt"

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
