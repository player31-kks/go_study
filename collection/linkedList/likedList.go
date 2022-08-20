package linkedlist

import "fmt"

type List interface {
	Append(val interface{})
	GetAt(index int)
	Insert(index int, val interface{})
	Contains(val interface{})
	Remove(val interface{})
	RemoveAt(index int)
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Append(val interface{}) {
	if l.isEmpty() {
		newNode := NewNode(val)
		l.head = newNode
		l.tail = newNode
		return
	}
	nextNode := NewNode(val)
	l.tail.SetNext(nextNode)
	l.tail = nextNode
}

func (l *LinkedList) Insert(index int, val interface{}) {
	prevNode := l.getNodeByIndex(index - 1)

	// 아무것도 없을 때
	if prevNode == nil {
		fmt.Printf("not insert")
		return
	}

	// 헤드 노드를 변경할 때
	if prevNode == l.head {
		newNode := NewNode(val)
		newNode.SetNext(prevNode)
		l.head = newNode
		return
	}

	// 일반 노드를 변경할 때
	newNode := NewNode(val)
	prevNode.InsertNext(newNode)
	if prevNode == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList) RemoveAt(index int) {
	// 아무것도 없을 때
	prevNode := l.getNodeByIndex(index - 1)
	if prevNode == nil {
		fmt.Printf("not remove")
		return
	}
	// 헤드일 때
	if prevNode == l.head {
		l.head = nil
		return
	}
	prevNode.RemoveNext()
}

func (l *LinkedList) Remove(val interface{}) {
	node, idx := l.getNodeByValue(val)
	if node == nil {
		fmt.Printf("not remove")
		return
	}
	if node == l.head {
		l.head = nil
		return
	}
	l.RemoveAt(idx)
}

func (l *LinkedList) GetAt(index int) interface{} {
	node := l.getNodeByIndex(index)
	return node.val
}

func (l *LinkedList) Contains(val interface{}) bool {
	node, _ := l.getNodeByValue(val)
	return node != nil
}

func (l *LinkedList) getNodeByIndex(index int) *Node {
	travelNode := l.head
	for i := 0; i < index; i++ {
		if travelNode == nil {
			return nil
		}
		travelNode = travelNode.next
	}
	return travelNode
}

func (l *LinkedList) getNodeByValue(value interface{}) (*Node, int) {
	travelNode := l.head
	idx := 0
	for travelNode != nil {
		if travelNode.val == value {
			return travelNode, idx
		}
		travelNode = travelNode.next
		idx += 1
	}
	return travelNode, 0
}
