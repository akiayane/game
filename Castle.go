package main

type Castle struct {
	name string
	level int
	currentGold int
	income int
	friendly bool
}

var counter int32

func (c *Castle) upgrade() {
	c.level++
	c.income = c.income + 100
}
func (c *Castle) setFriendly(){
	c.friendly = true
}

func newCastle(name string) *Castle {
	Castle := &Castle{
		name: name,
		level: 1,
		currentGold: 100,
		income: 150,
		friendly: false,
	}
	return Castle
}

