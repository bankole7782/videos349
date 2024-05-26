package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/videos349/internal"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	_, err := internal.GetRootPath()
	if err != nil {
		panic(err)
	}

	runtime.LockOSThread()

	internal.Instructions = make([]map[string]string, 0)
	internal.InChannel = make(chan bool)

	window := g143.NewWindow(1200, 800, "videos349: a simple video editor for teachers", false)
	internal.DrawBeginView(window)

	ffmpegPath := GetFFMPEGCommand()
	ffprobePath := GetFFPCommand()
	go func() {
		for {
			<-internal.InChannel
			internal.Render(internal.Instructions, ffmpegPath, ffprobePath)
			internal.ClearAfterRender = true
		}
	}()

	// respond to the mouse
	window.SetMouseButtonCallback(workViewMouseBtnCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		if internal.ClearAfterRender {
			// clear the UI and redraw
			internal.Instructions = make([]map[string]string, 0)
			internal.AllDraws(window)
			internal.DrawEndRenderView(window, internal.CurrentWindowFrame)
			time.Sleep(5 * time.Second)
			internal.AllDraws(window)
			// register the ViewMain mouse callback
			window.SetMouseButtonCallback(workViewMouseBtnCallback)
			internal.ClearAfterRender = false
		}

		time.Sleep(time.Second/time.Duration(internal.FPS) - time.Since(t))
	}

}
