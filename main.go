// +build js,wasm

package main

import (
	"github.com/rajnikant12345/wasmapp/app"
	"github.com/rajnikant12345/wasmapp/engine"
)

func main() {
	c := make(chan struct{}, 1)

	println("WASM Go Initialized")

	engine.InitApp()
	engine.App.AddChild( app.CreateApp() )
	engine.StartApp()

	<-c
}
