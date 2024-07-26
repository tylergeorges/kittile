package main

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/0xrawsec/golang-win32/win32"
	"github.com/0xrawsec/golang-win32/win32/kernel32"
	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/rect"
	"golang.org/x/sys/windows"
)

var (
	libuser32 = windows.NewLazySystemDLL("user32.dll")
)

func IsIconic(hWnd w32.HWND) bool {
	ret, _, _ := syscall.Syscall(libuser32.NewProc("IsIconic").Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func QueryFullProcessImageName(hProcess w32.HANDLE) (string, error) {
	text, lastErr := kernel32.QueryFullProcessImageName(win32.HANDLE(hProcess))

	return text, lastErr
}

func Exe(hwnd w32.HANDLE) (string, error) {
	exe_path, err := QueryFullProcessImageName(hwnd)

	if err != nil {
		return "", err
	}

	path_arr := strings.Split(exe_path, "\\")

	return path_arr[len(path_arr)-1], nil
}

func CloseProcess(window w32.HANDLE) {
	w32.CloseHandle(window)
}

func SetWindowPos(window, position w32.HWND, layout *rect.Rect, flags int) {
	w32.SetWindowPos(
		window,
		position,
		layout.X,
		layout.Y,
		layout.Width,
		layout.Height,
		uint(flags),
	)
}

func window_rect(hwnd w32.HWND) *w32.RECT {

	var r windows.Rect

	if hwnd == 0 {
		return w32.GetWindowRect(hwnd)
	}

	ptr := unsafe.Pointer(&r)
	size := (uint32)(unsafe.Sizeof(r))

	handle := (windows.HWND)(hwnd)

	err := windows.DwmGetWindowAttribute(handle, windows.DWMWA_EXTENDED_FRAME_BOUNDS, ptr, size)

	if err != nil {
		window_rect := w32.GetWindowRect(hwnd)
		return window_rect
	}

	return &w32.RECT{
		Left:   r.Left,
		Right:  r.Right,
		Bottom: r.Bottom,
		Top:    r.Top,
	}

}

func shadow_rect(hwnd w32.HWND) rect.Rect {
	window_rect := rect.FromRECT(*window_rect(hwnd))
	shadow_rect := rect.FromRECT(*w32.GetWindowRect(hwnd))

	return shadow_rect.Sub(window_rect)
}

func PositionWindow(window w32.HWND, layout *rect.Rect, top bool) {
	flags := w32.SWP_NOACTIVATE | w32.SWP_NOSENDCHANGING | w32.SWP_NOCOPYBITS | w32.SWP_FRAMECHANGED

	srect := shadow_rect(window)

	offset_rect := layout.Add(srect)

	SetWindowPos(
		window,
		w32.HWND_TOP,
		&offset_rect,
		flags,
	)
}

func NewPoint(x, y int) w32.POINT {
	return w32.POINT{X: int32(x), Y: int32(y)}
}

func CursorPos() w32.POINT {
	x, y, _ := w32.GetCursorPos()

	return NewPoint(x, y)
}

func MonitorFromPos(p w32.POINT) w32.HMONITOR {
	return w32.MonitorFromPoint(int(p.X), int(p.Y), w32.MONITOR_DEFAULTTOPRIMARY)
}

func MonitorAtCursor() w32.HMONITOR {
	return MonitorFromPos(CursorPos())
}

func OpenProcess(access_rights uint32, inherit_handle bool, process_id uint32) w32.HANDLE {
	return w32.OpenProcess(access_rights, inherit_handle, process_id)
}

func ProcessHandle(process_id uint32) w32.HANDLE {
	return OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, process_id)
}
