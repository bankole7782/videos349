package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps      = 10
	fontSize = 20
)

var objCoords map[int]g143.RectSpecs
var currentWindowFrame image.Image

var allFiles []string

func main() {
	if len(os.Args) != 3 {
		panic("expecting a rootpath and a extension")
	}
	rootPath := os.Args[1]
	exts := os.Args[2]
	extsParts := strings.Split(exts, "|")

	toPickFrom := make([]string, 0)
	for _, extPart := range extsParts {
		tmp := GetFilesOfType(rootPath, extPart)
		toPickFrom = append(toPickFrom, tmp...)
	}

	slices.Sort(toPickFrom)
	allFiles = toPickFrom

	runtime.LockOSThread()

	objCoords = make(map[int]g143.RectSpecs)

	window := g143.NewWindow(1200, 800, "Videos349 file picker", false)
	allDraws(window, toPickFrom)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func GetFilesOfType(rootPath, ext string) []string {
	dirFIs, err := os.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}
	files := make([]string, 0)
	for _, dirFI := range dirFIs {
		if !dirFI.IsDir() && !strings.HasPrefix(dirFI.Name(), ".") && strings.HasSuffix(dirFI.Name(), ext) {
			files = append(files, dirFI.Name())
		}

		if dirFI.IsDir() && !strings.HasPrefix(dirFI.Name(), ".") {
			innerDirFIs, _ := os.ReadDir(filepath.Join(rootPath, dirFI.Name()))
			innerFiles := make([]string, 0)

			for _, innerDirFI := range innerDirFIs {
				if !innerDirFI.IsDir() && !strings.HasPrefix(innerDirFI.Name(), ".") && strings.HasSuffix(innerDirFI.Name(), ext) {
					innerFiles = append(innerFiles, filepath.Join(dirFI.Name(), innerDirFI.Name()))
				}
			}

			if len(innerFiles) > 0 {
				files = append(files, innerFiles...)
			}
		}

	}

	return files
}

func allDraws(window *glfw.Window, files []string) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// load font
	fontPath := getDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	currentX := 20
	currentY := 20

	for i, aFile := range files {
		aFileStrW, _ := ggCtx.MeasureString(aFile)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(aFile, float64(currentX), float64(currentY)+fontSize)
		aFileRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY, Width: int(aFileStrW), Height: fontSize}
		objCoords[i+1] = aFileRS

		newX := currentX + int(aFileStrW) + 20
		if newX > (wWidth - int(aFileStrW)) {
			currentY += 40
			currentX = 20
		} else {
			currentX += int(aFileStrW) + 20 + 10
		}

	}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func getDefaultFontPath() string {
	fontPath := filepath.Join(os.TempDir(), "fpicker_font.ttf")
	os.WriteFile(fontPath, DefaultFont, 0777)
	return fontPath
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

	fmt.Println(allFiles[widgetCode-1])
}
