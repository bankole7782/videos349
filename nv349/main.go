package main

import (
	"image"
	"os"
	"path/filepath"
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10
)

var objCoords map[int]g143.RectSpecs
var currentWindowFrame image.Image

func main() {
	runtime.LockOSThread()

	objCoords = make(map[int]g143.RectSpecs)

	window := g143.NewWindow(1200, 800, "videos349: a simple video editor", false)
	allDraws(window)

	// // respond to the mouse
	// window.SetMouseButtonCallback(mouseBtnCallback)
	// window.SetKeyCallback(keyCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}

}

func getDefaultFontPath() string {
	fontPath := filepath.Join(os.TempDir(), "v349_font.ttf")
	os.WriteFile(fontPath, DefaultFont, 0777)
	return fontPath
}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// // intro text
	fontPath := getDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()

}
