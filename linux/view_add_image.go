package main

import (
	"image"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/videos349/v3shared"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	VAI_SelectImg   = 21
	VAI_SelectAudio = 22
	VAI_DurInput    = 23
	VAI_AddBtn      = 24
	VAI_CloseBtn    = 25
)

var vaiObjCoords map[int]g143.RectSpecs
var vaiInputsStore map[string]string
var vaiEnteredText string

func drawViewAddImage(window *glfw.Window, currentFrame image.Image) {
	vaiObjCoords = make(map[int]g143.RectSpecs)
	vaiInputsStore = make(map[string]string)

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
	dialogHeight := 450

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth), float64(dialogHeight))
	ggCtx.Fill()

	// Add Image
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Add Image Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	// Add Btn
	ggCtx.SetHexColor("#56845A")
	addStrW, _ := ggCtx.MeasureString("Add")
	addBtnOriginX := dialogWidth - int(addStrW) - 50 + dialogOriginX
	ggCtx.DrawRoundedRectangle(float64(addBtnOriginX), float64(dialogOriginY)+20, addStrW+20, 30, 10)
	ggCtx.Fill()
	addBtnRS := g143.RectSpecs{OriginX: addBtnOriginX, OriginY: dialogOriginY + 20, Width: int(addStrW) + 20, Height: 30}
	vaiObjCoords[VAI_AddBtn] = addBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Add", float64(addBtnOriginX)+10, float64(dialogOriginY)+20+20)

	// Close Btn
	ggCtx.SetHexColor("#B75F5F")
	closeStrW, _ := ggCtx.MeasureString("Close")
	closeBtnX := dialogOriginX + dialogWidth - 50 - int(addStrW) - 30 - int(closeStrW)
	ggCtx.DrawRoundedRectangle(float64(closeBtnX), float64(dialogOriginY)+20, closeStrW+20, 30, 10)
	ggCtx.Fill()
	closeBtnRS := g143.RectSpecs{OriginX: closeBtnX, OriginY: dialogOriginY + 20, Width: int(closeStrW) + 20, Height: 30}
	vaiObjCoords[VAI_CloseBtn] = closeBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Close", float64(closeBtnX)+10, float64(dialogOriginY)+20+20)

	// add image form
	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX)+40, float64(dialogOriginY)+20+50, float64(dialogWidth)-80, 240, 10)
	ggCtx.Fill()
	selectImgRS := g143.RectSpecs{Width: dialogWidth - 80, Height: 240, OriginX: dialogOriginX + 40, OriginY: dialogOriginY + 20 + 50}
	vaiObjCoords[VAI_SelectImg] = selectImgRS

	aicStr := "click to pick an image"
	ggCtx.SetHexColor("#444")
	aicStrW, _ := ggCtx.MeasureString(aicStr)
	aicStrOriginX := selectImgRS.OriginX + (selectImgRS.Width-int(aicStrW))/2
	ggCtx.DrawString(aicStr, float64(aicStrOriginX), float64(selectImgRS.OriginY)+20+20)

	sfl := "audio file (optional)"
	sflW, _ := ggCtx.MeasureString(sfl)
	sflY := selectImgRS.OriginY + selectImgRS.Height + 10 + 20
	ggCtx.DrawString(sfl, float64(dialogOriginX)+40, float64(sflY))

	ggCtx.SetHexColor("#eee")
	sflBtnX := sflW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(sflBtnX, float64(sflY)-20, float64(dialogWidth)/2, 30, 10)
	ggCtx.Fill()
	sflBtnRS := g143.RectSpecs{Width: dialogWidth / 2, Height: 30, OriginX: int(sflBtnX), OriginY: sflY - 20}
	vaiObjCoords[VAI_SelectAudio] = sflBtnRS

	durLabel := "duration (in seconds)"
	durLabelW, _ := ggCtx.MeasureString(durLabel)
	durLabelY := sflBtnRS.OriginY + sflBtnRS.Height + 10 + 20
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(durLabel, float64(dialogOriginX)+40, float64(durLabelY))

	ggCtx.SetHexColor("#eee")
	durInputX := durLabelW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(durInputX, float64(durLabelY)-20, 200, 30, 10)
	ggCtx.Fill()
	durInputRS := g143.RectSpecs{OriginX: int(durInputX), OriginY: durLabelY - 20, Width: 100, Height: 30}
	vaiObjCoords[VAI_DurInput] = durInputRS

	// default duration
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("5", durInputX+20, float64(durLabelY))

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func viewAddImageMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range vaiObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := v3shared.GetRootPath()

	switch widgetCode {
	case VAI_CloseBtn:
		allDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case VAI_SelectImg:
		filename := pickFileUbuntu("png|jpg")
		if filename == "" {
			return
		}
		vaiInputsStore["image"] = filename

		// show picked image
		ggCtx := gg.NewContextForImage(currentWindowFrame)

		img, _ := imaging.Open(filename)
		img = imaging.Fit(img, widgetRS.Width-20, widgetRS.Height-20, imaging.Lanczos)
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY), float64(widgetRS.Width),
			float64(widgetRS.Height), 10)
		ggCtx.Fill()
		ggCtx.DrawImage(img, widgetRS.OriginX+10, widgetRS.OriginY+10)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	case VAI_SelectAudio:
		filename := pickFileUbuntu("mp3")
		if filename == "" {
			return
		}
		vaiInputsStore["audio_optional"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(currentWindowFrame)
		// load font
		fontPath := getDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY),
			float64(widgetRS.Width), float64(widgetRS.Height), 10)
		ggCtx.Fill()

		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(displayFilename, float64(widgetRS.OriginX)+10, float64(widgetRS.OriginY)+20)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()

	case VAI_AddBtn:
		if vaiInputsStore["image"] == "" {
			return
		}

		if vaiEnteredText == "" {
			vaiInputsStore["duration"] = "5"
		} else {
			vaiInputsStore["duration"] = vaiEnteredText
			vaiEnteredText = ""
		}

		instructions = append(instructions, map[string]string{
			"kind":           "image",
			"image":          vaiInputsStore["image"],
			"duration":       vaiInputsStore["duration"],
			"audio_optional": vaiInputsStore["audio_optional"],
		})

		allDraws(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}

func vaikeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	// enforce number types
	if isKeyNumeric(key) {
		vaiEnteredText += glfw.GetKeyName(key, scancode)
	} else if key == glfw.KeyBackspace && len(vaiEnteredText) != 0 {
		vaiEnteredText = vaiEnteredText[:len(vaiEnteredText)-1]
	}

	ggCtx := gg.NewContextForImage(currentWindowFrame)
	// load font
	fontPath := getDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	durInputRS := vaiObjCoords[VAI_DurInput]

	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRoundedRectangle(float64(durInputRS.OriginX), float64(durInputRS.OriginY), float64(durInputRS.Width),
		float64(durInputRS.Height), 10)
	ggCtx.Fill()

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(vaiEnteredText, float64(durInputRS.OriginX+10), float64(durInputRS.OriginY)+20)

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}
