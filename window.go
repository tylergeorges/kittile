package main

import (
	"strings"

	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/rect"
	"github.com/tylergeorges/kittile/windows_api"
)

type Window struct {
	hwnd w32.HWND
}

func (w *Window) SetPosition(layout rect.Rect, top bool) {
	windows_api.PositionWindow(w.hwnd, &layout, top)
}

func GetWindowThreadProcessId(hwnd w32.HWND) (uint32, uint32) {

	thread_id, process_id := w32.GetWindowThreadProcessId(hwnd)

	return uint32(thread_id), uint32(process_id)
}

func (w *Window) Exe() bool {
	_, process_id := GetWindowThreadProcessId(w.hwnd)

	handle := windows_api.ProcessHandle(process_id)
	_, err := windows_api.Exe(handle)

	windows_api.CloseProcess(handle)

	return err == nil
}

func (w Window) GetWindowTitle() string {
	return w32.GetWindowText(w.hwnd)
}

func (w Window) HasTitle() bool {
	title := w.GetWindowTitle()

	if title == "Program Manager" || title == "NVIDIA GeForce Overlay" || title == "Microsoft Text Input Application" {
		return false
	}

	return len(strings.TrimSpace(title)) > 0
}

func (w Window) IsVisible() bool {
	return w32.IsWindowVisible(w.hwnd)
}

func (w Window) IsWindow() bool {
	return w32.IsWindow(w.hwnd)
}

func (w *Window) Center(work_area *rect.Rect) {

	half_width := work_area.Width / 2
	half_height := work_area.Height / 2

	x := work_area.X + ((work_area.Width - half_width) / 2)
	y := work_area.Y + ((work_area.Height - half_height) / 2)

	layout := rect.New(x, y, half_width, half_height)

	w.SetPosition(layout, true)
}

func NewWindow(hwnd w32.HWND) *Window {
	return &Window{
		hwnd: hwnd,
	}
}

func (w *Window) AlignEnd(work_area rect.Rect) {

	half_width := work_area.Width / 2

	half_height := work_area.Height / 2

	end_position := rect.Rect{
		X: 0,
		Y: half_height,

		Width:  half_width,
		Height: half_height,
	}

	w.SetPosition(end_position, true)
}

func (w *Window) AlignStart(work_area rect.Rect) {

	half_width := work_area.Width / 2
	half_monitor_height := work_area.Height / 2

	end_position := rect.Rect{
		X: 0,
		Y: 0,

		Width:  half_width,
		Height: half_monitor_height,
	}

	w.SetPosition(end_position, true)
}

func (w *Window) JustifyEnd(work_area rect.Rect) {
	// monitor_end := work_area.Width

	half_monitor_width := work_area.Width / 2

	end_position := rect.Rect{
		X: half_monitor_width,
		Y: 0,

		Width:  half_monitor_width,
		Height: work_area.Height,
	}

	w.SetPosition(end_position, true)
}
