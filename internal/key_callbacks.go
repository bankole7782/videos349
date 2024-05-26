package internal

import (
	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func ProjKeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if key == glfw.KeyBackspace && len(NameInputEnteredTxt) != 0 {
		NameInputEnteredTxt = NameInputEnteredTxt[:len(NameInputEnteredTxt)-1]
	} else if key == glfw.KeySpace {
		NameInputEnteredTxt += " "
	} else {
		NameInputEnteredTxt += glfw.GetKeyName(key, scancode)
	}

	ggCtx := gg.NewContextForImage(CurrentWindowFrame)
	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	nameInputRS := ProjObjCoords[PROJ_NameInput]

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(nameInputRS.OriginX+3), float64(nameInputRS.OriginY+3),
		float64(nameInputRS.Width)-6, float64(nameInputRS.Height)-6)
	ggCtx.Fill()

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(NameInputEnteredTxt, float64(nameInputRS.OriginX+25), float64(nameInputRS.OriginY+25))

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}

func VaikeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	// enforce number types
	if IsKeyNumeric(key) {
		VAI_DurationEnteredTxt += glfw.GetKeyName(key, scancode)
	} else if key == glfw.KeyBackspace && len(VAI_DurationEnteredTxt) != 0 {
		VAI_DurationEnteredTxt = VAI_DurationEnteredTxt[:len(VAI_DurationEnteredTxt)-1]
	}

	ggCtx := gg.NewContextForImage(CurrentWindowFrame)
	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	durInputRS := VaiObjCoords[VAI_DurInput]

	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRoundedRectangle(float64(durInputRS.OriginX), float64(durInputRS.OriginY), float64(durInputRS.Width),
		float64(durInputRS.Height), 10)
	ggCtx.Fill()

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(VAI_DurationEnteredTxt, float64(durInputRS.OriginX+10), float64(durInputRS.OriginY)+20)

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}

func VaiskeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if VAIS_SelectedInput == VAIS_AudioBeginInput {
		// enforce number types
		if IsKeyNumeric(key) {
			VaisBeginInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VaisBeginInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VaisBeginInputEnteredTxt) != 0 {
			VaisBeginInputEnteredTxt = VaisBeginInputEnteredTxt[:len(VaisBeginInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		aBInputRS := VaisObjCoords[VAIS_AudioBeginInput]

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(aBInputRS.OriginX), float64(aBInputRS.OriginY), float64(aBInputRS.Width),
			float64(aBInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(VaisBeginInputEnteredTxt, float64(aBInputRS.OriginX+10), float64(aBInputRS.OriginY)+20)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	} else if VAIS_SelectedInput == VAIS_AudioEndInput {
		// enforce number types
		if IsKeyNumeric(key) {
			VaisEndInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VaisEndInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VaisEndInputEnteredTxt) != 0 {
			VaisEndInputEnteredTxt = VaisEndInputEnteredTxt[:len(VaisEndInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		aEINputRS := VaisObjCoords[VAIS_AudioEndInput]

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(aEINputRS.OriginX), float64(aEINputRS.OriginY),
			float64(aEINputRS.Width), float64(aEINputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(VaisEndInputEnteredTxt, float64(aEINputRS.OriginX+10), float64(aEINputRS.OriginY)+20)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()
	}

}

func VavkeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if VAV_SelectedInput == VAV_BeginInput {
		beginInputRS := VavObjCoords[VAV_BeginInput]

		// enforce number types, semicolon and backspace
		if IsKeyNumeric(key) {
			BeginInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			BeginInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(BeginInputEnteredTxt) != 0 {
			BeginInputEnteredTxt = BeginInputEnteredTxt[:len(BeginInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(beginInputRS.OriginX), float64(beginInputRS.OriginY), float64(beginInputRS.Width),
			float64(beginInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(BeginInputEnteredTxt, float64(beginInputRS.OriginX+10), float64(beginInputRS.OriginY)+FontSize)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	} else if VAV_SelectedInput == VAV_EndInput {
		endInputRS := VavObjCoords[VAV_EndInput]

		// enforce number types, semicolon and backspace
		if IsKeyNumeric(key) {
			EndInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			EndInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(EndInputEnteredTxt) != 0 {
			EndInputEnteredTxt = EndInputEnteredTxt[:len(EndInputEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(endInputRS.OriginX), float64(endInputRS.OriginY), float64(endInputRS.Width),
			float64(endInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(EndInputEnteredTxt, float64(endInputRS.OriginX+10), float64(endInputRS.OriginY)+FontSize)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	} else if VAV_SelectedInput == VAV_AudioBegin {
		audioBeginInputRS := VavObjCoords[VAV_AudioBegin]

		// enforce number types, semicolon and backspace
		if IsKeyNumeric(key) {
			VAV_AudioBeginEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VAV_AudioBeginEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VAV_AudioBeginEnteredTxt) != 0 {
			VAV_AudioBeginEnteredTxt = VAV_AudioBeginEnteredTxt[:len(VAV_AudioBeginEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(audioBeginInputRS.OriginX), float64(audioBeginInputRS.OriginY),
			float64(audioBeginInputRS.Width), float64(audioBeginInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(VAV_AudioBeginEnteredTxt, float64(audioBeginInputRS.OriginX+10), float64(audioBeginInputRS.OriginY)+FontSize)

		// send the frame to glfw window
		windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = ggCtx.Image()

	}
}
