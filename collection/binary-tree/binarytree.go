package binarytree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value any
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) GetChildren() []*TreeNode {
	var children []*TreeNode
	if t.Left != nil {
		children = append(children, t.Left)
	}
	if t.Right != nil {
		children = append(children, t.Right)
	}
	return children
}

func (t *TreeNode) GetValue() any {
	return t.Value
}
