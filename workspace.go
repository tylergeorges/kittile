package main

import (
	"github.com/tylergeorges/kittile/layout"
)

type Workspace struct {
	Tree *layout.TreeNode
}

func NewWorkspace() *Workspace {
	return &Workspace{
		Tree: nil,
	}
}

func (ws *Workspace) ApplyLayout(n *layout.TreeNode) {
	if n == nil || (n.Layout == nil && !n.IsLeaf()) {
		return
	}

	
	if n.IsLeaf() {
		ws.ApplyLayout(n.FirstChild)
		ws.ApplyLayout(n.SecondChild)
	}
	
	PositionWindow(n.Id, n.Layout, true)

}

func (ws *Workspace) Update() {
	node := ws.Tree

	ws.ApplyLayout(node)
}
