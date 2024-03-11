
//go:build js && wasm

package main

import "syscall/js"
	

func GreetWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		
		
		// Llamar a la función original de Go
		result := Greet(name)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}


func SumarWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		
		
		// Llamar a la función original de Go
		result := Sumar(a, b)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}


func DividirX3Wrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 3 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		
		
		// Llamar a la función original de Go
		result := DividirX3(a, b, c)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}

