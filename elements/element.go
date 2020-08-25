// +build js,wasm

package elements

import (
	"syscall/js"

	"github.com/encryptgig/wasmapp/engine"
)

type Style struct {
	Key   string
	Value string
}

type Action struct {
	Command string
	Func    js.Func
}

func CreateElement(tag string, innerhtml string, id string, class string, styles []Style, actions []Action, parent *engine.Element) *engine.Element {
	
	e := engine.NewElement("tag").SetId(id).SetClass(class)


	return e
}
