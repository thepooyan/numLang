package internal

import (
	"os"
	"strings"
)

func (program *Program) Interpret(mainContent string) {
  lines := strings.Split(mainContent, "\n")

  for _,line := range lines {
    program.cleanupLine(line)
  }
}

func (program *Program) cleanupLine(line string) {
  tokens := strings.Split(line, " ")
  tokens = cleanupTokens(tokens)
  if len(tokens) == 0 {return}
  if strings.HasPrefix(tokens[0], "//") {return} 

  program.interpretLine(tokens)
}

func (program *Program) interpretLine(tokens []string) {
  actions := program.getActions()
  for key, value := range actions {
    if key.MatchString(tokens[0]) {
      value(tokens...)
      return
    }
  }
  println("Undefined keyword", tokens[0])
  os.Exit(1)
}
