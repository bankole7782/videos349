package main

import (
	"image"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10

	AddImgBtn = 1
	AddVidBtn = 2
	OpenWDBtn = 3
)

var objCoords map[int]g143.RectSpecs
var currentWindowFrame image.Image

func main() {
	_, err := GetRootPath()
	if err != nil {
		panic(err)
	}

	runtime.LockOSThread()

	objCoords = make(map[int]g143.RectSpecs)

	window := g143.NewWindow(1200, 800, "videos349: a simple video editor", false)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
	// respond to the keyboard
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

	// add image button
	addImgStr := "Add Image"
	addImgStrWidth, addImgStrHeight := ggCtx.MeasureString(addImgStr)
	addImgBtnWidth := addImgStrWidth + 80
	addImgBtnHeight := addImgStrHeight + 30
	ggCtx.SetHexColor("#B75F5F")
	ggCtx.DrawRoundedRectangle(10, 10, addImgBtnWidth, addImgBtnHeight, addImgBtnHeight/2)
	ggCtx.Fill()

	addImgBtnRS := g143.RectSpecs{Width: int(addImgBtnWidth), Height: int(addImgBtnHeight),
		OriginX: 10, OriginY: 10}
	objCoords[AddImgBtn] = addImgBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(addImgStr, 30, 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#633232")
	ggCtx.DrawCircle(10+addImgBtnWidth-40, 10+addImgBtnHeight/2, 10)
	ggCtx.Fill()

	// Add Video Button
	addVidStr := "Add Video"
	addVidStrWidth, addVidStrHeight := ggCtx.MeasureString(addVidStr)
	addVidBtnWidth := addVidStrWidth + 80
	addVidBtnHeight := addVidStrHeight + 30
	ggCtx.SetHexColor("#81577F")
	addVidBtnOriginX := float64(addImgBtnRS.Width+addImgBtnRS.OriginX) + 10 // gutter
	ggCtx.DrawRoundedRectangle(addVidBtnOriginX, 10, addVidBtnWidth, addVidBtnHeight, addVidBtnHeight/2)
	ggCtx.Fill()

	addVidBtnRS := g143.RectSpecs{Width: int(addVidBtnWidth), Height: int(addVidBtnHeight),
		OriginX: int(addVidBtnOriginX), OriginY: 10}
	objCoords[AddVidBtn] = addVidBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(addVidStr, 30+float64(addVidBtnRS.OriginX), 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#633260")
	ggCtx.DrawCircle(float64(addVidBtnRS.OriginX)+addVidBtnWidth-30, 10+addVidBtnHeight/2, 10)
	ggCtx.Fill()

	// Open Working Directory button
	owdStr := "Open Working Directory"
	owdStrWidth, owdStrHeight := ggCtx.MeasureString(owdStr)
	openWDBtnWidth := owdStrWidth + 30
	openWDBtnHeight := owdStrHeight + 30
	ggCtx.SetHexColor("#56845A")
	openWDBtnOriginX := float64(addVidBtnRS.OriginX+addVidBtnRS.Width) + 40
	ggCtx.DrawRoundedRectangle(openWDBtnOriginX, 10, openWDBtnWidth, openWDBtnHeight, openWDBtnHeight/2)
	ggCtx.Fill()

	openWDBtnRS := g143.RectSpecs{Width: int(openWDBtnWidth), Height: int(openWDBtnHeight),
		OriginX: int(openWDBtnOriginX), OriginY: 10}
	objCoords[OpenWDBtn] = openWDBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(owdStr, 15+float64(openWDBtnRS.OriginX), 10+owdStrHeight+15)

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()

}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range objCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case AddImgBtn:
		// ggCtx := gg.NewContextForImage(currentWindowFrame)
	case AddVidBtn:

	case OpenWDBtn:
		rootPath, _ := GetRootPath()

		if runtime.GOOS == "windows" {
			exec.Command("cmd", "/C", "start", rootPath).Run()
		} else if runtime.GOOS == "linux" {
			exec.Command("xdg-open", rootPath).Run()
		}
	}
}
