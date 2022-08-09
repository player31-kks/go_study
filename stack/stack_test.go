package stack

import (
	"testing"
)

func TestAdd(t *testing.T) {
	stack := NewStack(10)
	stack.Add(10)

	data := stack.data[0]
	want := 10
	if data != want {
		t.Errorf("expected '%d' but got '%d'", 10, want)
	}
}

func TestPop(t *testing.T) {

	t.Run("stack에 가장 마지막값을 가지고 옴", func(t *testing.T) {
		stack := NewStack(10)
		stack.Add(100)

		want := 100
		if result := stack.Pop(); result != want {
			t.Errorf("expected '%d' but got '%d'", want, result)
		}
	})

	t.Run("stack에 값이 비여있을 때 nil 반환", func(t *testing.T) {
		stack := NewStack(10)
		if result := stack.Pop(); result != nil {
			t.Errorf("expected '%#v' but got '%d'", nil, result)
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Run("스택에 값이 없을 때", func(t *testing.T) {
		stack := NewStack(10)
		if result := stack.Empty(); result != true {
			t.Errorf("expected '%#v' but got '%t'", true, result)
		}
	})

	t.Run("스택에 값이 있을 때", func(t *testing.T) {
		stack := NewStack(10)
		stack.Add(10)
		if result := stack.Empty(); result != false {
			t.Errorf("expected '%#v' but got '%t'", false, result)
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("스택에 값이 없을 때", func(t *testing.T) {
		stack := NewStack(10)
		if result := stack.Peek(); result != nil {
			t.Errorf("expected '%#v' but got '%t'", nil, result)
		}
	})

	t.Run("스택에 값이 있을 때", func(t *testing.T) {
		stack := NewStack(10)
		stack.Add(100)
		if result := stack.Peek(); result != 100 {
			t.Errorf("expected '%#v' but got '%t'", 100, result)
		}
	})
}
