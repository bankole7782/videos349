package main

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	rootPath, err := GetRootPath()
	if err != nil {
		panic(err)
	}

	outPath := filepath.Join(rootPath, "renders")
	os.MkdirAll(outPath, 0777)

	runtime.LockOSThread()

	Instructions = make([]map[string]string, 0)
	InChannel = make(chan bool)

	window := g143.NewWindow(1200, 800, ProgTitle, false)
	DrawBeginView(window)

	ffmpegPath := GetFFMPEGCommand()
	ffprobePath := GetFFPCommand()
	go func() {
		for {
			<-InChannel
			Render(Instructions, ffmpegPath, ffprobePath)
			ClearAfterRender = true
		}
	}()

	// respond to the mouse
	window.SetMouseButtonCallback(projViewMouseCallback)
	// respond to the keyboard
	window.SetKeyCallback(ProjKeyCallback)
	// save the project file
	window.SetCloseCallback(SaveProjectCloseCallback)
	// quick hover effect
	window.SetCursorPosCallback(getHoverCB(ProjObjCoords, CurrentWindowFrame))

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		if ClearAfterRender {
			// clear the UI and redraw
			// Instructions = make([]map[string]string, 0)
			DrawWorkView(window, 1)
			DrawEndRenderView(window, CurrentWindowFrame)
			time.Sleep(5 * time.Second)
			DrawWorkView(window, 1)
			// register the ViewMain mouse callback
			window.SetMouseButtonCallback(workViewMouseBtnCallback)
			// quick hover effect
			window.SetCursorPosCallback(getHoverCB(ProjObjCoords, CurrentWindowFrame))
			ClearAfterRender = false
		}

		time.Sleep(time.Second/time.Duration(FPS) - time.Since(t))
	}

}
