package main

import (
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/videos349/internal"
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

	for code, RS := range internal.VaiObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	// rootPath, _ := internal.GetRootPath()

	switch widgetCode {
	case internal.VAI_CloseBtn:
		internal.DrawWorkView(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAI_SelectImg:
		filename := pickFileUbuntu("png|jpg")
		if filename == "" {
			return
		}
		internal.VaiInputsStore["image"] = filename

		// show picked image
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

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
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAI_AddBtn:
		if internal.VaiInputsStore["image"] == "" {
			return
		}

		if internal.VAI_DurationEnteredTxt == "" {
			internal.VaiInputsStore["duration"] = "5"
		} else {
			internal.VaiInputsStore["duration"] = internal.VAI_DurationEnteredTxt
			internal.VAI_DurationEnteredTxt = ""
		}

		if internal.ToUpdateInstrNum != 0 {
			internal.Instructions[internal.ToUpdateInstrNum] = map[string]string{
				"kind":     "image",
				"image":    internal.VaiInputsStore["image"],
				"duration": internal.VaiInputsStore["duration"],
			}

			internal.ToUpdateInstrNum = 0
		} else {
			internal.Instructions = append(internal.Instructions, map[string]string{
				"kind":     "image",
				"image":    internal.VaiInputsStore["image"],
				"duration": internal.VaiInputsStore["duration"],
			})

		}

		internal.DrawWorkView(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

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

	for code, RS := range internal.VaisObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := internal.GetRootPath()

	clearIndicators := func(window *glfw.Window) {
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		aBInputRS := internal.VaisObjCoords[internal.VAIS_AudioBeginInput]
		aEInputRS := internal.VaisObjCoords[internal.VAIS_AudioEndInput]

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
		internal.CurrentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case internal.VAIS_CloseBtn:
		internal.DrawWorkView(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAIS_SelectImg:
		filename := pickFileUbuntu("png|jpg")
		if filename == "" {
			return
		}
		internal.VaisInputsStore["image"] = filename

		// show picked image
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

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
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAIS_SelectAudio:
		filename := pickFileUbuntu("mp3|flac|wav")
		if filename == "" {
			return
		}
		internal.VaisInputsStore["audio"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)
		// load font
		fontPath := internal.GetDefaultFontPath()
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
		endInputRS := internal.VaisObjCoords[internal.VAIS_AudioEndInput]
		videoLength := internal.LengthOfVideo(filename, ffprobe)
		internal.VaisEndInputEnteredTxt = videoLength

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY),
			float64(endInputRS.Width), float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(videoLength, float64(endInputRS.OriginX)+10, float64(endInputRS.OriginY+internal.FontSize))

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAIS_AudioBeginInput:
		internal.VAIS_SelectedInput = internal.VAIS_AudioBeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAIS_AudioEndInput:
		internal.VAIS_SelectedInput = internal.VAIS_AudioEndInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAIS_AddBtn:
		if internal.VaisInputsStore["image"] == "" {
			return
		}

		if internal.VaisInputsStore["audio"] == "" {
			return
		}

		if internal.VaisBeginInputEnteredTxt == "" {
			internal.VaisInputsStore["audio_begin"] = "5"
		} else {
			internal.VaisInputsStore["audio_begin"] = internal.VaisBeginInputEnteredTxt
			internal.VaisBeginInputEnteredTxt = ""
		}

		if internal.VaisEndInputEnteredTxt == "" {
			internal.VaisInputsStore["audio_end"] = "5"
		} else {
			internal.VaisInputsStore["audio_end"] = internal.VaisEndInputEnteredTxt
			internal.VaisEndInputEnteredTxt = ""
		}

		if internal.ToUpdateInstrNum != 0 {
			internal.Instructions[internal.ToUpdateInstrNum] = map[string]string{
				"kind":        "image",
				"image":       internal.VaisInputsStore["image"],
				"audio":       internal.VaisInputsStore["audio"],
				"audio_begin": internal.VaisInputsStore["audio_begin"],
				"audio_end":   internal.VaisInputsStore["audio_end"],
			}
			internal.ToUpdateInstrNum = 0
		} else {
			internal.Instructions = append(internal.Instructions, map[string]string{
				"kind":        "image",
				"image":       internal.VaisInputsStore["image"],
				"audio":       internal.VaisInputsStore["audio"],
				"audio_begin": internal.VaisInputsStore["audio_begin"],
				"audio_end":   internal.VaisInputsStore["audio_end"],
			})

		}

		internal.DrawWorkView(window)

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

	for code, RS := range internal.VavObjCoords {
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
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		beginInputRS := internal.VavObjCoords[internal.VAV_BeginInput]
		endInputRS := internal.VavObjCoords[internal.VAV_EndInput]

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
		internal.CurrentWindowFrame = ggCtx.Image()
	}

	rootPath, _ := internal.GetRootPath()

	switch widgetCode {
	case internal.VAV_CloseBtn:
		internal.DrawWorkView(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAV_PickVideo:
		filename := pickFileUbuntu("mp4|mkv|webm")
		if filename == "" {
			return
		}
		internal.VavInputsStore["video"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)
		// load font
		fontPath := internal.GetDefaultFontPath()
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
		endInputRS := internal.VavObjCoords[internal.VAV_EndInput]
		videoLength := internal.LengthOfVideo(filename, ffprobe)
		internal.EndInputEnteredTxt = videoLength

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY),
			float64(endInputRS.Width), float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(videoLength, float64(endInputRS.OriginX)+10, float64(endInputRS.OriginY+internal.FontSize))

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAV_BeginInput:
		internal.VAV_SelectedInput = internal.VAV_BeginInput

		clearIndicators(window)

		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAV_EndInput:
		internal.VAV_SelectedInput = internal.VAV_EndInput
		clearIndicators(window)

		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawCircle(float64(widgetRS.OriginX)+float64(widgetRS.Width)+20, float64(widgetRS.OriginY)+15, 10)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAV_AddBtn:
		if internal.VavInputsStore["video"] == "" {
			return
		}

		if internal.ToUpdateInstrNum != 0 {
			internal.Instructions[internal.ToUpdateInstrNum] = map[string]string{
				"kind":  "video",
				"video": internal.VavInputsStore["video"],
				"begin": internal.BeginInputEnteredTxt,
				"end":   internal.EndInputEnteredTxt,
			}
			internal.ToUpdateInstrNum = 0
		} else {
			internal.Instructions = append(internal.Instructions, map[string]string{
				"kind":  "video",
				"video": internal.VavInputsStore["video"],
				"begin": internal.BeginInputEnteredTxt,
				"end":   internal.EndInputEnteredTxt,
			})
		}

		internal.DrawWorkView(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}
