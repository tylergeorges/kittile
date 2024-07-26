package main

import (
	w32 "github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/layout"
	"github.com/tylergeorges/kittile/rect"
)

type Monitor struct {
	size rect.Rect // size of monitor

	work_area rect.Rect

	workspace Workspace

	id w32.HMONITOR
}

func (m *Monitor) LoadFocusedWorkspace() {
	tree := layout.NewTree(&m.work_area)

	w32.EnumWindows(func(w w32.HWND) bool {
		is_minimized := IsIconic(w)

		window := NewWindow(w)

		if window.IsVisible() && window.IsWindow() && !is_minimized && window.HasTitle() && window.Exe() {
			node := layout.NewNode(w)

			tree.Insert(node)
		}

		return true

	})

	m.workspace.Tree = tree

}

func (m Monitor) UpdateWorkspace() {
	m.workspace.Update()
}

func NewMonitor(w w32.HWND) *Monitor {
	monitor_id := w32.MonitorFromWindow(w, w32.MONITOR_DEFAULTTOPRIMARY)

	monitor_info := w32.MONITORINFO{}

	w32.GetMonitorInfo(monitor_id, &monitor_info)

	work_area := rect.FromRECT(monitor_info.RcWork)

	monitor := &Monitor{
		work_area: work_area,
		size:      rect.FromRECT(monitor_info.RcMonitor),
		id:        monitor_id,
		workspace: *NewWorkspace(),
	}

	monitor.LoadFocusedWorkspace()

	return monitor
}
