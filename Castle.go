package main

type Castle struct{
	name string
	level int
	income int
	friendly bool
	group *group
}

func (c *Castle) upgrade() {
	c.level++
	c.income = c.income + 10
	if !c.friendly {
		c.group.levelUpGroup()
	}
}

func (c *Castle) setFriendly(){
	c.friendly = true
}

func (c *Castle) setGroup(group *group){
	c.group = group
}

func newCastle(name string) *Castle {
	Castle := &Castle{
		name: name,
		level: 1,
		income: 50,
		friendly: false,
	}
	return Castle
}
func (c *Castle) accept(v visitor, group *group) {
	v.visitCastle(c, group)
}

