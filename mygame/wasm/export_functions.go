//go:build js && wasm

package main

import "fmt"

// your functions comes here
func Greet(name string) string {
	return fmt.Sprintf("Hello ,%s from Aury", name)
}
