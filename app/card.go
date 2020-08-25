// +build js,wasm

package app

import (
	"github.com/encryptgig/wasmapp/engine"
	"syscall/js"
)

func CreateCard(url, name, email string ) *engine.Element {

	cb := js.FuncOf( func(this js.Value, args []js.Value) interface{}{ js.Global().Call("alert", url) ; return nil; })

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
