package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var goFilePath = "../../mygame/wasm/export_functions.go"
var tsFilePath = "../../frontend/src/aury.ts"

var funcWrappersFilePath = "../../mygame/wasm/wrappers_generated.go"

func main() {
	// Leer el archivo Go que contiene las funciones exportadas
	goFile, err := os.Open(goFilePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo Go:", err)
		return
	}
	defer goFile.Close()

	// Crear el archivo TypeScript para las declaraciones
	tsFile, err := os.Create(tsFilePath)
	if err != nil {
		fmt.Println("Error al crear el archivo TypeScript:", err)
		return
	}
	defer tsFile.Close()

	// Escanear el archivo Go y generar declaraciones TypeScript
	scanner := bufio.NewScanner(goFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Utilizar expresiones regulares más detalladas para buscar funciones exportadas y sus parámetros
		if matches := regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)\s*(\w*)`).FindStringSubmatch(line); matches != nil {
			functionName := matches[1]
			parameters := matches[2]
			returnType := matches[3]
			if matches[3] == "" {
				returnType = "void"
			} else {
				returnType = goToTSType(returnType)
			}

			// Procesar los parámetros y generar la declaración TypeScript
			tsParameters := processParameters(parameters)
			fmt.Fprintf(tsFile, "declare function %s(%s): %s;\n", functionName, tsParameters, returnType)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo Go:", err)
		return
	}

	fmt.Println("Declaraciones TypeScript generadas con éxito en", tsFilePath)

	// generate wrapper_functions

	outputFile, err := os.Create(funcWrappersFilePath)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer outputFile.Close()

	// Generar funciones wrapper desde el archivo Go
	if err := generateWrapperFunctions(goFilePath, outputFile); err != nil {
		fmt.Println("Error al generar funciones wrapper:", err)
		return
	}

	fmt.Println("Funciones wrapper generadas con éxito en", funcWrappersFilePath)

}
