package main

import (
	"os"
	"path/filepath"

	g143 "github.com/bankole7782/graphics143"
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
	} else if key == glfw.KeyEnter && len(NameInputEnteredTxt) != 0 {
		// create file
		rootPath, _ := GetRootPath()

		ProjectName = NameInputEnteredTxt + ".v3p"
		outPath := filepath.Join(rootPath, ProjectName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		DrawWorkView(window, 1)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		return
	} else {
		NameInputEnteredTxt += glfw.GetKeyName(key, scancode)
	}

	nIRS := ProjObjCoords[PROJ_NameInput]
	theCtx := Continue2dCtx(CurrentWindowFrame, &ProjObjCoords)
	theCtx.drawInput(PROJ_NameInput, nIRS.OriginX, nIRS.OriginY, nIRS.Width, NameInputEnteredTxt, true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
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

	dIRS := VaiObjCoords[VAI_DurInput]
	theCtx := Continue2dCtx(CurrentWindowFrame, &VaiObjCoords)
	theCtx.drawInput(VAI_DurInput, dIRS.OriginX, dIRS.OriginY, dIRS.Width, VAI_DurationEnteredTxt, true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
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

		aBRect := VaisObjCoords[VAIS_AudioBeginInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VaisObjCoords)
		theCtx.drawInput(VAIS_AudioBeginInput, aBRect.OriginX, aBRect.OriginY, aBRect.Width, VaisBeginInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	} else if VAIS_SelectedInput == VAIS_AudioEndInput {
		// enforce number types
		if IsKeyNumeric(key) {
			VaisEndInputEnteredTxt += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeySemicolon {
			VaisEndInputEnteredTxt += ":"
		} else if key == glfw.KeyBackspace && len(VaisEndInputEnteredTxt) != 0 {
			VaisEndInputEnteredTxt = VaisEndInputEnteredTxt[:len(VaisEndInputEnteredTxt)-1]
		}

		aEIRect := VaisObjCoords[VAIS_AudioEndInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VaisObjCoords)
		theCtx.drawInput(VAIS_AudioEndInput, aEIRect.OriginX, aEIRect.OriginY, aEIRect.Width, VaisEndInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()
	}

}

func VavkeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
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

		bIRect := VavObjCoords[VAV_BeginInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VavObjCoords)
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

		eIRect := VavObjCoords[VAV_EndInput]
		theCtx := Continue2dCtx(CurrentWindowFrame, &VavObjCoords)
		theCtx.drawInput(VAV_EndInput, eIRect.OriginX, eIRect.OriginY, eIRect.Width, EndInputEnteredTxt, true)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	}
}
