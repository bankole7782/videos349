package main

import (
	"image"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sqweek/dialog"
)

const (
	VAV_AddBtn     = 31
	VAV_CloseBtn   = 32
	VAV_PickVideo  = 33
	VAV_BeginInput = 34
	VAV_EndInput   = 35
)

var vavObjCoords map[int]g143.RectSpecs
var vavInputsStore map[string]string

var beginInputEnteredTxt string = "0:00"
var endInputEnteredTxt string = "0:00"

func drawViewAddVideo(window *glfw.Window, currentFrame image.Image) {
	vavObjCoords = make(map[int]g143.RectSpecs)
	vavInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	ggCtx.DrawImage(img, 0, 0)

	// load font
	fontPath := getDefaultFontPath()
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
	ggCtx.DrawString("Add Video Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	// Add Btn
	ggCtx.SetHexColor("#56845A")
	addStrW, _ := ggCtx.MeasureString("Add")
	addBtnOriginX := dialogWidth - int(addStrW) - 50 + dialogOriginX
	ggCtx.DrawRoundedRectangle(float64(addBtnOriginX), float64(dialogOriginY)+20, addStrW+20, 30, 10)
	ggCtx.Fill()
	addBtnRS := g143.RectSpecs{OriginX: addBtnOriginX, OriginY: dialogOriginY + 20, Width: int(addStrW) + 20, Height: 30}
	vavObjCoords[VAV_AddBtn] = addBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Add", float64(addBtnOriginX)+10, float64(dialogOriginY)+20+20)

	// Close Btn
	ggCtx.SetHexColor("#B75F5F")
	closeStrW, _ := ggCtx.MeasureString("Close")
	closeBtnX := dialogOriginX + dialogWidth - 50 - int(addStrW) - 30 - int(closeStrW)
	ggCtx.DrawRoundedRectangle(float64(closeBtnX), float64(dialogOriginY)+20, closeStrW+20, 30, 10)
	ggCtx.Fill()
	closeBtnRS := g143.RectSpecs{OriginX: closeBtnX, OriginY: dialogOriginY + 20, Width: int(closeStrW) + 20, Height: 30}
	vavObjCoords[VAV_CloseBtn] = closeBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Close", float64(closeBtnX)+10, float64(dialogOriginY)+20+20)
	// end of top bar

	// pick video label and btn
	ggCtx.SetHexColor("#444")
	pvStr := "video file"
	pvStrW, _ := ggCtx.MeasureString(pvStr)
	pvStrX := dialogOriginX + 40
	pvStrY := dialogOriginY + fontSize + 70
	ggCtx.DrawString(pvStr, float64(pvStrX), float64(pvStrY))

	ggCtx.SetHexColor("#eee")
	pvInputX := pvStrW + 30 + float64(dialogOriginX) + 40
	pvInputW := dialogWidth - int(pvStrW) - 90

	ggCtx.DrawRoundedRectangle(pvInputX, float64(pvStrY)-fontSize, float64(pvInputW), 60, 10)
	ggCtx.Fill()
	pfRS := g143.RectSpecs{OriginX: int(pvInputX), OriginY: pvStrY - fontSize, Width: pvInputW, Height: 60}
	vavObjCoords[VAV_PickVideo] = pfRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("[click to pick file]", pvInputX+10, float64(pvStrY))

	// begin str label and input
	beginStr := "begin (mm:ss)"
	beginStrW, _ := ggCtx.MeasureString(beginStr)
	beginStrX := dialogOriginX + 40
	beginStrY := pvStrY + fontSize + 60
	ggCtx.DrawString(beginStr, float64(beginStrX), float64(beginStrY))

	ggCtx.SetHexColor("#eee")
	beginInputX := beginStrW + 30 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(beginInputX, float64(beginStrY)-fontSize, 100, 30, 10)
	ggCtx.Fill()
	biRS := g143.RectSpecs{OriginX: int(beginInputX), OriginY: beginStrY - fontSize, Width: 100, Height: 30}
	vavObjCoords[VAV_BeginInput] = biRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", beginInputX+10, float64(beginStrY))

	// end str label and input
	endStr := "end (mm:ss)"
	endStrW, _ := ggCtx.MeasureString(endStr)
	endStrX := dialogOriginX + 40
	endStrY := beginStrY + fontSize + 30
	ggCtx.DrawString(endStr, float64(endStrX), float64(endStrY))

	ggCtx.SetHexColor("#eee")
	endInputX := endStrW + 30 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(endInputX, float64(endStrY)-fontSize, 100, 30, 10)
	ggCtx.Fill()
	eiRS := g143.RectSpecs{OriginX: int(endInputX), OriginY: endStrY - fontSize, Width: 100, Height: 30}
	vavObjCoords[VAV_EndInput] = eiRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", endInputX+10, float64(endStrY))

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func viewAddVideoMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range vavObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case VAV_CloseBtn:
		allDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case VAV_PickVideo:
		filePath, err := dialog.File().Filter("Video MP4", "mp4").Filter("Video MKV", "mkv").
			Filter("Video WEBM", "webm").Load()
		if err != nil {
			return
		}
		vavInputsStore["video"] = filePath

		// write audio name
		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err = ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}
		fileName := filepath.Base(filePath)
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY),
			float64(widgetRS.Width), float64(widgetRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(fileName, float64(widgetRS.OriginX+10), float64(widgetRS.OriginY+20))

		// update end str
		endInputRS := vavObjCoords[VAV_EndInput]
		videoLength := lengthOfVideo(filePath)
		endInputEnteredTxt = videoLength

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY),
			float64(endInputRS.Width), float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(videoLength, float64(endInputRS.OriginX)+10, float64(endInputRS.OriginY+fontSize))

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	}
}

func vavkeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

}
