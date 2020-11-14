package main

import (
	"fmt"
	"math/rand"
	"time"
)

type character struct {
	Name    string
	Hp      int
	Damage  int
	Defence int
	CritChance int
	Speed   int
	Enemy   bool
	Level   int
	Status  bool
}

type Character interface {
	Attack(g *group)
	getStats()
}

func (m *character) Attack(g *group) {
var a int
fmt.Println("Choose target for ",m.Name)
	g.readGroup()
fmt.Scanln(&a)
var c = m.calcCrit()
var crit = 1
	switch a {
	case 1:
		if g.cells[0] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(float32(m.Damage)*((100-g.cells[0].CalcDefence())/100)*float32(crit))
			//fmt.Println(dd) // damage done
			g.cells[0].Hp=g.cells[0].Hp-dd
			fmt.Println(m.getName()," hits ",g.cells[0].getName()," for ",dd)
			if g.cells[0].Hp==0 || g.cells[0].Hp <0 {
				fmt.Println(g.cells[0].getName()," dies! ")
				g.cells[0].Status=false
			}
		}
	case 2:
		if g.cells[1] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(float32(m.Damage)*((100-g.cells[1].CalcDefence())/100)*float32(crit))
			g.cells[1].Hp=g.cells[1].Hp-dd
			fmt.Println(m.getName()," hits ",g.cells[1].getName()," for ",dd)
			if g.cells[1].Hp==0 || g.cells[1].Hp <0 {
				fmt.Println(g.cells[1].getName()," dies! ")
				g.cells[1].Status=false
			}
		}
	case 3:
		if g.cells[2] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(float32(m.Damage)*((100-g.cells[2].CalcDefence())/100)*float32(crit))
			g.cells[2].Hp=g.cells[2].Hp-dd
			fmt.Println(m.getName()," hits ",g.cells[2].getName()," for ",dd)
			if g.cells[2].Hp==0 || g.cells[2].Hp <0 {
				fmt.Println(g.cells[2].getName()," dies! ")
				g.cells[2].Status=false
			}
		}
	case 4:
		if g.cells[3] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(float32(m.Damage)*((100-g.cells[3].CalcDefence())/100)*float32(crit))
			g.cells[3].Hp=g.cells[3].Hp-dd
			fmt.Println(m.getName()," hits ",g.cells[3].getName()," for ",dd)
			if g.cells[3].Hp==0 || g.cells[3].Hp <0 {
				fmt.Println(g.cells[3].getName()," dies! ")
				g.cells[3].Status=false
			}
		}
	default:
		fmt.Println("Incorrect input")
		m.Attack(g)
	}
}

func (m *character) CalcDefence() float32 {
	var defense int = m.Defence
	var defensePercent float32 = 3.0
	var tempDef float32 = 5.5
	var divider float32 = 1.11

	for i:=0; i<defense; i++{
		tempDef = tempDef/divider
		defensePercent += tempDef
	}

	return defensePercent
}

func (m *character) calcCrit() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	var r = rand.Intn(100)
	if m.CritChance>=r {
		return true
	}
	return false
}

func (m *character)  getStats(){
fmt.Println("Name - ",m.Name)
fmt.Println("Hp - ",m.Hp)
fmt.Println("Damage - ",m.Damage)
fmt.Println("Defence - ",m.Defence)
fmt.Println("Crit Chance - ",m.CritChance)
fmt.Println("Speed - ",m.Speed)
}

func (m *character) getHp() {
	fmt.Println(m.Hp)
}

func (m *character)  getName() string{
	return m.Name
}

func (m *character)  equipWeapon(w weapon) {
	m.Damage=m.Damage + w.GetDamage()
}

/////////////////////////////////////BUILDER////////////////////////////
func (cb *Builder) setName(name string) CharBuilder {
	cb.name = name
	return cb
}

func (cb *Builder) setHp(num int) CharBuilder {
	cb.hp=num
	return cb
}

func (cb *Builder) setDamage(num int) CharBuilder {
	cb.damage=num
	return cb
}

func (cb *Builder) setDefence(num int) CharBuilder {
	cb.defence=num
	return cb
}

func (cb *Builder) setCritChance(num int) CharBuilder {
	cb.critchance=num
	return cb
}

func (cb *Builder) setSpeed(num int) CharBuilder {
	cb.speed=num
	return cb
}

func (cb *Builder) setStatus(num bool) CharBuilder {
	cb.status=num
	return cb
}

type CharBuilder interface {
	setName(name string) CharBuilder
	setHp(num int) CharBuilder
	setDamage(num int) CharBuilder
	setDefence(num int) CharBuilder
	setSpeed(num int) CharBuilder
	setStatus(num bool) CharBuilder
	Create() Character
}

type Builder struct {
	name string
	level int
	hp int
	damage int
	defence int
	critchance int
	speed int
	status bool
}

func NewCharBuilder() *Builder {
	return &Builder{}
}

func (cb *Builder) Create() Character {
	return &character{
		cb.name,
		cb.hp,
		cb.damage,
		cb.defence,
		cb.critchance,
		cb.speed,
		false,
		1,
		true,
	}
}
/////////////////////////////////////WEAPONS////////////////////////////