package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/videos349/internal"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func projViewMouseCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range internal.ProjObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := internal.GetRootPath()

	switch widgetCode {
	case internal.PROJ_NewProject:
		if internal.NameInputEnteredTxt == "" {
			return
		}

		// create file
		internal.ProjectName = internal.NameInputEnteredTxt
		outPath := filepath.Join(rootPath, internal.ProjectName+".v3p")
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		internal.DrawWorkView(window)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
	}
}

func workViewMouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
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
