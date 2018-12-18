package app

import "github.com/rajnikant12345/wasmapp/engine"
import "syscall/js"

var AppObj js.Value

func CreateApp() *engine.Element {

	FetchRobots()

	e := engine.NewElement("div")
	e.SetClass("tc")
	heading := engine.NewElement("h1")
	heading.SetInnerHtml("Robo Friends")
	heading.SetClass("f1")
	s := CreateSearchBox()
	e.AddChild(heading)
	e.AddChild(s)
	scrl := CreateScroll()
	cl := CreateCardList()
	scrl.AddChild(cl)
	e.AddChild(scrl)

	return e
}



