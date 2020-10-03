//Package scanner ...
package scanner

import (
	"sort"
)

//GScanner - The scanner
type GScanner struct {
	input  string
	tokens []string
	itio   bool
}

type lengthWise []string

func NewScanner(input string, tokens []string, includeTokensInOutput bool) *GScanner {
	scanner := &GScanner{
		input:  input,
		tokens: tokens,
		itio:   includeTokensInOutput,
	}
	sort.Sort(lengthWise(scanner.tokens)) //sort tokens by length.. the longer appear before the shorter
	return scanner
}

//Scan - Scans the input into simple tokens.
func (cs *GScanner) Scan() []string {

	in := cs.input

	parse := make([]string, 0)

	for i := 0; i < len(in); i++ {

		for _, token := range cs.tokens {

			tokenLen := len(token)

			if tokenLen > 0 && i+tokenLen <= len(in) { //Ignore zero length tokens
				portion := in[i : i+tokenLen]

				if portion == token {
					if i != 0 { //avoid empty spaces
						parse = append(parse, in[0:i])
					}
					if cs.itio {
						parse = append(parse, token)
					}
					//in = in.substring(i + len)
					in = in[i+tokenLen:]
					i = -1
					break
				}

			}
		}

	}

	if len(in) > 0 {
		parse = append(parse, in)
	}

	return parse
}

func (s lengthWise) Len() int {
	return len(s)
}
func (s lengthWise) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s lengthWise) Less(i, j int) bool {
	return len(s[i]) > len(s[j]) // Sort in descending order
}
