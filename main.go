// +build js,wasm

package main

import (
	"github.com/encryptgig/wasmapp/app"
	"github.com/encryptgig/wasmapp/engine"
)

func main() {
	c := make(chan struct{}, 1)

	println("WASM Go Initialized")

	engine.InitApp()
	engine.App.AddChild( app.CreateApp() )
	engine.StartApp()

	<-c
}
