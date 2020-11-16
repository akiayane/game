package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character interface {
	Attack(g *group)
	getStats()
	getName() string
	setName(name string)
	getHp() int
	setHp(int)
	getSpeed() int
	setSpeed(int)
	isStatus() bool
	setStatus(bool)
	isFriend() bool
	setFriend(bool)
	isEnemy() bool
	setEnemy(bool)
	getGold() int
	setGold(int)
	getWeapon() weapon
	equipWeapon(w weapon)
	CalcDefence() float32
	CalcDamage() float32
	calcCrit() bool
	levelUp()
}

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
	Friend  bool
	Gold    int

	Weapon weapon
}

func (m *character) getHp() int{
	//fmt.Println(m.Hp)
	return m.Hp
}

func (m *character) setHp (hp int) {
	m.Hp = hp
}

func (m *character) getSpeed() int{
	return m.Speed
}

func (m *character) setSpeed (speed int) {
	m.Speed = speed
}

func (m *character) getName() string{
	return m.Name
}

func (m *character) setName(name string) {
	m.Name = name
}

func (m *character) isStatus() bool {
	return m.Status
}

func (m *character) setStatus(status bool) {
	m.Status = status
}

func (m *character) isFriend() bool {
	return m.Friend
}

func (m *character) setFriend(friend bool) {
	m.Friend = friend
}

func (m *character) isEnemy() bool {
	return m.Enemy
}

func (m *character) setEnemy(enemy bool) {
	m.Enemy = enemy
}

func (m *character) getGold() int{
	return m.Gold
}

func (m *character) setGold (gold int) {
	m.Gold = gold
}

func (m *character) getWeapon() weapon {
	return m.Weapon
}

func (m *character) equipWeapon(w weapon) {
	m.Weapon = w
}

func (m *character) CalcDefence() float32 {
	var defense = m.Defence
	var defensePercent float32 = 3.0
	var tempDef float32 = 5.5
	var divider float32 = 1.11

	for i:=0; i<defense; i++{
		tempDef = tempDef/divider
		defensePercent += tempDef
	}

	return defensePercent
}

func (m *character) CalcDamage() float32 {
	if m.Weapon == nil {
		return float32(m.Damage)
	}
	return float32(m.Damage + m.Weapon.GetDamage())
}

func (m *character) calcCrit() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	var r = rand.Intn(100)
	if m.CritChance>=r {
		return true
	}
	return false
}

func (m *character) levelUp() {
	m.Level++
	m.Hp+=5
	m.Damage+=2
	m.Defence+=2
	m.Speed+=2
}

func (m *character)  getStats(){
	fmt.Println("Name - ",m.Name)
	fmt.Println("Hp - ",m.Hp)
	fmt.Println("Damage - ",m.Damage)
	fmt.Println("Defence - ",m.Defence)
	//fmt.Println("Crit Chance - ",m.CritChance)
	fmt.Println("Speed - ",m.Speed)
	fmt.Println("Price:", m.Gold)
	fmt.Println()
}


func (m *character) Attack(g *group) {

	var a int
	if m.Friend==true {
		fmt.Println("Choose target for ",m.Name)
		g.readGroup()
		fmt.Scanln(&a)
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		a=rand.Intn(g.Count())+1

		for g.cells[a-1].getHp() <= 0 {
			rand.Seed(time.Now().UTC().UnixNano())
			a=rand.Intn(g.Count())+1
		}

	}

var c = m.calcCrit()
var crit = 1
	switch a {
	case 1:
		if g.cells[0] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(m.CalcDamage()*((100-g.cells[0].CalcDefence())/100)*float32(crit))
			//fmt.Println(dd) // damage done
			g.cells[0].setHp(g.cells[0].getHp()-dd)
			fmt.Println(m.getName()," hits ",g.cells[0].getName()," for ",dd)
			if g.cells[0].getHp()==0 || g.cells[0].getHp() <0 {
				fmt.Println(g.cells[0].getName()," dies! ")
				g.cells[0].setStatus(false)
			}
		}
	case 2:
		if g.cells[1] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(m.CalcDamage()*((100-g.cells[1].CalcDefence())/100)*float32(crit))
			//fmt.Println(dd) // damage done
			g.cells[1].setHp(g.cells[1].getHp()-dd)
			fmt.Println(m.getName()," hits ",g.cells[1].getName()," for ",dd)
			if g.cells[1].getHp()==0 || g.cells[1].getHp() <0 {
				fmt.Println(g.cells[1].getName()," dies! ")
				g.cells[1].setStatus(false)
			}
		}
	case 3:
		if g.cells[2] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(m.CalcDamage()*((100-g.cells[2].CalcDefence())/100)*float32(crit))
			//fmt.Println(dd) // damage done
			g.cells[2].setHp(g.cells[2].getHp()-dd)
			fmt.Println(m.getName()," hits ",g.cells[2].getName()," for ",dd)
			if g.cells[2].getHp()==0 || g.cells[2].getHp() <0 {
				fmt.Println(g.cells[2].getName()," dies! ")
				g.cells[2].setStatus(false)
			}
		}
	case 4:
		if g.cells[3] != nil {
			if c {
				fmt.Println("CRITICAL HIT!")
				crit=2
			}
			dd:= int(m.CalcDamage()*((100-g.cells[3].CalcDefence())/100)*float32(crit))
			//fmt.Println(dd) // damage done
			g.cells[3].setHp(g.cells[3].getHp()-dd)
			fmt.Println(m.getName()," hits ",g.cells[3].getName()," for ",dd)
			if g.cells[3].getHp()==0 || g.cells[3].getHp() <0 {
				fmt.Println(g.cells[3].getName()," dies! ")
				g.cells[3].setStatus(false)
			}
		}
	default:
		fmt.Println("Incorrect input")
		m.Attack(g)
	}
}
