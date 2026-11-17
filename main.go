package main

import (
	"fmt"
	"os"
	"push-swap/parser"
	"push-swap/sorter"
	"push-swap/stack"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	a, err := parser.ParseArguments(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	b := stack.NewStack()
	sorter.Sort(a, b)
}
