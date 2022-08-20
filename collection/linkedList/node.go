package linkedlist

type Node struct {
	val  interface{}
	next *Node
}

func NewNode(val interface{}) *Node {
	node := new(Node)
	node.val = val
	return node
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

func (n *Node) InsertNext(node *Node) {
	if node != nil {
		node.next = n.next
		n.next = node
	}
}

func (n *Node) RemoveNext() *Node {
	nextNode := n.next
	if nextNode != nil {
		n.next = nextNode.next
	}
	return nextNode
}
