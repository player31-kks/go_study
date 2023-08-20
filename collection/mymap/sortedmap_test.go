package mymap

import "testing"

func TestSortedMap(t *testing.T) {
	m := SortedMap[int, string]{}
	m.Add(1, "one")
	m.Add(2, "two")
	m.Add(3, "three")
	m.Add(4, "four")
	m.Add(5, "five")

	if v, ok := m.Get(3); !ok || v != "three" {
		t.Errorf("expected 3 to be three, got %v", v)
	}
}
