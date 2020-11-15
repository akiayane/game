package main

import (
	"fmt"
	"math/rand"
)

type subject interface {
	register(Observer *HeroIcon)
	deregister(Observer *HeroIcon)
	notifyAll()
}

type weatherCondition struct {
	observerList []*HeroIcon
	state int
}

func (w *weatherCondition) updateState() {
	r := rand.Intn(4)
	
	condition := weather[r]

	fmt.Println("Weather today:", condition)

	switch condition {
	case "sun":
		w.state = 2
	case "normal":
		w.state = 0
	case "rain":
		w.state = -2
	case "snow":
		w.state = -5
	}

	w.notifyAll()
}

func (w *weatherCondition) notifyAll() {
	for _, observer := range w.observerList {
		observer.update(w.state)
	}
}

func (w *weatherCondition) register(o *HeroIcon) {
	w.observerList = append(w.observerList, o)
}

func (w *weatherCondition) deregister(o *HeroIcon) {
	w.observerList = removeFromslice(w.observerList, o)
}

func removeFromslice(observerList []*HeroIcon, observerToRemove *HeroIcon) []*HeroIcon {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.group.cells[0].getName() == observer.group.cells[0].getName() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
