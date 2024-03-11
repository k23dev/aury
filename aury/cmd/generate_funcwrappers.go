package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func generateWrapperFunctions(goFilePath string, outputFile *os.File) error {
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

func processGoFile(goFile *os.File, outputFile *os.File) error {
	// Mapa para almacenar funciones ya procesadas
	processedFunctions := make(map[string]bool)

	scanner := bufio.NewScanner(goFile)

	setFileHeader(outputFile)

	for scanner.Scan() {
		line := scanner.Text()

		// Utilizar expresiones regulares para buscar declaraciones de funciones
		if matches := regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)\s*(\w*)`).FindStringSubmatch(line); matches != nil {
			functionName := matches[1]

			// Verificar si la función ya ha sido procesada
			if processedFunctions[functionName] {
				continue
			}

			parameters := matches[2]
			returnType := matches[3]

			// Procesar los parámetros y generar la función wrapper
			wrapperFunction := generateWrapperFunction(functionName, parameters, returnType)
			fmt.Fprintln(outputFile, wrapperFunction)

			// Marcar la función como procesada
			processedFunctions[functionName] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func generateWrapperFunction(functionName, parameters, returnType string) string {
	parameters = parseFunctionParameters(parameters)
	tsParameters := parseFunctionParameters4Typescript(parameters)
	auxParametersNoType := parseFunctionParametersNoType(parameters)

	// Crear la función wrapper
	wrapperFunction := fmt.Sprintf(`
func %sWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != %d {
			return "Invalid number of arguments passed"
		}
		
		%s
	})
}
`, functionName, len(strings.Split(tsParameters, ",")), generateWrapperBody(functionName, tsParameters, auxParametersNoType, returnType))

	return wrapperFunction
}

func generateWrapperBody(functionName, tsParameters, parametersNoType, returnType string) string {
	// Puedes personalizar la lógica interna del wrapper aquí según tus necesidades
	return fmt.Sprintf(`
		// Extraer los argumentos
		%s
		
		// Llamar a la función original de Go
		result := %s(%s)
		
		// Puedes traducir el resultado a JS si es necesario
		return result
`, generateExtractArguments(tsParameters), functionName, parametersNoType)
}

func generateExtractArguments(tsParameters string) string {
	fmt.Println("GENERATE!!!!")
	fmt.Printf("TS >> \"%s\" \n", tsParameters)
	var extractArguments []string
	for index, param := range strings.Split(tsParameters, ",") {
		paramParts := strings.Split(param, ":")
		if len(paramParts) > 1 {
			paramName := strings.TrimSpace(paramParts[0])
			paramType := strings.TrimSpace(paramParts[1])
			extractArguments = append(extractArguments, fmt.Sprintf("%s := args[%d].%s()", paramName, index, generateGoTypeConversion(paramType)))
		}
		// extractArguments = append(extractArguments, fmt.Sprintf("%s := args[%d].%s()", paramName, index, generateGoTypeConversion(paramType)))
	}

	return strings.Join(extractArguments, "\n\t\t")
}

func generateGoTypeConversion(paramType string) string {
	// Agrega conversiones de tipo según sea necesario
	// Puedes personalizar esto según los tipos que estás utilizando en tus funciones Go
	switch paramType {
	case "int":
		return "Int"
	case "int8":
		return "Int8"
	case "int16":
		return "Int16"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint":
		return "Uint"
	case "uint8":
		return "Uint8"
	case "uint16":
		return "Uint16"
	case "uint32":
		return "Uint32"
	case "uint64":
		return "Uint64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "string":
		return "String"
	case "bool":
		return "Bool"
	default:
		// Tipo por defecto
		return "Any"
	}
}

func setFileHeader(outputFile *os.File) {
	header := `
//go:build js && wasm

package main

import "syscall/js"
	`
	// file header
	fmt.Fprintln(outputFile, header)
}
