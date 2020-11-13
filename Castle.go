package main

type Castle struct{
	name string
	level int
	currentGold int
	income int
	friendly bool
	group group
}


func (c *Castle) upgrade() {
	c.level++
	c.income = c.income + 100
}
func (c *Castle) setFriendly(){
	c.friendly = true
}

func (c *Castle) setGroup(){
	
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

