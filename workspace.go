package main

import (
	"github.com/tylergeorges/kittile/layout"
)

type Workspace struct {
	tree *layout.TreeNode
}

func NewWorkspace() *Workspace {
	return &Workspace{
		tree: nil,
	}
}

func (ws *Workspace) SetTree(t *layout.TreeNode) {
	ws.tree = t
}

func (ws *Workspace) GetTree() *layout.TreeNode {
	return ws.tree

}

func (ws *Workspace) ApplyLayout(n *layout.TreeNode) {

	if n == nil || (n.Layout.IsEmpty() && !n.IsLeaf()) {
		return
	}

	if n.IsLeaf() {
		ws.ApplyLayout(n.FirstChild)
		ws.ApplyLayout(n.SecondChild)
	}

}

// func (ws *Workspace) Update() {
// 	node := ws.Tree

// 	node.FlipTree(layout.FlipVertical)
// 	node.Render()
// }
