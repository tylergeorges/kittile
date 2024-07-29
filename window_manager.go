package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/gonutz/w32/v2"
	"github.com/tylergeorges/kittile/layout"
)

type Msg interface{}

type FlipVertMsg struct{}
type QuitMsg struct{}

type WindowManager struct {
	Monitors *Monitor
	msgs     chan Msg
	quitting bool
}

func (w *WindowManager) Run() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	m := w.Monitors
	ws := m.workspace

	tree := ws.GetTree()

	tree.FlipTree(layout.FlipVertical)
	tree.Render()
	// tree.Render()

	// tree.FlipTree(layout.FlipVertical)
	// tree.Render()

	go func() {
		for {
			_, key, err := keyboard.GetKey()

			if err != nil {
				panic(err)
			}

			fmt.Printf("key: %v\n", key)

			if key == keyboard.KeyArrowLeft {
				w.msgs <- FlipVert()
			}

			if key == keyboard.KeyEsc {
				w.msgs <- Quit()
				w.quitting = true
			}

			if w.quitting {
				break
			}
		}

	}()

	for msg := range w.msgs {

		switch msg.(type) {
		case FlipVertMsg:
			tree.FlipTree(layout.FlipVertical)
			tree.Render()

			fmt.Print("flip tree")

		case QuitMsg:
			return
		}
	}

}

func FlipVert() Msg {
	return FlipVertMsg{}
}

func Quit() Msg {
	return QuitMsg{}
}

func (w *WindowManager) listen_to_keyboard() {
	fmt.Println("Press ESC to quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {

		_, key, err := keyboard.GetKey()

		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			// w.msgs <- Quit()
			break
		}

		// select {
		// case msg := <-w.msgs:
		// 	if msg == nil {
		// 		continue
		// 	}

		// 	switch msg.(type) {
		// 	case QuitMsg:
		// 		fmt.Print("quit")
		// 		return
		// 	}
		// }

	}
}

func NewWindowManager() *WindowManager {
	w := w32.GetForegroundWindow()

	m := NewMonitor(w)

	return &WindowManager{
		Monitors: m,
		msgs:     make(chan Msg),
		quitting: false,
	}
}
