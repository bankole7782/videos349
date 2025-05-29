package main

import (
	"fmt"
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
	os.MkdirAll(filepath.Join(rootPath, "errors"), 0777)

	runtime.LockOSThread()

	Instructions = make([]map[string]string, 0)
	InChannel = make(chan bool)

	window := g143.NewWindow(1200, 800, ProgTitle, false)
	drawFirstView(window)

	ffmpegPath := GetFFMPEGCommand()
	ffprobePath := GetFFPCommand()
	go func() {
		for {
			<-InChannel
			IsRendering = true
			_, err := Render(Instructions, ffmpegPath, ffprobePath)
			if err != nil {
				RenderErrorHappened = true
				RenderErrorMsg = fmt.Sprintf("%+v", err)
			}
			IsRendering = false
			ClearAfterRender = true
		}
	}()

	// respond to the mouse
	window.SetMouseButtonCallback(fVMouseCB)
	// respond to the keyboard
	window.SetKeyCallback(fVKeyCB)
	// save the project file
	window.SetCloseCallback(SaveProjectCloseCallback)
	// quick hover effect
	window.SetCursorPosCallback(getHoverCB(ProjObjCoords))

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		if IsRendering {
			percentageRendered := float64(RenderProgress) / float64(len(Instructions))
			DrawRenderView(window, SavedWorkViewFrame, percentageRendered)
		}

		if ClearAfterRender && RenderErrorHappened {
			drawItemsView(window, 1)
			DrawEndRenderView(window, CurrentWindowFrame)

			time.Sleep(5 * time.Second)
			drawItemsView(window, 1)
			// register the ViewMain mouse callback
			window.SetMouseButtonCallback(iVMouseBtnCB)
			// quick hover effect
			window.SetCursorPosCallback(getHoverCB(ObjCoords))
			ClearAfterRender = false
			RenderErrorHappened = false

		} else if ClearAfterRender {
			// clear the UI and redraw
			// Instructions = make([]map[string]string, 0)
			drawItemsView(window, 1)
			DrawEndRenderView(window, CurrentWindowFrame)
			time.Sleep(5 * time.Second)
			drawItemsView(window, 1)
			// register the ViewMain mouse callback
			window.SetMouseButtonCallback(iVMouseBtnCB)
			// quick hover effect
			window.SetCursorPosCallback(getHoverCB(ObjCoords))
			ClearAfterRender = false
		}

		time.Sleep(time.Second/time.Duration(FPS) - time.Since(t))
	}

}
