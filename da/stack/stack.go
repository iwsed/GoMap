package main

import (
	"errors"
	"fmt"
	"sync"
)

//实现stack push, pop, peek

type stk struct {
	lock sync.Mutex
	s    []int
}

func NewStk() *stk {
	return &stk{sync.Mutex{}, make([]int, 0)}
}

func (s *stk) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stk) Pop() (int, error) {
	s.lock.Lock()
	s.lock.Unlock()

	len := len(s.s)
	if len == 0 {
		return -1, errors.New("Empty Stack")
	}
	v := s.s[len-1]
	s.s = s.s[:len-1]
	return v, nil
}

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

func main() {
	fmt.Println("hello go")

	ss := NewStk()
	ss.Push(10)
	ss.Push(11)
	ss.Push(12)
	fmt.Println(ss.s)
	vv, _ := ss.Pop()
	vv, _ = ss.Pop()
	fmt.Println(ss.s, vv)

	return

	s := make(stack, 0)
	s = s.Push(10)
	s = s.Push(11)
	s = s.Push(12)
	fmt.Println(s)
	s, v := s.Pop()
	fmt.Println(s, v)
	s, v = s.Pop()
	fmt.Println(s, v)
	s, v = s.Pop()
	fmt.Println(s, v)
	s, v = s.Pop()
	fmt.Println(s, v)

	return

	stack := make([]int, 0)
	//v := 5
	fmt.Printf("aa %p, %p\n", stack, &stack)

	PushS(&stack, 5)
	PushS(&stack, 10)
	PushS(&stack, 11)
	//PushS(&stack, 12)

	//stack = append(stack, 11)

	fmt.Printf("aa %p, %p\n", stack, &stack)
	fmt.Println(stack)

	v = PopS(&stack)
	fmt.Println(v, stack)
	v = PopS(&stack)
	fmt.Println(v, stack)
	v = PopS(&stack)
	fmt.Println(v, stack)
	v = PopS(&stack)
	fmt.Println(v, stack)

}

func PushS(s *[]int, v int) {
	fmt.Printf("%p, %p\n", s, *s)
	*s = append(*s, v)
}

func PopS(s *[]int) int {
	if len(*s) == 0 {
		return 0
	}
	len := len(*s)
	fmt.Println("len of stack", len)
	v := (*s)[len-1]
	*s = (*s)[:len-1]

	return v
}
