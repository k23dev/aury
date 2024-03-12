package main

import (
	"fmt"
	"regexp"
	"strings"
)

func processParameters(parameters string) string {
	// Dividir los parámetros teniendo en cuenta las comas dentro de los tipos de parámetros
	var tsParameters []string
	parenthesisCount := 0
	lastIndex := 0

	for i, char := range parameters {
		switch char {
		case '(':
			parenthesisCount++
		case ')':
			parenthesisCount--
		case ',':
			if parenthesisCount == 0 {
				// Dividir los parámetros
				tsParameters = append(tsParameters, parameters[lastIndex:i])
				lastIndex = i + 1
			}
		}
	}

	// Agregar el último parámetro
	tsParameters = append(tsParameters, parameters[lastIndex:])

	// Procesar cada parámetro
	var processedParameters []string
	for _, parameter := range tsParameters {
		// Utilizar expresiones regulares para extraer el nombre y el tipo del parámetro
		if matches := regexp.MustCompile(`(\w+)`).FindStringSubmatch(strings.TrimSpace(parameter)); matches != nil {
			parameterName := matches[1]

			// Asumir un tipo por defecto (puedes ajustar esto según tus necesidades)
			// parameterType := "any"
			parameterType := parseParamType(parameter)

			// Crear la parte de la declaración TypeScript para el parámetro
			tsParameter := fmt.Sprintf("%s: %s", parameterName, parameterType)
			processedParameters = append(processedParameters, tsParameter)
		}
	}

	// Unir los parámetros procesados
	return strings.Join(processedParameters, ", ")
}

func parseParamType(goType string) string {
	// Utilizar expresiones regulares para extraer el nombre y el tipo del parámetro
	if matches := regexp.MustCompile(`(\w+)\s+(\w+)`).FindStringSubmatch(strings.TrimSpace(goType)); matches != nil {
		_, parameterType := matches[1], matches[2]

		return converGo2TSType(parameterType)
	}

	// Tipo por defecto si no hay coincidencia
	return "any"
}

func converGo2TSType(goType string) string {
	switch goType {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "number"
	case "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	default:
		// Tipo por defecto si no hay coincidencia
		return "any"
	}
}

// parsea los parámetros de una función
func parseFunctionParameters(parameters string) string {
	splitedParams := strings.Split(parameters, ",")
	auxParamsIndex := 0
	var auxParams []string
	auxParamType := ""
	for i := 0; i < len(splitedParams); i++ {
		// remove spaces from the begening of the str
		splitedParams[i] = strings.TrimLeft(splitedParams[i], " ")
		if strings.Contains(splitedParams[i], " ") {
			splittedBySpace := strings.Split(splitedParams[i], " ")
			auxParams = append(auxParams, splittedBySpace[0])
			auxParamType = splittedBySpace[len(splittedBySpace)-1]
			// una vez que se encuentra el tipo de dato se le agrega a los parámetros anteriores
			auxParamsLen := len(auxParams)
			if auxParamsLen > 0 {
				for _, val := range auxParams {
					splitedParams[auxParamsIndex] = val + " " + auxParamType
					auxParamsIndex++
				}
				// resetea los valores
				auxParams = []string{}
				auxParamType = ""
			}
		} else {
			auxParams = append(auxParams, splitedParams[i])
		}
	}
	return strings.Join(splitedParams, ", ")
}

func parseFunctionParametersNoType(parameters string) string {

	splitedParams := strings.Split(parameters, ",")
	auxParams := []string{}
	for _, val := range splitedParams {
		param := strings.Split(strings.TrimLeft(val, " "), " ")
		auxParams = append(auxParams, param[0])
	}
	return strings.Join(auxParams, ", ")
}

func parseFunctionParameters4Typescript(parameters string) string {

	splitedParams := strings.Split(parameters, ",")
	auxParams := []string{}
	for _, val := range splitedParams {
		if val != "" {
			param := strings.Split(strings.TrimLeft(val, " "), " ")
			paramName := param[0]
			paramType := param[1]
			auxParam := paramName + " " + converGo2TSType(paramType)
			auxParams = append(auxParams, auxParam)
		}
	}
	return strings.Join(auxParams, ", ")
}
