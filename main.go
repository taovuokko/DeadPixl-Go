package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	window, err := glfw.CreateWindow(mode.Width, mode.Height, "go-pixl", monitor, nil)
	if err != nil {
		log.Fatalf("could not create GLFW window: %v", err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		log.Fatalf("could not initialize GL: %v", err)
	}

	setFixedClearColor()

	window.SetKeyCallback(handleKeyboardInput)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		glfw.PollEvents()
		window.SwapBuffers()
	}
}
