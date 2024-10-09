package main

import (
	"image"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawViewAddImage(window *glfw.Window, currentFrame image.Image) {
	VaiObjCoords = make(map[int]g143.Rect)
	VaiInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &VaiObjCoords)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 200

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	theCtx.ggCtx.Fill()

	// Add Image
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("Add Image Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(VAI_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(VAI_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	placeholder := "[click to pick an image]"
	if IsUpdateDialog {
		filename := Instructions[ToUpdateInstrNum]["image"]
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		placeholder = displayFilename
	}
	pHRect := theCtx.drawFileInput(VAI_SelectImg, dialogOriginX+20, dialogOriginY+40+30, dialogWidth-40, placeholder)
	_, durLabelY := nextVerticalCoords(pHRect, 30)
	durLabel := "duration (in seconds)"
	durLabelW, _ := theCtx.ggCtx.MeasureString(durLabel)
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString(durLabel, float64(dialogOriginX)+20, float64(durLabelY))

	theCtx.drawInput(VAI_DurInput, dialogOriginX+int(durLabelW)+40, durLabelY-FontSize, 80, "5", true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

func DrawViewAIS(window *glfw.Window, currentFrame image.Image) {
	VaisObjCoords = make(map[int]g143.Rect)
	VaisInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &VaisObjCoords)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 300

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	theCtx.ggCtx.Fill()

	// Add Image
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("Add Image + Sound Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(VAI_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(VAI_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	// file pickers
	placeholder := "[click to pick an image]"
	if IsUpdateDialog {
		filename := Instructions[ToUpdateInstrNum]["image"]
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		placeholder = displayFilename
	}
	pHRect := theCtx.drawFileInput(VAIS_SelectImg, dialogOriginX+20, dialogOriginY+40+30, dialogWidth-40, placeholder)

	_, audioBtnY := nextVerticalCoords(pHRect, 20)

	placeholder2 := "[click to pick audio]"
	if IsUpdateDialog {
		filename := Instructions[ToUpdateInstrNum]["audio"]
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		placeholder2 = displayFilename
	}
	aPHRect := theCtx.drawFileInput(VAIS_SelectAudio, dialogOriginX+20, audioBtnY, dialogWidth-40, placeholder2)

	// audio begin
	_, audioBeginY := nextVerticalCoords(aPHRect, 30)
	aBL := "audio begin (mm:ss)"
	theCtx.ggCtx.SetHexColor("#444")
	aBLW, _ := theCtx.ggCtx.MeasureString(aBL)
	theCtx.ggCtx.DrawString(aBL, float64(dialogOriginX)+40, float64(audioBeginY))
	aBIX := dialogOriginX + 40 + int(aBLW) + 20
	value := "0:00"
	if IsUpdateDialog {
		value = Instructions[ToUpdateInstrNum]["audio_begin"]
	}
	aBRect := theCtx.drawInput(VAIS_AudioBeginInput, aBIX, audioBeginY-FontSize, 80, value, true)

	// audio end
	aEL := "audio end (mm:ss)"
	aELW, _ := theCtx.ggCtx.MeasureString(aEL)
	_, aELY := nextVerticalCoords(aBRect, 30)
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString(aEL, float64(dialogOriginX)+40, float64(aELY))
	aEIX := dialogOriginX + 40 + int(aELW) + 30
	value2 := "0:00"
	if IsUpdateDialog {
		value2 = Instructions[ToUpdateInstrNum]["audio_end"]
	}
	theCtx.drawInput(VAIS_AudioEndInput, aEIX, aELY-FontSize, 80, value2, true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()

}
