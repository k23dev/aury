//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

// here you will define your wrappers for your functions

func greetingsWrapper() js.Func {
	// this returns a function
	// To return another type of data you must translate it to JS the js.ValueOf do that.
	// or just pass the the function js.FuncOf
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}
		input := args[0].String()
		fmt.Printf("input %s\n", input)
		pretty, err := Greetings(input)
		if err != nil {
			fmt.Printf("unable to convert to run the func %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
