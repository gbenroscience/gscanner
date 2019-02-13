package main

import (
	"fmt"

	"github.com/gbenroscience/customscanner/scanner"
)

func main() {

	input := "(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+" +
		"(28+32+11-9E12+sin(3.2E9/cos(-3))-sinsinh(5)+sinh(8)+zopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekmzopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekmzopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekm-zopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekmzopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekmzopmkdmekdekekdmekdmekdmekdmekdmekdmekdmekdmekdmekdmekm"

	scanner := &scanner.CustomScanner{
		Input:                 input,
		Tokens:                []string{"-", "sin", "sinh", ""},
		IncludeTokensInOutput: true,
	}

	out := scanner.Scan()

	fmt.Println(out, "len > ", len(out))
}
