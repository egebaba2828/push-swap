package parser

import (
	"errors"
	"push-swap/stack"
	"strconv"
	"strings"
)

func ParseArguments(args []string) (*stack.Stack, error) {
	if len(args) == 0 {
		return nil, errors.New("no args")
	}

	var rawNumbers []string
	for _, arg := range args {
		fields := strings.Fields(arg)
		rawNumbers = append(rawNumbers, fields...)
	}

	a := stack.NewStack()
	seen := make(map[int]bool)

	for i := len(rawNumbers) - 1; i >= 0; i-- {
		val, err := strconv.Atoi(rawNumbers[i])
		if err != nil || seen[val] {
			return nil, errors.New("Error")
		}
		seen[val] = true
		a.Push(val)
	}

	return a, nil
}
