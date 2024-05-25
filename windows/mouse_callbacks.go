package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/videos349/internal"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sqweek/dialog"
)

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range internal.ObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case internal.AddImgBtn:
		// tmpFrame = internal.CurrentWindowFrame
		internal.DrawViewAddImage(window, internal.CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAddImageMouseCallback)
		window.SetKeyCallback(internal.VaikeyCallback)

	case internal.AddImgSoundBtn:
		internal.DrawViewAIS(window, internal.CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAISMouseCallback)
		window.SetKeyCallback(internal.VaiskeyCallback)

	case internal.AddVidBtn:
		internal.DrawViewAddVideo(window, internal.CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAddVideoMouseCallback)
		window.SetKeyCallback(internal.VavkeyCallback)

	case internal.OpenWDBtn:
		rootPath, _ := internal.GetRootPath()
		internal.ExternalLaunch(rootPath)

	case internal.OurSite:
		if runtime.GOOS == "windows" {
			exec.Command("cmd", "/C", "start", "https://sae.ng").Run()
		} else if runtime.GOOS == "linux" {
			exec.Command("xdg-open", "https://sae.ng").Run()
		}

	case internal.RenderBtn:
		if len(internal.Instructions) == 0 {
			return
		}
		internal.DrawRenderView(window, internal.CurrentWindowFrame)
		window.SetMouseButtonCallback(nil)
		window.SetKeyCallback(nil)
		internal.InChannel <- true
	}

	// for generated buttons
	if widgetCode > 1000 && widgetCode < 2000 {
		instrNum := widgetCode - 1000
		internal.ExternalLaunch(internal.Instructions[instrNum-1]["image"])
	} else if widgetCode > 2000 && widgetCode < 3000 {
		instrNum := widgetCode - 2000
		if _, ok := internal.Instructions[instrNum-1]["audio_optional"]; ok {
			internal.ExternalLaunch(internal.Instructions[instrNum-1]["audio_optional"])
		} else {
			internal.ExternalLaunch(internal.Instructions[instrNum-1]["audio"])
		}
	} else if widgetCode > 3000 {
		instrNum := widgetCode - 3000
		internal.ExternalLaunch(internal.Instructions[instrNum-1]["video"])
	}

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

	clearIndicators := func(window *glfw.Window) {
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)

		durationInputRS := internal.VaiObjCoords[internal.VAI_DurInput]

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(durationInputRS.OriginX)+float64(durationInputRS.Width)+20, float64(durationInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case internal.VAI_CloseBtn:
		internal.AllDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAI_SelectImg:
		filename, err := dialog.File().Filter("PNG Image", "png").Filter("JPEG Image", "jpg").Load()
		if filename == "" || err != nil {
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

	case internal.VAI_DurInput:
		internal.VAI_SelectedInput = internal.VAI_DurInput

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

		internal.Instructions = append(internal.Instructions, map[string]string{
			"kind":     "image",
			"image":    internal.VaiInputsStore["image"],
			"duration": internal.VaiInputsStore["duration"],
		})

		internal.AllDraws(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
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
		internal.AllDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAIS_SelectImg:
		filename, err := dialog.File().Filter("PNG Image", "png").Filter("JPEG Image", "jpg").Load()
		if filename == "" || err != nil {
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
		filename, err := dialog.File().Filter("MP3 Audio", "mp3").Filter("FLAC Audio", "flac").
			Filter("WAV Audio", "wav").Load()
		if filename == "" || err != nil {
			return
		}
		internal.VaisInputsStore["audio"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)
		// load font
		fontPath := internal.GetDefaultFontPath()
		err = ggCtx.LoadFontFace(fontPath, 20)
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

		internal.Instructions = append(internal.Instructions, map[string]string{
			"kind":        "image",
			"image":       internal.VaisInputsStore["image"],
			"audio":       internal.VaisInputsStore["audio"],
			"audio_begin": internal.VaisInputsStore["audio_begin"],
			"audio_end":   internal.VaisInputsStore["audio_end"],
		})

		internal.AllDraws(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
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
		audioBeginInputRS := internal.VavObjCoords[internal.VAV_AudioBegin]

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(beginInputRS.OriginX)+float64(beginInputRS.Width)+20, float64(beginInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(endInputRS.OriginX)+float64(endInputRS.Width)+20, float64(endInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawCircle(float64(audioBeginInputRS.OriginX)+float64(audioBeginInputRS.Width)+20,
			float64(audioBeginInputRS.OriginY)+15, 20)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		internal.CurrentWindowFrame = ggCtx.Image()
	}

	switch widgetCode {
	case internal.VAV_CloseBtn:
		internal.AllDraws(window)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	case internal.VAV_PickVideo:
		filename, err := dialog.File().Filter("MP4 Video", "mp4").Filter("WEBM Video", "webm").Filter("MKV Video", "mkv").Load()
		if filename == "" || err != nil {
			return
		}
		internal.VavInputsStore["video"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)
		// load font
		fontPath := internal.GetDefaultFontPath()
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

	case internal.VAV_AudioBegin:
		internal.VAV_SelectedInput = internal.VAV_AudioBegin
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

	case internal.VAV_PickAudio:
		filename, err := dialog.File().Filter("MP3 Audio", "mp3").Filter("FLAC Audio", "flac").
			Filter("WAV Audio", "wav").Load()
		if filename == "" || err != nil {
			return
		}
		internal.VavInputsStore["audio_optional"] = filename

		// write audio name
		ggCtx := gg.NewContextForImage(internal.CurrentWindowFrame)
		// load font
		fontPath := internal.GetDefaultFontPath()
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
		internal.CurrentWindowFrame = ggCtx.Image()

	case internal.VAV_AddBtn:
		if internal.VavInputsStore["video"] == "" {
			return
		}

		internal.Instructions = append(internal.Instructions, map[string]string{
			"kind":                 "video",
			"video":                internal.VavInputsStore["video"],
			"begin":                internal.BeginInputEnteredTxt,
			"end":                  internal.EndInputEnteredTxt,
			"audio_optional":       internal.VavInputsStore["audio_optional"],
			"audio_begin_optional": internal.VAV_AudioBeginEnteredTxt,
		})

		internal.AllDraws(window)

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(mouseBtnCallback)
		// unregister the keyCallback
		window.SetKeyCallback(nil)

	}

}
