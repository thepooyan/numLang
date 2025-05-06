package internal

import (
	"os"
	"strings"
)

type Action = func(args ...string)

func (program *Program) getActions() map[string]Action {
  return map[string]Action {
    "var": program.variable,
  }
}

func (program *Program) variable(args ...string) {
  if len(args) != 5 || args[3] != "=" {
    println("Invalid syntax:")
    command := strings.Join(args, " ")
    println(command)
    os.Exit(1)
  }
  err := program.addVariable(args[1], args[2], args[4])
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
}
