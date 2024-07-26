package main

func main() {
	window_manager := NewWindowManager()

	monitor := window_manager.Monitors

	monitor.UpdateWorkspace()

}
