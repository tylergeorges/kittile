package main

import "github.com/tylergeorges/kitty/layout"

type Container struct {
	windows []Window
	Layout  *layout.Layout
}

func (c *Container) AddWindow(w *Window) {
	c.windows = append(c.windows, *w)

}

func (c *Container) Len() int {
	return len(c.windows)

}

func (c *Container) IsFull() bool {
	return c.Len() >= cap(c.windows)
}

// base_window_rect := *w32.GetWindowRect(window)
// raw_window_rect := rect.FromRECT(base_window_rect)

// raw_window_rect.X -= work_area_rect.X
// raw_window_rect.Y -= work_area_rect.Y
