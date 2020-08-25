// +build js,wasm

package app

import "github.com/encryptgig/wasmapp/engine"

var CardList *engine.Element

var child map[string]*engine.Element

func CreateCardList() *engine.Element {

	e := engine.NewElement("div")
	child = make(map[string]*engine.Element)

	for _, j := range RoboArray {

		ch := CreateCard("https://robohash.org/"+j.Name, j.Name, j.Email)

		child[j.Email] = ch

		e.AddChild(ch)

	}

	CardList = e
	return e
}

func UpdateCardList(tmp []Robot) {
	CardList.RemoveChild()
	for _, j := range tmp {
		CardList.AddChild(child[j.Email])
	}
}
