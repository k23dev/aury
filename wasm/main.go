//go:build js && wasm

package main

import "syscall/js"

func main() {
	// now you expose your function to js
	js.Global().Set("Greet", greetWrapper())
	<-make(chan struct{})
}
