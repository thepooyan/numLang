package internal

type Action = func(args ...string)

func getActions() map[string]Action {
  return map[string]Action {
    "var": variable,
  }
}

func variable(args ...string) {
  println("variable")
  println(args)
}
