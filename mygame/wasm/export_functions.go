//go:build js && wasm

package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello ,%s from Aury", name)
}

func Sumar(a, b int) int {
	return a + b
}

func DividirX3(a, b int16, c float32) int16 {
	return a + b
}

func loguear() interface{} {
	return nil
}

func Prueba(a, b, c, d int, e float32, z, x, w string) int {
	return a + b
}

func GetPepe(a interface{}) interface{} {
	return a
}
