package main

import (
	"strconv"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
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

	var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range VaiObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
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
		DrawWorkView(window, TotalPages())
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

		// show picked image
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

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
		CurrentWindowFrame = ggCtx.Image()

	case VAI_AddBtn:
		if VaiInputsStore["image"] == "" {
			return
		}

		if VAI_DurationEnteredTxt == "" {
			VaiInputsStore["duration"] = "5"
		} else {
			VaiInputsStore["duration"] = VAI_DurationEnteredTxt
			VAI_DurationEnteredTxt = ""
		}

		if ToUpdateInstrNum != 0 {
			Instructions[ToUpdateInstrNum] = map[string]string{
				"kind":     "image",
				"image":    VaiInputsStore["image"],
				"duration": VaiInputsStore["duration"],
			}

			ToUpdateInstrNum = 0
		} else {
			Instructions = append(Instructions, map[string]string{
				"kind":     "image",
				"image":    VaiInputsStore["image"],
				"duration": VaiInputsStore["duration"],
			})

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

	var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range VaisObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
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
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case VAIS_CloseBtn:
		DrawWorkView(window, TotalPages())
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

		// show picked image
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

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
		CurrentWindowFrame = ggCtx.Image()

	case VAIS_SelectAudio:
		filename := PickAudioFile()
		if filename == "" {
			return
		}
		VaisInputsStore["audio"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
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

		// update end str
		ffprobe := GetFFPCommand()
		endInputRS := VaisObjCoords[VAIS_AudioEndInput]
		videoLength := LengthOfVideo(filename, ffprobe)
		VaisEndInputEnteredTxt = videoLength

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY),
			float64(endInputRS.Width), float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(videoLength, float64(endInputRS.OriginX)+10, float64(endInputRS.OriginY+FontSize))

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAIS_AudioBeginInput:
		VAIS_SelectedInput = VAIS_AudioBeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
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
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAIS_AddBtn:
		if VaisInputsStore["image"] == "" {
			return
		}

		if VaisInputsStore["audio"] == "" {
			return
		}

		if VaisBeginInputEnteredTxt == "" {
			VaisInputsStore["audio_begin"] = "5"
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

		if ToUpdateInstrNum != 0 {
			Instructions[ToUpdateInstrNum] = map[string]string{
				"kind":        "image",
				"image":       VaisInputsStore["image"],
				"audio":       VaisInputsStore["audio"],
				"audio_begin": VaisInputsStore["audio_begin"],
				"audio_end":   VaisInputsStore["audio_end"],
			}
			ToUpdateInstrNum = 0
		} else {
			Instructions = append(Instructions, map[string]string{
				"kind":        "image",
				"image":       VaisInputsStore["image"],
				"audio":       VaisInputsStore["audio"],
				"audio_begin": VaisInputsStore["audio_begin"],
				"audio_end":   VaisInputsStore["audio_end"],
			})

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

	var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range VavObjCoords {
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
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()
	}

	rootPath, _ := GetRootPath()

	switch widgetCode {
	case VAV_CloseBtn:
		DrawWorkView(window, TotalPages())
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

		// write audio name
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
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
		ggCtx.DrawString(displayFilename, float64(widgetRS.OriginX+10), float64(widgetRS.OriginY+20))

		// update end str
		ffprobe := GetFFPCommand()
		endInputRS := VavObjCoords[VAV_EndInput]
		videoLength := LengthOfVideo(filename, ffprobe)
		EndInputEnteredTxt = videoLength

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY),
			float64(endInputRS.Width), float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(videoLength, float64(endInputRS.OriginX)+10, float64(endInputRS.OriginY+FontSize))

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAV_BeginInput:
		VAV_SelectedInput = VAV_BeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
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
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAV_SpeedUpCheckbox:
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		suRS := VavObjCoords[VAV_SpeedUpCheckbox]
		if VAV_SpeedUpCheckboxSelected {
			ggCtx.SetHexColor("#fff")
			ggCtx.DrawRectangle(float64(suRS.OriginX)+2, float64(suRS.OriginY)+2, float64(suRS.Width)-4,
				float64(suRS.Height)-4)
			ggCtx.Fill()

			VAV_SpeedUpCheckboxSelected = false
		} else {

			ggCtx.SetHexColor("#444")
			ggCtx.DrawRectangle(float64(suRS.OriginX)+4, float64(suRS.OriginY)+4, float64(suRS.Width)-8,
				float64(suRS.Height)-8)
			ggCtx.Fill()

			VAV_SpeedUpCheckboxSelected = true
		}

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAV_BlackAndWhiteCheckbox:
		ggCtx := gg.NewContextForImage(CurrentWindowFrame)

		bwRS := VavObjCoords[VAV_BlackAndWhiteCheckbox]
		if VAV_BlackAndWhiteCheckboxSelected {
			ggCtx.SetHexColor("#fff")
			ggCtx.DrawRectangle(float64(bwRS.OriginX)+2, float64(bwRS.OriginY)+2, float64(bwRS.Width)-4,
				float64(bwRS.Height)-4)
			ggCtx.Fill()

			VAV_BlackAndWhiteCheckboxSelected = false
		} else {

			ggCtx.SetHexColor("#444")
			ggCtx.DrawRectangle(float64(bwRS.OriginX)+4, float64(bwRS.OriginY)+4, float64(bwRS.Width)-8,
				float64(bwRS.Height)-8)
			ggCtx.Fill()

			VAV_BlackAndWhiteCheckboxSelected = true
		}

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	case VAV_AddBtn:
		if VavInputsStore["video"] == "" {
			return
		}

		if ToUpdateInstrNum != 0 {
			Instructions[ToUpdateInstrNum] = map[string]string{
				"kind":       "video",
				"video":      VavInputsStore["video"],
				"begin":      BeginInputEnteredTxt,
				"end":        EndInputEnteredTxt,
				"speedup":    strconv.FormatBool(VAV_SpeedUpCheckboxSelected),
				"blackwhite": strconv.FormatBool(VAV_BlackAndWhiteCheckboxSelected),
			}
			ToUpdateInstrNum = 0
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

		DrawWorkView(window, TotalPages())
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}
