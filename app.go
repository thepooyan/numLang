package main

import (
	"os"
  "numLang/internal"
)

func main() {
  if len(os.Args) != 2 {
    println("Usage: numlang <filename>")
    os.Exit(1)
  }

  filename := os.Args[1]

  content, err := os.ReadFile(filename)
  if err != nil {
    println("Error reading file: %s", filename)
    os.Exit(1)
  }

  fileText := string(content)
  internal.Interpret(fileText)
}
