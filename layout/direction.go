package layout

type direction int

const (
	Vertical direction = iota + 1
	Horizontal
)

type flip_direction int

const (
	FlipVertical flip_direction = iota + 1
	FlipHorizontal
)
