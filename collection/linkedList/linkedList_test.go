package linkedlist

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	list := new(LinkedList)
	list.Append(10)
	list.Append(20)
	list.Insert(0, 30)
	list.Insert(3, 50)
	list.Append(60)
	list.RemoveAt(2)
	list.Remove(60)

	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", list.GetAt(i))
	}

	want := 30
	if result := list.GetAt(0); result != want {
		t.Errorf("expected '%d' but got '%d'", want, result)
	}
}
