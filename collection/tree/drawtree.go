package tree

import (
	"gonum.org/v1/plot/plotter"
)

type DrawTreeNode[T any] struct {
	Value T
	X     int
	Y     int
	Child []*DrawTreeNode[T]
}

func (d *DrawTreeNode[T]) getLocations(xys *plotter.XYs) {
	*xys = append(*xys, plotter.XY{
		float64(d.X),
		float64(d.Y)})

	for _, child := range d.Child {
		child.getLocations(xys)
	}
}

func makeDrawTree[T any](node *TreeNode[T], level int, order *int) *DrawTreeNode[T] {
	if node == nil {
		return nil
	}
	drawTreeNode := &DrawTreeNode[T]{
		Value: node.Value,
		Y:     level - 1,
	}
	// in order
	half := len(node.Child) / 2
	for i := 0; i < half; i++ {
		drawTreeNode.Child = append(drawTreeNode.Child, makeDrawTree(node.Child[i], level-1, order))
	}

	drawTreeNode.X = *order
	*order++

	for i := half; i < len(node.Child); i++ {
		drawTreeNode.Child = append(drawTreeNode.Child, makeDrawTree(node.Child[i], level-1, order))
	}
	return drawTreeNode
}
