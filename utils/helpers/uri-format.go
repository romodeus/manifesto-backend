package helpers

import (
	"strings"
)

func URIFormat(words string) string {
	// convert to lowercase
	words = strings.ToLower(words)
	words = strings.ReplaceAll(words, " ", "-")
	return words
}
