package private

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func formatName(fullName string) string {
	// Replace multiple spaces/tabs with a single space and trim leading/trailing spaces
	fields := strings.Fields(fullName)

	// Use a language-aware title case function
	caser := cases.Title(language.Russian) // You can adjust this for other languages

	// Convert each word to title case
	for i, word := range fields {
		fields[i] = caser.String(strings.ToLower(word))
	}

	// Join the formatted parts back into a single string
	return strings.Join(fields, " ")
}
