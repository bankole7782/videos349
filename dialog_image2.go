package main

import (
	"image"
	"slices"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func drawViewAIS(window *glfw.Window, currentFrame image.Image) {
	VAISObjCoords = make(map[int]g143.Rect)
	VAISInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &VAISObjCoords)

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
	addBtnRect := theCtx.drawButtonA(VAIS_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(VAIS_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

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

func vAISMouseCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VAISObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := GetRootPath()

	clearIndicators := func(window *glfw.Window) {
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		aBInputRS := VAISObjCoords[VAIS_AudioBeginInput]
		aEInputRS := VAISObjCoords[VAIS_AudioEndInput]

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(aBInputRS.OriginX)+float64(aBInputRS.Width)+20, float64(aBInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(aEInputRS.OriginX)+float64(aEInputRS.Width)+20, float64(aEInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case VAIS_CloseBtn:
		IsUpdateDialog = false
		IsInsertBeforeDialog = false

		drawItemsView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(iVScrollBtnCB)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAIS_SelectImg:
		filename := PickImageFile()
		if filename == "" {
			return
		}
		VAISInputsStore["image"] = filename
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		theCtx := Continue2dCtx(CurrentWindowFrame, &VAISObjCoords)
		theCtx.drawFileInput(VAIS_SelectImg, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAIS_SelectAudio:
		filename := PickAudioFile()
		if filename == "" {
			return
		}
		VAISInputsStore["audio"] = filename

		theCtx := Continue2dCtx(CurrentWindowFrame, &VAISObjCoords)
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		theCtx.drawFileInput(VAIS_SelectAudio, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// update end str
		ffprobe := GetFFPCommand()
		eIRect := VAISObjCoords[VAIS_AudioEndInput]
		videoLength := LengthOfVideo(filename, ffprobe)
		VAISEndInputEnteredTxt = videoLength
		theCtx.drawInput(VAIS_AudioEndInput, eIRect.OriginX, eIRect.OriginY, eIRect.Width, VAISEndInputEnteredTxt, true)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAIS_AudioBeginInput:
		VAIS_SelectedInput = VAIS_AudioBeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAIS_AudioEndInput:
		VAIS_SelectedInput = VAIS_AudioEndInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAIS_AddBtn:

		if IsUpdateDialog {
			oldInstr := Instructions[ToUpdateInstrNum]
			if imagePath, ok := oldInstr["image"]; ok {
				oldInstr["image"] = imagePath
			}
			if audioPath, ok := oldInstr["audio"]; ok {
				oldInstr["audio"] = audioPath
			}
			if VAISBeginInputEnteredTxt != "" {
				oldInstr["audio_begin"] = VAISBeginInputEnteredTxt
				VAISBeginInputEnteredTxt = ""
			}

			if VAISEndInputEnteredTxt != "" {
				oldInstr["audio_end"] = VAISEndInputEnteredTxt
				VAISEndInputEnteredTxt = ""
			}
			Instructions[ToUpdateInstrNum] = oldInstr
			ToUpdateInstrNum = 0
			IsUpdateDialog = false

		} else {
			if VAISInputsStore["image"] == "" {
				return
			}

			if VAISInputsStore["audio"] == "" {
				return
			}

			if VAISBeginInputEnteredTxt == "" {
				VAISInputsStore["audio_begin"] = "0:00"
			} else {
				VAISInputsStore["audio_begin"] = VAISBeginInputEnteredTxt
				VAISBeginInputEnteredTxt = ""
			}

			if VAISEndInputEnteredTxt == "" {
				VAISInputsStore["audio_end"] = "5"
			} else {
				VAISInputsStore["audio_end"] = VAISEndInputEnteredTxt
				VAISEndInputEnteredTxt = ""
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":        "image",
					"image":       VAISInputsStore["image"],
					"audio":       VAISInputsStore["audio"],
					"audio_begin": VAISInputsStore["audio_begin"],
					"audio_end":   VAISInputsStore["audio_end"],
				}
				Instructions = slices.Insert(Instructions, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				Instructions = append(Instructions, map[string]string{
					"kind":        "image",
					"image":       VAISInputsStore["image"],
					"audio":       VAISInputsStore["audio"],
					"audio_begin": VAISInputsStore["audio_begin"],
					"audio_end":   VAISInputsStore["audio_end"],
				})

			}

		}

		drawItemsView(window, TotalPages())
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}

func vAISKeyCB(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if VAIS_SelectedInput == VAIS_AudioBeginInput {
		// enforce number types
		if IsKeyNumeric(key) {
			VAISBeginInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VAISBeginInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VAISBeginInputEnteredTxt) != 0 {
			VAISBeginInputEnteredTxt = VAISBeginInputEnteredTxt[:len(VAISBeginInputEnteredTxt)-1]
		}

		aBRect := VAISObjCoords[VAIS_AudioBeginInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VAISObjCoords)
		theCtx.drawInput(VAIS_AudioBeginInput, aBRect.OriginX, aBRect.OriginY, aBRect.Width, VAISBeginInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if VAIS_SelectedInput == VAIS_AudioEndInput {
		// enforce number types
		if IsKeyNumeric(key) {
			VAISEndInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VAISEndInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VAISEndInputEnteredTxt) != 0 {
			VAISEndInputEnteredTxt = VAISEndInputEnteredTxt[:len(VAISEndInputEnteredTxt)-1]
		}

		aEIRect := VAISObjCoords[VAIS_AudioEndInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VAISObjCoords)
		theCtx.drawInput(VAIS_AudioEndInput, aEIRect.OriginX, aEIRect.OriginY, aEIRect.Width, VAISEndInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}

}
