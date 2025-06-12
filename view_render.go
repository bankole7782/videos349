package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawRenderView(window *glfw.Window, currentFrame image.Image, percentage float64) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	ggCtx.DrawImage(img, 0, 0)

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 300

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth), float64(dialogHeight))
	ggCtx.Fill()

	// Add Image
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Rendering! Please Wait", float64(dialogOriginX)+20, float64(dialogOriginY)+FontSize+20)

	progressX := float64(dialogOriginX) + 20
	progressY := float64(dialogOriginY) + 80
	widthOfBar := 600 - 40

	// background rectangle
	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRectangle(progressX, progressY, float64(widthOfBar), 20)
	ggCtx.Fill()

	// progress rectangle
	progressBarX := progressX + 5 + float64(50*LastIndicatorIndex)
	ggCtx.SetHexColor("#444")
	ggCtx.DrawRectangle(progressBarX, progressY+5, 100, 10)
	ggCtx.Fill()

	if LastIndicatorIndex == MaxIndicatorIndex {
		LastIndicatorIndex = 0
	} else {
		LastIndicatorIndex += 1
	}

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}

func DrawEndRenderView(window *glfw.Window, currentFrame image.Image) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	ggCtx.DrawImage(img, 0, 0)

	// load font
	fontPath := GetDefaultFontPath()
	ggCtx.LoadFontFace(fontPath, 20)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 300

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth), float64(dialogHeight))
	ggCtx.Fill()

	if RenderErrorHappened {
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString("An error occurred", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

		rootPath, _ := GetRootPath()
		logsPath := filepath.Join(rootPath, "errors", UntestedRandomString(10)+"_log.txt")
		os.WriteFile(logsPath, []byte(RenderErrorMsg), 0777)

		ggCtx.DrawString(fmt.Sprintf("View log %s", logsPath), float64(dialogOriginX)+20, float64(dialogOriginY)+40+30)
	} else {
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString("Done Rendering! Open the Working Directory", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)
	}

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}
