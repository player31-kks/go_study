package mymap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Element[Key constraints.Ordered, TValue any] struct {
	Key   Key
	Value TValue
}

type SortedMap[Key constraints.Ordered, TValue any] struct {
	Arr []Element[Key, TValue]
}

func (m *SortedMap[Key, TValue]) Add(key Key, value TValue) {
	idx := sort.Search(len(m.Arr), func(i int) bool {
		return m.Arr[i].Key >= key
	})

	if idx < len(m.Arr) && m.Arr[idx].Key == key {
		m.Arr[idx].Value = value
		return
	}

	before := m.Arr[:idx]
	after := m.Arr[idx:]
	m.Arr = append(before, Element[Key, TValue]{key, value})
	m.Arr = append(m.Arr, after...)
}

func (m *SortedMap[Key, TValue]) Get(key Key) (TValue, bool) {
	idx := sort.Search(len(m.Arr), func(i int) bool {
		return m.Arr[i].Key >= key
	})
	if idx < len(m.Arr) && m.Arr[idx].Key == key {
		return m.Arr[idx].Value, true
	}
	var zeroValue TValue
	return zeroValue, false
}

func (m *SortedMap[Key, TValue]) Delete(key Key) bool {
	idx := sort.Search(len(m.Arr), func(i int) bool {
		return m.Arr[i].Key >= key
	})
	if idx < len(m.Arr) && m.Arr[idx].Key == key {
		m.Arr = append(m.Arr[:idx], m.Arr[idx+1:]...)
		return true
	}
	return false
}
