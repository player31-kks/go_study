package tree

type TreeNode[T any] struct {
	Value T
	Child []*TreeNode[T]
}

func (t *TreeNode[T]) Add(value T) *TreeNode[T] {
	treeNode := new(TreeNode[T])
	treeNode.Value = value
	t.Child = append(t.Child, treeNode)
	return treeNode
}

// 자기 자신을 먼저 호출
func (t *TreeNode[T]) PreOrder(fn func(value T)) {
	if t == nil {
		return
	}
	fn(t.Value)
	for _, child := range t.Child {
		child.PreOrder(fn)
	}
}

// 자기 자신을 나중에 호출
func (t *TreeNode[T]) PostOrder(fn func(value T)) {
	if t == nil {
		return
	}
	for _, child := range t.Child {
		child.PostOrder(fn)
	}
	fn(t.Value)
}

//BFS
func (t *TreeNode[T]) LevelOrder(fn func(value T)) {
	if t == nil {
		return
	}
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		fn(front.Value)
		queue = append(queue, front.Child...)
	}
}
