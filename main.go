package main

func main() {
	window_manager := NewWindowManager()

	// window_manager.Monitors.CenterWindow()

	monitor := window_manager.Monitors

	monitor.UpdateWorkspace()
	// window_manager.JustifyEnd()

	// window_manager.NextWindow()
}
