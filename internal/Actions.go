package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Action = func(args ...string)

func (program *Program) getActions() map[*regexp.Regexp]Action {
  return map[*regexp.Regexp]Action {
    regexp.MustCompile("^var$"): program.variable,
    regexp.MustCompile("^output\\(.*$"): program.output,
    regexp.MustCompile("^input\\(.*$"): program.input,
  }
}

func (p *Program) assignment(args ...string) {
  joined := strings.Join(args, " ")
  if len(args) != 5 || args[1] != "=" {
    println("Invalid syntax:")
    println(joined)
    os.Exit(1)
  }
  val1, _,err := p.retrieveValue(args[2])
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
  int1, _ := strconv.Atoi(val1)

  val2, _,err := p.retrieveValue(args[4])
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
  int2, _ := strconv.Atoi(val2)
  result, err := mathOperations(args[3], int1, int2)
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }

  p.addVariable(args[0], "int", fmt.Sprintf("%d",result) )

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
  err = program.addVariable(str, "int", in)
  if err != nil {
    println(err.Error())
    os.Exit(1)
  }
}

func (program *Program) variable(args ...string) {
  if len(args) < 5 || args[3] != "=" {
    println("Invalid syntax:")
    command := strings.Join(args, " ")
    println(command)
    os.Exit(1)
  }
  err := program.addVariable(args[1], args[2], strings.Join(args[4:], " ") )
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
