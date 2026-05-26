// Package utils provides core utility functions for the search engine.
package utils

import "strings"

// stopWords is a map of common English stop words that are excluded from the index.
var stopWords = map[string]struct{}{
	"a":    {},
	"and":  {},
	"be":   {},
	"have": {},
	"i":    {},
	"in":   {},
	"of":   {},
	"that": {},
	"the":  {},
	"to":   {},
	"was":  {},
	"with": {},
	"is":   {},
	"for":  {},
	"on":   {},
	"it":   {},
	"as":   {},
	"at":   {},
	"by":   {},
	"an":   {},
	"this": {},
	"are":  {},
	"but":  {},
	"not":  {},
}

// Filter removes stop words and performs simple stemming on the list of tokens.
func Filter(tokens []string) []string {
	var filtered []string
	for _, token := range tokens {
		if _, exists := stopWords[token]; !exists {
			stemmed := stem(token)
			if len(stemmed) > 0 {
				filtered = append(filtered, stemmed)
			}
		}
	}
	return filtered
}

// stem performs a basic, naive suffix-stripping stemming (e.g. fishing -> fish).
func stem(token string) string {
	if len(token) > 4 && strings.HasSuffix(token, "ing") {
		return token[:len(token)-3]
	}
	if len(token) > 3 && strings.HasSuffix(token, "ed") {
		return token[:len(token)-2]
	}
	if len(token) > 2 && strings.HasSuffix(token, "s") && !strings.HasSuffix(token, "ss") {
		return token[:len(token)-1]
	}
	return token
}