
//go:build js && wasm

package main

import "syscall/js"

func main() {
	

	js.Global().Set("Greet", GreetWrapper())

	js.Global().Set("Sumar", SumarWrapper())

	js.Global().Set("DividirX3", DividirX3Wrapper())

	js.Global().Set("loguear", loguearWrapper())

	js.Global().Set("Prueba", PruebaWrapper())

	js.Global().Set("GetPepe", GetPepeWrapper())

	<-make(chan struct{})
}
	
