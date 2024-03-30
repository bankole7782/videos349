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
	VAV_PickAudio  = 36
)

var vavObjCoords map[int]g143.RectSpecs
var vavInputsStore map[string]string

var beginInputEnteredTxt string = "0:00"
var endInputEnteredTxt string = "0:00"
var selectedInput int

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

	// Add Video Header
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

	// add audio input
	sfl := "audio file (optional)"
	sflW, _ := ggCtx.MeasureString(sfl)
	sflY := eiRS.OriginY + eiRS.Height + 10 + 20
	ggCtx.DrawString(sfl, float64(dialogOriginX)+40, float64(sflY))

	ggCtx.SetHexColor("#eee")
	sflBtnX := sflW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(sflBtnX, float64(sflY)-20, float64(dialogWidth)/2, 30, 10)
	ggCtx.Fill()
	sflBtnRS := g143.RectSpecs{Width: dialogWidth / 2, Height: 30, OriginX: int(sflBtnX), OriginY: sflY - 20}
	vavObjCoords[VAV_PickAudio] = sflBtnRS

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

	clearIndicators := func(window *glfw.Window) {
		ggCtx := gg.NewContextForImage(currentWindowFrame)

		beginInputRS := vavObjCoords[VAV_BeginInput]
		endInputRS := vavObjCoords[VAV_EndInput]

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(beginInputRS.OriginX)+float64(beginInputRS.Width)+20, float64(beginInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(endInputRS.OriginX)+float64(endInputRS.Width)+20, float64(endInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case VAV_CloseBtn:
		allDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case VAV_PickVideo:
		filename, err := dialog.File().Filter("MP4 Video", "mp4").Filter("WEBM Video", "webm").Filter("MKV Video", "mkv").Load()
		if filename == "" || err != nil {
			return
		}
		vavInputsStore["video"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err = ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY),
			float64(widgetRS.Width), float64(widgetRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(filepath.Base(filename), float64(widgetRS.OriginX+10), float64(widgetRS.OriginY+20))

		// update end str
		endInputRS := vavObjCoords[VAV_EndInput]
		videoLength := lengthOfVideo(filename)
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

	case VAV_BeginInput:
		selectedInput = VAV_BeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(currentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	case VAV_EndInput:
		selectedInput = VAV_EndInput
		clearIndicators(window)

		ggCtx := gg.NewContextForImage(currentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	case VAV_PickAudio:
		filename, err := dialog.File().Filter("MP3 Audio", "mp3").Load()
		if filename == "" || err != nil {
			return
		}
		vavInputsStore["audio_optional"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err = ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY),
			float64(widgetRS.Width), float64(widgetRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(filepath.Base(filename), float64(widgetRS.OriginX)+10, float64(widgetRS.OriginY)+20)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	case VAV_AddBtn:
		if vavInputsStore["video"] == "" {
			return
		}

		instructions = append(instructions, map[string]string{
			"kind":           "video",
			"video":          vavInputsStore["video"],
			"begin":          beginInputEnteredTxt,
			"end":            endInputEnteredTxt,
			"audio_optional": vavInputsStore["audio_optional"],
		})

		allDraws(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}

func vavkeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if selectedInput == VAV_BeginInput {
		beginInputRS := vavObjCoords[VAV_BeginInput]

		// enforce number types, semicolon and backspace
		if isKeyNumeric(key) {
			beginInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			beginInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(beginInputEnteredTxt) != 0 {
			beginInputEnteredTxt = beginInputEnteredTxt[:len(beginInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(beginInputRS.OriginX), float64(beginInputRS.OriginY), float64(beginInputRS.Width),
			float64(beginInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(beginInputEnteredTxt, float64(beginInputRS.OriginX+10), float64(beginInputRS.OriginY)+fontSize)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	} else if selectedInput == VAV_EndInput {
		endInputRS := vavObjCoords[VAV_EndInput]

		// enforce number types, semicolon and backspace
		if isKeyNumeric(key) {
			endInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			endInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(endInputEnteredTxt) != 0 {
			endInputEnteredTxt = endInputEnteredTxt[:len(endInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY), float64(endInputRS.Width),
			float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(endInputEnteredTxt, float64(endInputRS.OriginX+10), float64(endInputRS.OriginY)+fontSize)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	}
}
