package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func generateWASMMain(goFilePath string, outputFile *os.File) error {
	goFile, err := os.Open(goFilePath)
	if err != nil {
		return err
	}
	defer goFile.Close()

	// Escanear el archivo Go y generar funciones wrapper
	if err := processGoFile(goFile, outputFile); err != nil {
		return err
	}

	return nil
}

func processGoMainFile(goFile *os.File, outputFile *os.File) error {
	// Mapa para almacenar funciones ya procesadas
	processedFunctions := make(map[string]bool)

	scanner := bufio.NewScanner(goFile)

	setMainFileHeader(outputFile)

	fmt.Println("----------------------------------------------")
	fmt.Printf(">>>>> %v \n", "a")
	for scanner.Scan() {
		line := scanner.Text()

		// Utilizar expresiones regulares para buscar declaraciones de funciones
		if matches := regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)\s*(\w*)`).FindStringSubmatch(line); matches != nil {
			functionName := matches[1]
			fmt.Printf(">>>>> %v \n", line)
			// Verificar si la función ya ha sido procesada
			if processedFunctions[functionName] {
				continue
			}

			// Procesar los parámetros y generar la función wrapper
			mainWrappers := generateMainFunction(functionName)
			fmt.Fprintln(outputFile, mainWrappers)

			// Marcar la función como procesada
			processedFunctions[functionName] = true
		}
	}
	setMainFileTail(outputFile)

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func generateMainFunction(functionName string) string {
	// Crear la función wrapper
	wrapperFunction := fmt.Sprintf("%s", generateMainBody(functionName))

	return wrapperFunction
}

func generateMainBody(functionName string) string {
	// Puedes personalizar la lógica interna del wrapper aquí según tus necesidades
	return fmt.Sprintf(`
	js.Global().Set("%s", %sWrapper())`, functionName, functionName)
}

func setMainFileHeader(outputFile *os.File) {
	header := `
//go:build js && wasm

package main

import "syscall/js"

func main() {
	`
	// file header
	fmt.Fprintln(outputFile, header)
}

func setMainFileTail(outputFile *os.File) {
	tail := `
	<-make(chan struct{})
}
	`
	// file header
	fmt.Fprintln(outputFile, tail)
}
