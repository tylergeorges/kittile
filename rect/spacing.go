package rect

func (r *Rect) Padding(padding int) {
	r.X -= padding
	r.Y -= padding

	r.Width += padding * 2
	r.Height += padding * 2
}

func (r *Rect) PaddingLeft(padding int) {
	r.X += padding
}

func (r *Rect) PaddingRight(padding int) {

	r.Width -= padding
}

func (r *Rect) Margin(margin int) {

	r.Y += margin
	r.X += margin

	r.Width -= margin * 2
	r.Height -= margin * 2
}
