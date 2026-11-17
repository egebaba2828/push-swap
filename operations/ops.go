package operations

import (
	"fmt"
	"push-swap/stack"
)

func Swap(s *stack.Stack) {
	if len(s.Items) < 2 {
		return
	}
	s.Items[0], s.Items[1] = s.Items[1], s.Items[0]
}

func Sa(a *stack.Stack) {
	Swap(a)
	fmt.Println("sa")
}

func Push(src, dest *stack.Stack) {
	val, ok := src.Pop()
	if !ok {
		return
	}
	dest.Push(val)
}

func Pa(a, b *stack.Stack) {
	Push(b, a)
	fmt.Println("pa")
}

func Pb(a, b *stack.Stack) {
	Push(a, b)
	fmt.Println("pb")
}

func Rotate(s *stack.Stack) {
	if len(s.Items) < 2 {
		return
	}
	val := s.Items[0]
	s.Items = append(s.Items[1:], val)
}

func Ra(a *stack.Stack) {
	Rotate(a)
	fmt.Println("ra")
}

func ReverseRotate(s *stack.Stack) {
	if len(s.Items) < 2 {
		return
	}
	n := len(s.Items)
	val := s.Items[n-1]
	s.Items = append([]int{val}, s.Items[:n-1]...)
}

func Rra(a *stack.Stack) {
	ReverseRotate(a)
	fmt.Println("rra")
}
