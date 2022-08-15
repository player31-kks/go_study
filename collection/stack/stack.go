package stack

import "fmt"

type Stack struct {
	data []interface{}
}

func NewStack(size int) *Stack {
	stack := new(Stack)
	stack.data = make([]interface{}, 0, size)
	return stack
}

func (s *Stack) Empty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func (s *Stack) Add(value interface{}) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		fmt.Println("stack is empty")
		return nil
	}
	lastValue := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return lastValue
}

func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[len(s.data)-1]
}
