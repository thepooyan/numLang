package internal

type Action = func(args ...string)

func (program *Program) getActions() map[string]Action {
  return map[string]Action {
    "var": program.variable,
  }
}

func (program *Program) variable(args ...string) {
  println("variable")
  program.addVariable("name", "int", "hi")
  for _,t := range args {
    println(t)
  }
}
