package humanize

import (
	"strings"
	"unicode"
)

/* tokenize converts a string into a slice of strings. It
 * removes special characters, spaces, and gives you just the
 * string slice.
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

/* CamelCase converts a string into camelCase.
 *
 * For example, "`Happy birthday`` into `happyBirthday`."
 */
func CamelCase(input string) string {

	tokens := tokenize(input)

	var sb strings.Builder // building out each word

	// loop through tokens
	for i, t := range tokens {

		if i == 0 {
			sb.WriteString(t)
		} else {

			// capitalize first letter of word
			sb.WriteString(strings.ToUpper(string(t[0])))
			// apply the rest of the letters
			sb.WriteString(t[1:])
		}
	}

	return sb.String()
}

/* KebabCase converts a string into kebab-case.
 *
 * For example, "`Happy birthday`` into `happy-birthday`."
 */
func KebabCase(input string) string {

	tokens := tokenize(input)

	var sb strings.Builder // building out each word

	// loop through tokens
	for i, t := range tokens {

		if i > 0 {
			sb.WriteString("-")
		}

		sb.WriteString(t)
	}

	return sb.String()
}

/* PascalCase converts a string into PascalCase.
 *
 * For example, "`Happy birthday`` into `HappyBirthday`."
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
