package main

import (
	"image"
	"slices"
	"strconv"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func drawViewAddVideo(window *glfw.Window, currentFrame image.Image) {
	VAVObjCoords = make(map[int]g143.Rect)
	VAVInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &VAVObjCoords)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 280

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	theCtx.ggCtx.Fill()

	// Add Video Header
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("Add Video Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(VAV_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(VAV_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	// pick video
	placeholder := "[click to pick video file]"
	if IsUpdateDialog {
		filename := Instructions[ToUpdateInstrNum]["video"]
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		placeholder = displayFilename
	}
	vFIRect := theCtx.drawFileInput(VAV_PickVideo, dialogOriginX+20, dialogOriginY+40+30, dialogWidth-40, placeholder)

	// begin str label and input
	beginStr := "begin (mm:ss)"
	beginStrW, _ := theCtx.ggCtx.MeasureString(beginStr)
	beginStrX := dialogOriginX + 40
	_, beginStrY := nextVerticalCoords(vFIRect, 20)
	theCtx.ggCtx.DrawString(beginStr, float64(beginStrX), float64(beginStrY)+FontSize)
	bIX := dialogOriginX + 40 + int(beginStrW) + 20
	value := "0:00"
	if IsUpdateDialog {
		value = Instructions[ToUpdateInstrNum]["begin"]
	}
	bIRect := theCtx.drawInput(VAV_BeginInput, bIX, beginStrY, 80, value, true)

	// end str label and input
	_, endStrY := nextVerticalCoords(bIRect, 20)
	endStr := "end (mm:ss)"
	endStrW, _ := theCtx.ggCtx.MeasureString(endStr)
	endStrX := dialogOriginX + 40
	theCtx.ggCtx.DrawString(endStr, float64(endStrX), float64(endStrY)+FontSize)
	eIX := dialogOriginX + 40 + int(endStrW) + 20
	value2 := "0:00"
	if IsUpdateDialog {
		value2 = Instructions[ToUpdateInstrNum]["end"]
	}
	eIRect := theCtx.drawInput(VAV_EndInput, eIX, endStrY, 80, value2, true)

	// speedUp checkbox
	theCtx.ggCtx.SetHexColor("#444")
	suL := "speed up video"
	suLW, _ := theCtx.ggCtx.MeasureString(suL)
	sulX := dialogOriginX + 40
	_, sULY := nextVerticalCoords(eIRect, 20)
	theCtx.ggCtx.DrawString(suL, float64(sulX), float64(sULY)+FontSize)
	isSelected := false
	if IsUpdateDialog && strings.ToLower(Instructions[ToUpdateInstrNum]["speedup"]) == "true" {
		isSelected = true
	}
	sUCX := dialogOriginX + 40 + int(suLW) + 30
	sURect := theCtx.drawCheckbox(VAV_SpeedUpCheckbox, sUCX, sULY, isSelected)

	// blackAndWhite checkbox
	theCtx.ggCtx.SetHexColor("#444")
	bwL := "black and white video"
	bwLW, _ := theCtx.ggCtx.MeasureString(bwL)
	bWLX, _ := nextHorizontalCoords(sURect, 40)
	theCtx.ggCtx.DrawString(bwL, float64(bWLX), float64(sURect.OriginY)+FontSize)
	bWCX := bWLX + int(bwLW) + 30
	isSelected2 := false
	if IsUpdateDialog && strings.ToLower(Instructions[ToUpdateInstrNum]["blackwhite"]) == "true" {
		isSelected2 = true
	}
	theCtx.drawCheckbox(VAV_BlackAndWhiteCheckbox, bWCX, sURect.OriginY, isSelected2)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()

}

func vAVMouseCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VAVObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	clearIndicators := func(window *glfw.Window) {
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		beginInputRS := VAVObjCoords[VAV_BeginInput]
		endInputRS := VAVObjCoords[VAV_EndInput]

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(beginInputRS.OriginX)+float64(beginInputRS.Width)+20, float64(beginInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(endInputRS.OriginX)+float64(endInputRS.Width)+20, float64(endInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()
	}

	rootPath, _ := GetRootPath()

	switch widgetCode {
	case VAV_CloseBtn:
		IsUpdateDialog = false
		IsInsertBeforeDialog = false

		drawItemsView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(iVScrollBtnCB)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAV_PickVideo:
		filename := PickVideoFile()
		if filename == "" {
			return
		}
		VAVInputsStore["video"] = filename
		displayFilename := strings.ReplaceAll(filename, rootPath, "")

		theCtx := Continue2dCtx(CurrentWindowFrame, &VAVObjCoords)
		theCtx.drawFileInput(VAV_PickVideo, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// update end str
		ffprobe := GetFFPCommand()
		eIRect := VAVObjCoords[VAV_EndInput]
		videoLength := LengthOfVideo(filename, ffprobe)
		EndInputEnteredTxt = videoLength
		theCtx.drawFileInput(VAV_EndInput, eIRect.OriginX, eIRect.OriginY, eIRect.Width, EndInputEnteredTxt)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAV_BeginInput:
		VAV_SelectedInput = VAV_BeginInput

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

	case VAV_EndInput:
		VAV_SelectedInput = VAV_EndInput
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

	case VAV_SpeedUpCheckbox:
		if VAV_SpeedUpCheckboxSelected {
			VAV_SpeedUpCheckboxSelected = false
		} else {
			VAV_SpeedUpCheckboxSelected = true
		}

		theCtx := Continue2dCtx(CurrentWindowFrame, &VAVObjCoords)
		theCtx.drawCheckbox(VAV_SpeedUpCheckbox, widgetRS.OriginX, widgetRS.OriginY, VAV_SpeedUpCheckboxSelected)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAV_BlackAndWhiteCheckbox:

		if VAV_BlackAndWhiteCheckboxSelected {
			VAV_BlackAndWhiteCheckboxSelected = false
		} else {
			VAV_BlackAndWhiteCheckboxSelected = true
		}

		theCtx := Continue2dCtx(CurrentWindowFrame, &VAVObjCoords)
		theCtx.drawCheckbox(VAV_BlackAndWhiteCheckbox, widgetRS.OriginX, widgetRS.OriginY, VAV_BlackAndWhiteCheckboxSelected)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAV_AddBtn:

		if IsUpdateDialog {
			oldInstr := Instructions[ToUpdateInstrNum]
			if videoPath, ok := oldInstr["video"]; ok {
				oldInstr["video"] = videoPath
			}
			if BeginInputEnteredTxt != "" {
				oldInstr["begin"] = BeginInputEnteredTxt
			}
			if EndInputEnteredTxt != "" {
				oldInstr["end"] = EndInputEnteredTxt
			}
			if oldCBState, ok := oldInstr["speedup"]; ok {
				if oldCBState != strconv.FormatBool(VAV_SpeedUpCheckboxSelected) {
					oldInstr["speedup"] = strconv.FormatBool(VAV_SpeedUpCheckboxSelected)
				}
			}
			if oldBWState, ok := oldInstr["blackwhite"]; ok {
				if oldBWState != strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected) {
					oldInstr["blackwhite"] = strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected)
				}
			}
			Instructions[ToUpdateInstrNum] = oldInstr
			ToUpdateInstrNum = 0
			IsUpdateDialog = false

		} else {
			if VAVInputsStore["video"] == "" {
				return
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":       "video",
					"video":      VAVInputsStore["video"],
					"begin":      BeginInputEnteredTxt,
					"end":        EndInputEnteredTxt,
					"speedup":    strconv.FormatBool(VAV_SpeedUpCheckboxSelected),
					"blackwhite": strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected),
				}
				Instructions = slices.Insert(Instructions, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				Instructions = append(Instructions, map[string]string{
					"kind":       "video",
					"video":      VAVInputsStore["video"],
					"begin":      BeginInputEnteredTxt,
					"end":        EndInputEnteredTxt,
					"speedup":    strconv.FormatBool(VAV_SpeedUpCheckboxSelected),
					"blackwhite": strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected),
				})

			}

		}

		BeginInputEnteredTxt = ""
		EndInputEnteredTxt = ""

		drawItemsView(window, TotalPages())
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}

func vAVKeyCB(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if VAV_SelectedInput == VAV_BeginInput {

		// enforce number types, semicolon and backspace
		if IsKeyNumeric(key) {
			BeginInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			BeginInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(BeginInputEnteredTxt) != 0 {
			BeginInputEnteredTxt = BeginInputEnteredTxt[:len(BeginInputEnteredTxt)-1]
		}

		bIRect := VAVObjCoords[VAV_BeginInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VAVObjCoords)
		theCtx.drawInput(VAV_BeginInput, bIRect.OriginX, bIRect.OriginY, bIRect.Width, BeginInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if VAV_SelectedInput == VAV_EndInput {
		// enforce number types, semicolon and backspace
		if IsKeyNumeric(key) {
			EndInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			EndInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(EndInputEnteredTxt) != 0 {
			EndInputEnteredTxt = EndInputEnteredTxt[:len(EndInputEnteredTxt)-1]
		}

		eIRect := VAVObjCoords[VAV_EndInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VAVObjCoords)
		theCtx.drawInput(VAV_EndInput, eIRect.OriginX, eIRect.OriginY, eIRect.Width, EndInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	}
}
