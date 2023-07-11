package main

import (
	"syscall/js"
)

func hello(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("Hello wasm!!")
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("hello", js.FuncOf(hello))
	// main関数が終了しないようにstructのチャネルを受信し続ける
	<-c
}
