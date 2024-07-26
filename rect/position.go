package rect

// Completely position a child to the ceneter of the container.
func (r Rect) Center(child_rect Rect) {

	half_width := r.Width / 2
	half_height := r.Height / 2

	x := r.X + ((r.Width - half_width) / 2)
	y := r.Y + ((r.Height - half_height) / 2)

	child_rect.SetGeom(x, y, half_width, half_height)
}

// Vertically position a child to the end of container.
func (r Rect) AlignEnd(child_rect Rect) {
	half_width := r.Width / 2

	half_height := r.Height / 2

	child_rect.SetGeom(0, half_height, half_width, half_height)
}

// Vertically position a child to the start of container.
func (r Rect) AlignStart(child_rect Rect) {
	half_width := r.Width / 2
	half_height := r.Height / 2

	child_rect.SetGeom(0, 0, half_width, half_height)
}

// Horizontally position a child to the end of container.
func (r Rect) JustifyEnd(child_rect Rect) {
	half_width := r.Width / 2

	child_rect.SetGeom(half_width, 0, half_width, r.Height)
}
