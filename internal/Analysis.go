package internal

import "strings"

func (program *Program) Interpret(mainContent string) {
  lines := strings.Split(mainContent, "\n")

  for _,line := range lines {
    cleanupLine(line)
  }
}

func cleanupLine(line string) {
  tokens := strings.Split(line, " ")
  tokens = cleanupTokens(tokens)
  if len(tokens) == 0 {return}
  if strings.HasPrefix(tokens[0], "//") {return} 

  interpretLine(tokens)
}

func interpretLine(tokens []string) {
  for _,i := range tokens {
    println(i)
  }
}
