package utils

import (
	"bufio"
	"strconv"
	"strings"
)

func implode(intSlice []int, seperator string) string {
	var newString string = ""
	for _, value := range intSlice {
		newString += strconv.Itoa(value)
		newString += seperator
	}
	return newString[:len(newString)-1]
}

func stdInput(scanner *bufio.Scanner, varCount int) []string {
	var inputParameters []string

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	for i := 0; i < varCount; i++ {
		inputParameters = append(inputParameters, input[i])
	}
	return inputParameters
}

func convertInputToInt(variables []string) []int {
	var convertedInput []int
	for _, value := range variables {
		convertedVar, _ := strconv.Atoi(value)
		convertedInput = append(convertedInput, convertedVar)
	}
	return convertedInput
}
