// +build js,wasm

package app

import (
	"github.com/rajnikant12345/engine"
	"syscall/js"
)

func CreateCard(url, name, email string ) *engine.Element {

	cb := js.NewCallback(func(args []js.Value) { js.Global().Call("alert", url)})

	e := engine.NewElement("div").SetClass("bg-light-green dib br3 pa3 ma2 grow bw2 shadow-5")
	html :=
		`<img alt="robots" src=` + url + ` />
		<div>
                <h2>`+name+`</h2>
                <p>`+ email +`</p>
		</div>`
	e.SetInnerHtml(html).SetCallBack("click",cb)
	return e
}
