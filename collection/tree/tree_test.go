package tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeAdd(t *testing.T) {
	root := new(TreeNode[string])
	root.Value = "root"

	child1 := root.Add("child1")
	child2 := root.Add("child2")
	child3 := root.Add("child3")

	child1.Add("child1-1")
	child1.Add("child1-2")

	child2.Add("child2-1")
	child2.Add("child2-2")

	child3.Add("child3-1")

	var line strings.Builder
	root.PreOrder(func(value string) {
		line.WriteString(fmt.Sprint(value, "->"))
	})
	t.Log(line.String())
	line.Reset()

	root.PostOrder(func(value string) {
		line.WriteString(fmt.Sprint(value, "->"))
	})
	t.Log(line.String())
	line.Reset()

	root.LevelOrder(func(value string) {
		line.WriteString(fmt.Sprint(value, "->"))
	})
	t.Log(line.String())
	line.Reset()

}
