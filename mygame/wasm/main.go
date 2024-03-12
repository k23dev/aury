
//go:build js && wasm

package main

import "syscall/js"

func main() {
	

	js.Global().Set("Greet", GreetWrapper())

	<-make(chan struct{})
}
	
