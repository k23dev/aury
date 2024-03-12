
//go:build js && wasm

package main

import "syscall/js"
	

func GreetWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		name := args[0].String()
		
		// Llamar a la función original de Go
		result := Greet(name)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}

