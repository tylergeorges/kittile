package main

import "github.com/gonutz/w32/v2"

type WindowManager struct {
	Monitors *Monitor
}

func (wm *WindowManager) NextWindow() {

}

func NewWindowManager() *WindowManager {
	w := w32.GetForegroundWindow()

	m := NewMonitor(w)

	return &WindowManager{
		Monitors: m,
	}
}
