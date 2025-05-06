package internal

import "strings"

func cleanupTokens(tokens []string) []string {
  var cleaned []string

  for _,t := range tokens {
    nt := strings.TrimSpace(t)
    if nt == "" {continue}
    cleaned = append(cleaned, nt)
  }
  return cleaned
}
