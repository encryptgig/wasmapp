// +build js,wasm

package app

import "github.com/encryptgig/wasmapp/engine"

func CreateScroll( )  *engine.Element  {
	e := engine.NewElement("div")
	e.Node.Get("style").Set("overflowY","scroll")
	e.Node.Get("style").Set("border","1px solid black")
	e.Node.Get("style").Set("height","800px")

	return e
}
