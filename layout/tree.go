package layout

import (
	"fmt"

	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kitty/rect"
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

func (t TreeNode) get_next_direction() Direction {
	if t.direction == Vertical {
		return Horizontal
	}

	return Vertical

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

	is_root := t.SecondChild != nil // if true, second child should be replaced with leaf node

	if is_root {

		if t.SecondChild.IsLeaf() {
			t.SecondChild.Insert(second_child)

			return
		}

		var dir = t.get_next_direction()

		switch dir {
		case Horizontal: // windows should stack on top of eachother
			{
				var first_child = *t.FirstChild

				if t.SecondChild != nil {
					first_child = *t.SecondChild
				}

				second_child.SetGeom(first_child.Layout.X, first_child.Layout.Y+first_child.Layout.Height/2, first_child.Layout.Width, first_child.Layout.Height/2)

				first_child.SetGeom(first_child.Layout.X, 0, first_child.Layout.Width, first_child.Layout.Height/2)

				dir = Vertical

			}

		case Vertical:
			{
				var first_child = *t.FirstChild

				if t.SecondChild != nil {
					first_child = *t.SecondChild
				}

				second_child.SetGeom(first_child.Layout.X+first_child.Layout.Width/2, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

				first_child.SetGeom(first_child.Layout.X, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

				dir = Horizontal
			}

		}

		leaf := NewLeaf(t.SecondChild, second_child, t, dir)

		leaf.direction = dir

		t.SecondChild = leaf

		return
	}

	var first_child = *t.FirstChild

	if t.SecondChild != nil {
		first_child = *t.SecondChild
	}

	second_child.SetGeom(first_child.Layout.X+first_child.Layout.Width/2, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	first_child.SetGeom(first_child.Layout.X, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	t.SecondChild = second_child

	// if t.parent == nil { // if root node
	// 	var first_child = *t.FirstChild

	// 	if t.SecondChild != nil {
	// 		first_child = *t.SecondChild
	// 	}

	// 	second_child.SetGeom(first_child.Layout.X+first_child.Layout.Width/2, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	// 	first_child.SetGeom(first_child.Layout.X, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	// 	t.SecondChild = second_child
	// 	return
	// }

	// switch dir {
	// case Horizontal: // windows should stack on top of eachother
	// 	{

	// 		// if t.parent != nil {
	// 		// 	left_layout := t.Layout

	// 		// 	base_copy := rect.New(left_layout.X, left_layout.Y, left_layout.Width, left_layout.Height/2)

	// 		// 	copy = base_copy

	// 		// }

	// 		// else {
	// 		// 	left_layout := t.Layout.Copy()

	// 		// 	base_copy := rect.New(left_layout.X, left_layout.Y, left_layout.Width, 500)

	// 		// 	copy = base_copy
	// 		// 	// copy = t.Layout.Copy()
	// 		// }

	// 		// first_child := *t.Left
	// 		// first_child.Layout.Height = half_height

	// 		// copy.Height /= 2

	// 		// node.Layout.Height = half_height
	// 		// node.Layout.Y += half_height

	// 		// half_height := copy.Height / 2

	// 		second_child.SetGeom(t.Layout.X, t.Layout.Y+t.Layout.Height/2, t.Layout.Width, t.Layout.Height/2)

	// 		t.SetGeom(t.Layout.X, t.Layout.Height/2, t.Layout.Width, t.Layout.Height/2)

	// 	}

	// case Vertical:
	// 	{
	// 		var first_child = *t.FirstChild

	// 		if t.SecondChild != nil {
	// 			first_child = *t.SecondChild
	// 		}

	// 		second_child.SetGeom(first_child.Layout.X+first_child.Layout.Width/2, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	// 		first_child.SetGeom(first_child.Layout.X, first_child.Layout.Y, first_child.Layout.Width/2, first_child.Layout.Height)

	// 	}

	// }

	// var new_dir Direction

	// if dir == Horizontal {
	// 	new_dir = Vertical
	// } else {
	// 	new_dir = Horizontal
	// }

	// leaf := NewLeaf(t.SecondChild, second_child, t, new_dir)

	// leaf.direction = new_dir

	// t.SecondChild = leaf

	// return

	// }

	// t.SecondChild.Insert(second_child)

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
