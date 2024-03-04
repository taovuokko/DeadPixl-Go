package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var currentColorIndex = 0

var colors = [][]float32{
	{1.0, 1.0, 1.0, 1.0}, // White
	{0.0, 0.0, 0.0, 1.0}, // Black
	{1.0, 0.0, 0.0, 1.0}, // Red
	{0.0, 0.0, 1.0, 1.0}, // Blue
}

func setFixedClearColor() {
	color := colors[currentColorIndex]

	// Set the clear color
	gl.ClearColor(color[0], color[1], color[2], color[3])

	currentColorIndex = (currentColorIndex + 1) % len(colors)
}
