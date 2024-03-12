
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


func SumarWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		a := args[0].Int()
		b := args[1].Int()
		
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
		a_aux := args[0].Int()
		a := int16(a_aux)
		b_aux := args[1].Int()
		b := int16(b_aux)
		c_aux := args[2].Float()
		c := float32(c_aux)
		
		// Llamar a la función original de Go
		result := DividirX3(a, b, c)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}


func loguearWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		
		
		// Llamar a la función original de Go
		result := loguear()
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}


func PruebaWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 8 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		a := args[0].Int()
		b := args[1].Int()
		c := args[2].Int()
		d := args[3].Int()
		e_aux := args[4].Float()
		e := float32(e_aux)
		z := args[5].String()
		x := args[6].String()
		w := args[7].String()
		
		// Llamar a la función original de Go
		result := Prueba(a, b, c, d, e, z, x, w)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}


func GetPepeWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		
		
		// Extraer los argumentos
		a := args[0]
		
		// Llamar a la función original de Go
		result := GetPepe(a)
		
		// Puedes traducir el resultado a JS si es necesario
		return result

	})
}

