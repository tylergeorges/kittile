package layout

import (
	"fmt"

	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/rect"
	"github.com/tylergeorges/kittile/windows_api"
)

type TreeNode struct {
	FirstChild  *TreeNode
	SecondChild *TreeNode
	parent      *TreeNode
	Layout      *rect.Rect
	Id          w32.HWND
	direction   direction
}

func NewLeaf(first_child, second_child, parent *TreeNode, d direction) *TreeNode {
	return &TreeNode{FirstChild: first_child, SecondChild: second_child, direction: d, Layout: parent.Layout}

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
	return &TreeNode{FirstChild: nil, SecondChild: nil, Id: id, parent: nil, Layout: nil}
}

func (t *TreeNode) Display() {

	if t == nil || (t.Layout.IsEmpty() && !t.IsLeaf()) {
		return
	}
	t.FirstChild.Display()
	t.SecondChild.Display()

	windows_api.PositionWindow(t.Id, t.Layout, true)
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

func (t *TreeNode) IsLeaf() bool {
	return t.FirstChild != nil && t.SecondChild != nil
}

func (t *TreeNode) NodeName() string {
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

func (t *TreeNode) ApplyLayouts() {

	if t == nil {
		return
	}

	dir := t.direction

	switch dir {
	case Horizontal: // windows should stack on top of eachother
		{

			// first_child = *t.SecondChild
			x, y, width, height := t.Layout.Pieces()
			half_height := height / 2

			top_rect := rect.New(x, y, width, half_height)
			bottom_rect := rect.New(x, y+half_height, width, half_height)

			t.FirstChild.SetGeom(top_rect.Pieces())
			t.SecondChild.SetGeom(bottom_rect.Pieces())

		}

	case Vertical:
		{

			x, y, width, height := t.Layout.Pieces()

			half_width := width / 2

			left_rect := rect.New(x, y, half_width, height)
			right_rect := rect.New(x+half_width, y, half_width, height)

			t.FirstChild.SetGeom(left_rect.Pieces())
			t.SecondChild.SetGeom(right_rect.Pieces())

		}

	}

	t.FirstChild.ApplyLayouts()
	t.SecondChild.ApplyLayouts()

}

func (t *TreeNode) get_next_direction() direction {
	if t.direction == Vertical {
		return Horizontal
	}

	return Vertical
}

func (t *TreeNode) Insert(second_child *TreeNode) {
	exists := t.NodeExists(second_child.Id)

	if exists {
		return
	}

	if t.FirstChild == nil {
		second_child.SetGeom(t.Layout.Pieces())

		t.FirstChild = second_child
		return
	}

	if t.SecondChild == nil {
		t.SecondChild = second_child

		return
	}

	if t.SecondChild.IsLeaf() {
		t.SecondChild.Insert(second_child)
		return
	}

	first_child := t.SecondChild

	leaf := NewLeaf(first_child, second_child, t, t.get_next_direction())

	t.SecondChild = leaf

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

func (t *TreeNode) IsEmpty() bool {
	if t == nil {
		return true
	}

	if t.IsLeaf() {
		return t.FirstChild == nil && t.SecondChild == nil
	}

	if t.FirstChild != nil || t.SecondChild != nil {
		return false
	}

	return true
}

func (t *TreeNode) FlipTree(flp flip_direction) {
	if t == nil {
		return
	}

	var tmp *TreeNode

	if (flp == FlipVertical && t.direction == Vertical) || (flp == FlipHorizontal && t.direction == Horizontal) {
		tmp = t.FirstChild

		t.FirstChild = t.SecondChild
		t.SecondChild = tmp

	}

	t.FirstChild.FlipTree(flp)
	t.SecondChild.FlipTree(flp)
}

func (t TreeNode) Render() {
	t.ApplyLayouts()
	t.Display()
}

func NewTree(layout *rect.Rect, inverted bool) *TreeNode {
	return &TreeNode{
		Layout:      layout,
		FirstChild:  nil,
		SecondChild: nil,
		parent:      nil,
		direction:   Vertical,
		Id:          0,
	}
}
