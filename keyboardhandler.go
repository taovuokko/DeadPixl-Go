package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func handleKeyboardInput(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press && action != glfw.Repeat {
		return
	}

	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
		return
	}

	setFixedClearColor()
}
