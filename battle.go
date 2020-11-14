package main

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
