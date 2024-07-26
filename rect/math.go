package rect

func (a Rect) DivNum(num int) Rect {

	x1, y1, w1, h1 := a.Pieces()

	x, y := x1/num, y1/num
	w, h := w1/num, h1/num

	return New(x, y, w, h)
}

func (a Rect) Div(b Rect) Rect {

	x1, y1, w1, h1 := a.Pieces()
	x2, y2, w2, h2 := b.Pieces()

	x, y := x1/x2, y1/y2
	w, h := w1/w2, h1/h2

	return New(x, y, w, h)
}

func (a Rect) Mul(b Rect) Rect {
	x1, y1, w1, h1 := a.Pieces()
	x2, y2, w2, h2 := b.Pieces()

	x, y := x1*x2, y1*y2
	w, h := w1*w2, h1*h2

	return New(x, y, w, h)
}

func (a Rect) Add(b Rect) Rect {
	x1, y1, w1, h1 := a.Pieces()
	x2, y2, w2, h2 := b.Pieces()

	x, y := x1+x2, y1+y2
	w, h := w1+w2, h1+h2

	return New(x, y, w, h)
}

func (a Rect) Sub(b Rect) Rect {
	x1, y1, w1, h1 := a.Pieces()
	x2, y2, w2, h2 := b.Pieces()

	x, y := x1-x2, y1-y2
	w, h := w1-w2, h1-h2

	return New(x, y, w, h)
}
