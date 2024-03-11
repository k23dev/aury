//go:build js && wasm

package main

import "fmt"

func Greetings(name string) (string, error) {
	return fmt.Sprintf("Hello ,%s From Aury", name), nil
}
