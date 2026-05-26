// lowercase

package utils

import "strings"

func Tokenizer(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
