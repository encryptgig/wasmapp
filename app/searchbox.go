// +build js,wasm

package app

import (
	"strings"
	"syscall/js"
	"github.com/encryptgig/wasmapp/engine"
)


func CreateSearchBox() *engine.Element {

	e := engine.NewElement("div").SetClass("pa2")
	inp := engine.NewElement("input")

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		user := e.Get("target").Get("value").String()
		println(user)
		if user != "" {
			tmp := make([]Robot,0)
			for _,i:= range RoboArray {
				if strings.Contains(strings.ToLower(i.Name) , strings.ToLower(user)) {
					tmp = append(tmp , i)
				}
			}
			UpdateCardList(tmp)
			//RoboArray = tmp
		}else {
			UpdateCardList(RoboArray)
		}
		return nil
		//this should update card list
	})

	inp.SetCallBack("keyup",cb).Set("type","search").Set( "placeholder","search robots").SetClass("pa3 ba b--green bg-lightest-blue")

	e.AddChild(inp)

	return e
}
