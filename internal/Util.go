package internal

import (
	"bufio"
	"fmt"
	"os"
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

func Input() string {
  reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
  return strings.TrimSpace(res)
}

func mathOperations(symbol string, num1,num2 int) (int, error) {
  switch (symbol) {
    case "+":
      return num1 + num2, nil
    case "-":
      return num1 - num2, nil
    case "*":
      return num1 * num2, nil
    case "/":
      return num1 / num2, nil
  }
  return 0, fmt.Errorf("Unsupported math operation %s", symbol)
}
