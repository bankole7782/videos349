package main

import (
	"slices"
	"strconv"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func viewAddImageMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VaiObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	// rootPath, _ := GetRootPath()

	switch widgetCode {
	case VAI_CloseBtn:
		IsUpdateDialog = false
		IsInsertBeforeDialog = false

		DrawWorkView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAI_SelectImg:
		filename := PickImageFile()
		if filename == "" {
			return
		}
		VaiInputsStore["image"] = filename
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")

		theCtx := Continue2dCtx(CurrentWindowFrame, &VaiObjCoords)
		theCtx.drawFileInput(VAI_SelectImg, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAI_AddBtn:

		if IsUpdateDialog {
			oldInstr := Instructions[ToUpdateInstrNum]
			if filename, ok := VaiInputsStore["image"]; ok {
				oldInstr["image"] = filename
			}
			if VAI_DurationEnteredTxt != "" {
				oldInstr["duration"] = VAI_DurationEnteredTxt
				VAI_DurationEnteredTxt = ""
			}

			Instructions[ToUpdateInstrNum] = oldInstr
			IsUpdateDialog = false

		} else {

			if VaiInputsStore["image"] == "" {
				return
			}

			if VAI_DurationEnteredTxt == "" {
				VaiInputsStore["duration"] = "5"
			} else {
				VaiInputsStore["duration"] = VAI_DurationEnteredTxt
				VAI_DurationEnteredTxt = ""
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":     "image",
					"image":    VaiInputsStore["image"],
					"duration": VaiInputsStore["duration"],
				}
				Instructions = slices.Insert(Instructions, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				Instructions = append(Instructions, map[string]string{
					"kind":     "image",
					"image":    VaiInputsStore["image"],
					"duration": VaiInputsStore["duration"],
				})

			}

		}

		DrawWorkView(window, TotalPages())

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	}

}

func viewAISMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VaisObjCoords {
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

		aBInputRS := VaisObjCoords[VAIS_AudioBeginInput]
		aEInputRS := VaisObjCoords[VAIS_AudioEndInput]

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

		DrawWorkView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAIS_SelectImg:
		filename := PickImageFile()
		if filename == "" {
			return
		}
		VaisInputsStore["image"] = filename
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		theCtx := Continue2dCtx(CurrentWindowFrame, &VaisObjCoords)
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
		VaisInputsStore["audio"] = filename

		theCtx := Continue2dCtx(CurrentWindowFrame, &VaisObjCoords)
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		theCtx.drawFileInput(VAIS_SelectAudio, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// update end str
		ffprobe := GetFFPCommand()
		eIRect := VaisObjCoords[VAIS_AudioEndInput]
		videoLength := LengthOfVideo(filename, ffprobe)
		VaisEndInputEnteredTxt = videoLength
		theCtx.drawInput(VAIS_AudioEndInput, eIRect.OriginX, eIRect.OriginY, eIRect.Width, VaisEndInputEnteredTxt, true)

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
			if VaisBeginInputEnteredTxt != "" {
				oldInstr["audio_begin"] = VaisBeginInputEnteredTxt
				VaisBeginInputEnteredTxt = ""
			}

			if VaisEndInputEnteredTxt != "" {
				oldInstr["audio_end"] = VaisEndInputEnteredTxt
				VaisEndInputEnteredTxt = ""
			}
			Instructions[ToUpdateInstrNum] = oldInstr
			ToUpdateInstrNum = 0
			IsUpdateDialog = false

		} else {
			if VaisInputsStore["image"] == "" {
				return
			}

			if VaisInputsStore["audio"] == "" {
				return
			}

			if VaisBeginInputEnteredTxt == "" {
				VaisInputsStore["audio_begin"] = "0:00"
			} else {
				VaisInputsStore["audio_begin"] = VaisBeginInputEnteredTxt
				VaisBeginInputEnteredTxt = ""
			}

			if VaisEndInputEnteredTxt == "" {
				VaisInputsStore["audio_end"] = "5"
			} else {
				VaisInputsStore["audio_end"] = VaisEndInputEnteredTxt
				VaisEndInputEnteredTxt = ""
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":        "image",
					"image":       VaisInputsStore["image"],
					"audio":       VaisInputsStore["audio"],
					"audio_begin": VaisInputsStore["audio_begin"],
					"audio_end":   VaisInputsStore["audio_end"],
				}
				Instructions = slices.Insert(Instructions, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				Instructions = append(Instructions, map[string]string{
					"kind":        "image",
					"image":       VaisInputsStore["image"],
					"audio":       VaisInputsStore["audio"],
					"audio_begin": VaisInputsStore["audio_begin"],
					"audio_end":   VaisInputsStore["audio_end"],
				})

			}

		}

		DrawWorkView(window, TotalPages())
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}

func viewAddVideoMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VavObjCoords {
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

		beginInputRS := VavObjCoords[VAV_BeginInput]
		endInputRS := VavObjCoords[VAV_EndInput]

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

		DrawWorkView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAV_PickVideo:
		filename := PickVideoFile()
		if filename == "" {
			return
		}
		VavInputsStore["video"] = filename
		displayFilename := strings.ReplaceAll(filename, rootPath, "")

		theCtx := Continue2dCtx(CurrentWindowFrame, &VavObjCoords)
		theCtx.drawFileInput(VAV_PickVideo, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// update end str
		ffprobe := GetFFPCommand()
		eIRect := VavObjCoords[VAV_EndInput]
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

		theCtx := Continue2dCtx(CurrentWindowFrame, &VavObjCoords)
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

		theCtx := Continue2dCtx(CurrentWindowFrame, &VavObjCoords)
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
			if VavInputsStore["video"] == "" {
				return
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":       "video",
					"video":      VavInputsStore["video"],
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
					"video":      VavInputsStore["video"],
					"begin":      BeginInputEnteredTxt,
					"end":        EndInputEnteredTxt,
					"speedup":    strconv.FormatBool(VAV_SpeedUpCheckboxSelected),
					"blackwhite": strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected),
				})

			}

		}

		BeginInputEnteredTxt = ""
		EndInputEnteredTxt = ""

		DrawWorkView(window, TotalPages())
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}
