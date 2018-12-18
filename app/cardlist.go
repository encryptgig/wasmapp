// +build js,wasm

package app

import "github.com/rajnikant12345/wasmapp/engine"

var CardList *engine.Element

func CreateCardList() *engine.Element {

	e := engine.NewElement("div")

	for _, j := range RoboArray {
		e.AddChild(CreateCard("https://robohash.org/"+j.Name, j.Name, j.Email))
	}

	CardList = e
	return e
}

func UpdateCardList(tmp []Robot) {
	CardList.RemoveChild()
	for _, j := range tmp {
		CardList.AddChild(CreateCard("https://robohash.org/"+j.Name,j.Name,j.Email))
	}
}
