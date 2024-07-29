package rect

import (
	"fmt"

	"github.com/gonutz/w32/v2"
)

type Rect struct {
	X      int // The Left point in a [w32.RECT]
	Y      int // The Top point in a [w32.RECT]
	Width  int // The Right point in a [w32.RECT]
	Height int // The Bottom point in a [w32.RECT]
}

func (r Rect) String() string {
	return fmt.Sprintf("[(%d, %d) %dx%d]", r.X, r.Y, r.Width, r.Height)
}

func (r Rect) Pieces() (x, y, w, h int) {
	return r.X, r.Y, r.Width, r.Height
}

func (r Rect) Copy() Rect {
	return New(r.Pieces())
}

func (r *Rect) SetGeom(x, y, w, h int) {
	
	r.X = x
	r.Y = y

	r.Width = w
	r.Height = h
}

func (r *Rect) IsEmpty() bool {
	if r == nil {
		return true
	}

	return r.X == 0 && r.Y == 0 && r.Height == 0 && r.Width == 0
}

// Convert a [w32.RECT] to a [Rect]
func FromRECT(rect w32.RECT) Rect {
	x, y := int(rect.Left), int(rect.Top)
	w, h := int(rect.Right), int(rect.Bottom)

	return New(x, y, w, h)
}

func New(x, y, w, h int) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}
