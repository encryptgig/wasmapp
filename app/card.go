package app

import (
	"github.com/rajnikant12345/wasmapp/engine"
	"syscall/js"
)

func CreateCard(url, name, email string ) *engine.Element {

	e := engine.NewElement("div")
	e.SetClass("bg-light-green dib br3 pa3 ma2 grow bw2 shadow-5")
	html :=
		`<img alt="robots" src=` + url + ` />
		<div>
                <h2>`+name+`</h2>
                <p>`+ email +`</p>
		</div>`
	e.SetInnerHtml(html)
	cb := js.NewCallback(func(args []js.Value) { js.Global().Call("alert", url)})

	e.SetCallBack("click",cb)
	return e
}
