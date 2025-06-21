package main

import (
	"fmt"

	"github.com/gbenroscience/gscanner/scanner"
)

func printArray(arr []string) {

	numItems := len(arr)
	fmt.Print("[")
	for index, entry := range arr {
		if index < numItems-1 {
			fmt.Print(entry, ", ")
		} else {
			fmt.Print(entry, "]")
		}

	}
}

func main() {

	scan := scanner.NewScanner("(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)",
		[]string{"-", "sin", "sinh", "+", "(", ")", "cos", "/"},
		true)

	arr := scan.Scan()

	printArray(arr)

}
