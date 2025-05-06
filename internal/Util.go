package internal

import (
	"fmt"
	"strings"
)

func cleanupTokens(tokens []string) []string {
  var cleaned []string

  for _,t := range tokens {
    nt := strings.TrimSpace(t)
    if nt == "" {continue}
    cleaned = append(cleaned, nt)
  }
  return cleaned
}

func ExtractParenthesisContent(s string) (string, error) {
	startIndex := strings.Index(s, "(")
	if startIndex == -1 {
		return "", fmt.Errorf("no opening parenthesis found")
	}

	endIndex := strings.Index(s, ")")
	if endIndex == -1 {
		return "", fmt.Errorf("no closing parenthesis found")
	}

	if endIndex < startIndex {
		return "", fmt.Errorf("unbalanced parentheses: closing before opening")
	}

	return s[startIndex+1 : endIndex], nil
}

func ExtractQuotedString(s string) (string, bool) {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1], true
		}
		if s[0] == '\'' && s[len(s)-1] == '\'' {
			return s[1 : len(s)-1], true 
		}
	}
	return s, false
}
