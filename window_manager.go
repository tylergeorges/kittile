package main

import (
	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/layout"
)

type WindowManager struct {
	Monitors *Monitor
}

func (w *WindowManager) Run() {
	m := w.Monitors

	ws := m.workspace

	tree := ws.GetTree()
	//
	tree.FlipTree(layout.FlipVertical)
	tree.Render()
}

func NewWindowManager() *WindowManager {
	w := w32.GetForegroundWindow()

	m := NewMonitor(w)

	return &WindowManager{
		Monitors: m,
	}
}
