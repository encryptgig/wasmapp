// +build js,wasm

package app

import (
	"strings"
	"syscall/js"
	"github.com/rajnikant12345/wasmapp/engine"
)


func CreateSearchBox() *engine.Element {

	e := engine.NewElement("div")
	e.SetClass("pa2")
	inp := engine.NewElement("input")

	cb := js.NewCallback(func(args []js.Value) {
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

		//this should update card list
	})

	inp.SetCallBack("keyup",cb)

	inp.Set("type","search")
	inp.Set( "placeholder","search robots")
	inp.SetClass("pa3 ba b--green bg-lightest-blue")

	e.AddChild(inp)

	return e
}
