package iteration

import "strings"

const repeatCount = 5

// Repeat takes a character and returns it repeated 5 times.
func Repeat(character string) string {
	var repeated strings.Builder

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}