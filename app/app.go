// +build js,wasm
package app

import (
	"syscall/js"

	"github.com/encryptgig/wasmapp/engine"
)

var AppObj js.Value

func CreateApp() *engine.Element {

	e := engine.NewElement("div").SetClass("tc")
	heading := engine.NewElement("h1").SetInnerHtml("Robo Friends").SetClass("f1")
	s := CreateSearchBox()
	e.AddChild(heading)
	loading := engine.NewElement("h1").SetInnerHtml("Loading...").SetClass("f1")
	e.AddChild(s)
	e.AddChild(loading)

	go FetchRobots(func() {
		e.Node.Call("removeChild", loading.Node)
		//e.SetInnerHtml("")
		scrl := CreateScroll()
		cl := CreateCardList()
		scrl.AddChild(cl)
		e.AddChild(scrl)
	})
	return e
}
