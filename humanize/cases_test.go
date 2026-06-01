package humanize

import (
	"slices"
	"testing"
)

/* TestTokenize tests the Tokenize function.
 */
func TestTokenize(t *testing.T) {

	var testCases = []struct {
		name   string
		input  string
		output []string
	}{
		{
			"empty input",
			"",
			[]string{},
		},
		{
			"test a single word",
			"happy",
			[]string{"happy"},
		},
		{
			"test two words",
			"happy birthday",
			[]string{"happy", "birthday"},
		},
		{
			"a list of fruit",
			"apple banana cherry",
			[]string{"apple", "banana", "cherry"},
		},
		{
			"A normal sentence",
			"This is a normal sentence.",
			[]string{"this", "is", "a", "normal", "sentence"},
		},
	}

	// run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := tokenize(tc.input); !slices.Equal(result, tc.output) {
				t.Errorf("expected: %s, got: %s", tc.output, result)
			}
		})
	}
}

/* TestCamelCase tests the CamelCase function.
 */
func TestCamelCase(t *testing.T) {

	var testCases = []struct {
		name   string
		input  string
		output string
	}{
		{
			"empty input",
			"",
			"",
		},
		{
			"test a single word",
			"happy",
			"happy",
		},
		{
			"test two words",
			"happy birthday",
			"happyBirthday",
		},
		{
			"a list of fruit",
			"apple banana cherry",
			"appleBananaCherry",
		},
		{
			"A normal sentence",
			"This is a normal sentence.",
			"thisIsANormalSentence",
		},
	}

	// run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := CamelCase(tc.input); result != tc.output {
				t.Errorf("expected: %s, got: %s", tc.output, result)
			}
		})
	}
}

/* TestKebabCase tests the KebabCase function.
 */
func TestKebabCase(t *testing.T) {

	var testCases = []struct {
		name   string
		input  string
		output string
	}{
		{
			"empty input",
			"",
			"",
		},
		{
			"test a single word",
			"happy",
			"happy",
		},
		{
			"test two words",
			"happy birthday",
			"happy-birthday",
		},
		{
			"a list of fruit",
			"apple banana cherry",
			"apple-banana-cherry",
		},
		{
			"A normal sentence",
			"This is a normal sentence.",
			"this-is-a-normal-sentence",
		},
	}

	// run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := KebabCase(tc.input); result != tc.output {
				t.Errorf("expected: %s, got: %s", tc.output, result)
			}
		})
	}
}

/* TestPascalCase tests the PascalCase function.
 */
func TestPascalCase(t *testing.T) {

	var testCases = []struct {
		name   string
		input  string
		output string
	}{
		{
			"empty input",
			"",
			"",
		},
		{
			"test a single word",
			"happy",
			"Happy",
		},
		{
			"test two words",
			"happy birthday",
			"HappyBirthday",
		},
		{
			"a list of fruit",
			"apple banana cherry",
			"AppleBananaCherry",
		},
		{
			"A normal sentence",
			"This is a normal sentence.",
			"ThisIsANormalSentence",
		},
	}

	// run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := PascalCase(tc.input); result != tc.output {
				t.Errorf("expected: %s, got: %s", tc.output, result)
			}
		})
	}
}
