package internal

import (
	"os"
	"strings"
  "regexp"
)

type Action = func(args ...string)

func (program *Program) getActions() map[*regexp.Regexp]Action {
  return map[*regexp.Regexp]Action {
    regexp.MustCompile("^var$"): program.variable,
    regexp.MustCompile("^output\\(.*$"): program.output,
    regexp.MustCompile("^input\\(.*$"): program.input,
  }
}

func (program *Program) input(args ...string) {
  join := strings.Join(args, " ")
  str,ok := ExtractParenthesisContent(join)
  if ok != nil {
    println("Invalid syntax:")
    println(join)
    os.Exit(1)
  } 
  _, _, err := program.retrieveValue(str)
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
  in := Input()
  err = program.addVariable(str, "decimal", in)
  if err != nil {
    println(err.Error())
    os.Exit(1)
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

func (program *Program) output(args ...string) {
  joined := strings.Join(args, " ")
  str, err := ExtractParenthesisContent(joined)
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
  if str, ok := ExtractQuotedString(str); ok {
    println(str)
    return
  }
  val, _, err := program.retrieveValue(str)
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
  println(val)
}
