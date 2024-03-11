//go:build js && wasm

package main

import "syscall/js"

func main() {
	// now you expose your function to js
	js.Global().Set("greetings", greetingsWrapper())
	<-make(chan struct{})
}
