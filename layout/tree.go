package layout

import (
	"fmt"

	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/rect"
)

type TreeNode struct {
	FirstChild  *TreeNode
	SecondChild *TreeNode
	parent      *TreeNode
	Layout      *rect.Rect
	Id          w32.HWND
	direction   Direction
}

func NewLeaf(first_child, second_child, parent *TreeNode, direction Direction) *TreeNode {
	return &TreeNode{FirstChild: first_child, SecondChild: second_child, direction: direction, Layout: parent.Layout}

}

func (t *TreeNode) SetGeom(x, y, w, h int) {
	if t.Layout == nil {
		new_rect := rect.New(x, y, w, h)

		t.Layout = &new_rect

		return
	}

	t.Layout.SetGeom(x, y, w, h)
}

func NewNode(id w32.HWND) *TreeNode {
	return &TreeNode{FirstChild: nil, SecondChild: nil, Id: id, parent: nil}
}

func (t TreeNode) GetDirection() *string {
	var out string
	out = "<nil>"

	if t.direction == Vertical {
		dir := "Vertical"

		out = dir
	}

	if t.direction == Horizontal {
		dir := "Horizontal"

		out = dir
	}

	return &out
}

func (t TreeNode) IsLeaf() bool {
	return t.Id == 0
}

func (t TreeNode) NodeName() string {
	if t.IsLeaf() {
		return "Leaf"
	}

	return w32.GetWindowText(t.Id)
}

func (t TreeNode) NodeExists(id w32.HWND) bool {
	node := t.FindById(id)

	return node != nil
}

func (t TreeNode) String() string {
	var first_child string = "Not Set"
	var second_child string = "Not Set"
	var root string = t.NodeName()

	if t.IsLeaf() {
		root = "Leaf"
	}

	if t.FirstChild != nil {
		first_child = t.FirstChild.NodeName()
	}

	if t.SecondChild != nil {
		second_child = t.SecondChild.NodeName()
	}

	return fmt.Sprintf("root: %s\nFirst Child:%s \nSecond Child:%s \nDirection:%s  \n----------------------\n", root, first_child, second_child, *t.GetDirection())
}

func (t *TreeNode) UpdateLayout(second_child *TreeNode) {
	var dir Direction

	dir = t.direction

	switch dir {
	case Horizontal: // windows should stack on top of eachother
		{
			var first_child TreeNode

			first_child = *t.FirstChild

			if t.SecondChild != nil {
				first_child = *t.SecondChild
			}

			x, y, width, height := first_child.Layout.Pieces()

			half_height := height / 2

			second_child.SetGeom(x, y+half_height, width, half_height)
			first_child.SetGeom(x, y, width, half_height)

			dir = Vertical
		}

	default:
		{
			var first_child TreeNode
			first_child = *t.FirstChild

			if t.SecondChild != nil {
				first_child = *t.SecondChild
			}

			x, y, width, height := first_child.Layout.Pieces()

			half_width := width / 2

			second_child.SetGeom(x+half_width, y, half_width, height)
			first_child.SetGeom(x, y, half_width, height)

			dir = Horizontal
		}

	}

	if t.SecondChild == nil {
		t.SecondChild = second_child
	}

	leaf := NewLeaf(t.SecondChild, second_child, t, dir)
	leaf.direction = dir

	t.SecondChild = leaf

}

func (t *TreeNode) Insert(second_child *TreeNode) {
	exists := t.NodeExists(second_child.Id)

	if exists {
		return
	}

	if t.FirstChild == nil {
		layout_copy := *t.Layout

		second_child.SetGeom(layout_copy.Pieces())

		t.FirstChild = second_child

		return
	}

	if t.SecondChild != nil && t.SecondChild.IsLeaf() {
		t.SecondChild.Insert(second_child)

		return
	}

	t.UpdateLayout(second_child)
}

func (t *TreeNode) FindById(id w32.HWND) *TreeNode {
	if t.Id == id {
		return t
	}

	if t.Id == 0 {
		return nil
	}

	if t.FirstChild != nil {
		node := t.FirstChild.FindById(id)

		if node != nil {
			return node
		}
	}

	if t.SecondChild != nil {
		node := t.SecondChild.FindById(id)

		if node != nil {
			return node
		}
	}

	return nil
}

func NewTree(layout *rect.Rect) *TreeNode {
	return &TreeNode{
		Layout:      layout,
		FirstChild:  nil,
		SecondChild: nil,
		parent:      nil,
		direction:   Vertical,
		Id:          0,
	}
}
