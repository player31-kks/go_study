package queue

import "testing"

func TestAdd(t *testing.T) {
	queue := NewQueue(10)
	queue.Add(10)

	want := 10
	if result := queue.data[0]; result != want {
		t.Errorf("expected '%d' but got '%d'", 10, want)
	}
}

func TestEmpty(t *testing.T) {
	t.Run("queue가 비어있을 때", func(t *testing.T) {
		queue := NewQueue(10)
		want := true

		if result := queue.Empty(); result != want {
			t.Errorf("expected '%#v' but got '%t'", true, result)
		}
	})

	t.Run("queue가 존재할 때", func(t *testing.T) {
		queue := NewQueue(10)
		queue.Add(10)
		want := false

		if result := queue.Empty(); result != want {
			t.Errorf("expected '%#v' but got '%t'", false, result)
		}
	})
}

func TestPoll(t *testing.T) {
	t.Run("queue가 비어있을 때", func(t *testing.T) {
		queue := NewQueue(10)

		if result := queue.Poll(); result != nil {
			t.Errorf("expected '%#v' but got '%t'", true, result)
		}
	})

	t.Run("queue가 존재할 때", func(t *testing.T) {
		queue := NewQueue(10)
		queue.Add(10)
		want := 10

		if result := queue.Poll(); result != want && queue.Empty() {
			t.Errorf("expected '%#v' but got '%t'", false, result)
		}
	})
}
