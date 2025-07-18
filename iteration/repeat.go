package iteration

import "strings"

var repeatCount int

// Takes a string and an int-value, repeats back the string x-amout of times where x = entered int-value
func Repeat(character string, x int) string {
	repeatCount = x
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
