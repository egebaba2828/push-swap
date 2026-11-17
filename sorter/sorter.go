package sorter

import (
	"push-swap/operations"
	"push-swap/stack"
)

func Sort(a, b *stack.Stack) {
	if a.IsSorted() {
		return
	}
	if len(a.Items) == 2 && a.Items[0] > a.Items[1] {
		operations.Sa(a)
		return
	}
	for len(a.Items) > 3 {
		operations.Pb(a, b)
	}
	sortThree(a)
	for !b.IsEmpty() {
		operations.Pa(a, b)
	}
}

func sortThree(a *stack.Stack) {
	if a.Items[0] > a.Items[1] && a.Items[1] < a.Items[2] && a.Items[0] < a.Items[2] {
		operations.Sa(a)
	} else if a.Items[0] > a.Items[1] && a.Items[1] > a.Items[2] {
		operations.Sa(a)
		operations.Rra(a)
	} else if a.Items[0] > a.Items[1] && a.Items[1] < a.Items[2] && a.Items[0] > a.Items[2] {
		operations.Ra(a)
	}
}
