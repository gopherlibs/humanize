package humanize

import (
	"strings"
	"unicode"
)

/* TestTokenize breaks out strings into a string slice. It removes special
 * characters, spaces, and gives you just the string slice.
 */
func tokenize(input string) []string {

	// exit early if we don't need to do anything
	if input == "" {
		return []string{}
	}

	// prep input
	input = strings.Trim(input, " ")
	input = strings.ToLower(input)

	var output []string    // what this function will return
	var sb strings.Builder // building out each word

	// loop through each character/rune from the input
	for _, r := range input {

		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		} else if sb.Len() != 0 {
			output = append(output, sb.String())
			sb.Reset()
		}
	}

	// there's a last word we need to finish
	if sb.Len() != 0 {
		output = append(output, sb.String())
	}

	return output
}

/* PascalCase converts a string into PascalCase.
 */
func PascalCase(input string) string {

	tokens := tokenize(input)

	var sb strings.Builder // building out each word

	// loop through tokens
	for _, t := range tokens {

		// capitalize first letter of word
		sb.WriteString(strings.ToUpper(string(t[0])))
		// apply the rest of the letters
		sb.WriteString(t[1:])
	}

	return sb.String()
}
