package internal

import (
	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func VaikeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	if VAI_SelectedInput == VAI_DurInput {
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

	} else if VAI_SelectedInput == VAI_AudioBegin {
		// enforce number types
		if IsKeyNumeric(key) {
			VAI_AudioBeginEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VAI_AudioBeginEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VAI_AudioBeginEnteredTxt) != 0 {
			VAI_AudioBeginEnteredTxt = VAI_AudioBeginEnteredTxt[:len(VAI_AudioBeginEnteredTxt)-1]
		}

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		// load font
		fontPath := GetDefaultFontPath()
		err := ggCtx.LoadFontFace(fontPath, 20)
		if err != nil {
			panic(err)
		}

		audioBeginInputRS := VaiObjCoords[VAI_AudioBegin]

		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(audioBeginInputRS.OriginX), float64(audioBeginInputRS.OriginY),
			float64(audioBeginInputRS.Width), float64(audioBeginInputRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(VAI_AudioBeginEnteredTxt, float64(audioBeginInputRS.OriginX+10), float64(audioBeginInputRS.OriginY)+20)

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

	}
}
