package main

import (
	"bufio"
	"fmt"
	"os"
	"push-swap/operations"
	"push-swap/parser"
	"push-swap/stack"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch line {
		case "sa":
			operations.Swap(a)
		case "sb":
			operations.Swap(b)
		case "ss":
			operations.Swap(a)
			operations.Swap(b)
		case "pa":
			operations.Push(b, a)
		case "pb":
			operations.Push(a, b)
		case "ra":
			operations.Rotate(a)
		case "rb":
			operations.Rotate(b)
		case "rr":
			operations.Rotate(a)
			operations.Rotate(b)
		case "rra":
			operations.ReverseRotate(a)
		case "rrb":
			operations.ReverseRotate(b)
		case "rrr":
			operations.ReverseRotate(a)
			operations.ReverseRotate(b)
		}
	}

	if a.IsSorted() && b.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
