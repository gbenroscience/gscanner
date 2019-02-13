//Package scanner ...
package scanner

import (
	"sort"
)

//CustomScanner - The scanner
type GScanner struct {
	Input                 string
	Tokens                []string
	IncludeTokensInOutput bool
}

type lengthWise []string

//Scan - Scans the input into simple tokens.
func (cs *GScanner) Scan() []string {

	in := cs.Input

	parse := make([]string, 0)

	sort.Sort(lengthWise(cs.Tokens))

	for i := 0; i < len(in); i++ {

		for _, token := range cs.Tokens {

			strLen := len(token)

			if strLen > 0 && i+strLen <= len(in) { //Ignore zero length tokens
				portion := in[i : i+strLen]

				if portion == token {
					if i != 0 { //avoid empty spaces
						parse = append(parse, in[0:i])
					}
					if cs.IncludeTokensInOutput {
						parse = append(parse, token)
					}
					//in = in.substring(i + len)
					in = in[i+strLen:]
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
