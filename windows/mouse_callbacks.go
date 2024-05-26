package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"

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
		internal.ProjectName = internal.NameInputEnteredTxt + ".v3p"
		outPath := filepath.Join(rootPath, internal.ProjectName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		internal.DrawWorkView(window)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
	}

	if widgetCode > 1000 && widgetCode < 2000 {
		num := widgetCode - 1000 - 1
		projectFile := internal.GetProjectFiles()[num]

		internal.ProjectName = projectFile.Name

		// load instructions
		obj := make([]map[string]string, 0)
		rootPath, _ := internal.GetRootPath()
		inPath := filepath.Join(rootPath, internal.ProjectName)
		rawBytes, _ := os.ReadFile(inPath)
		json.Unmarshal(rawBytes, &obj)

		internal.Instructions = append(internal.Instructions, obj...)

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
		instrNum := widgetCode - 1000 - 1
		internal.ExternalLaunch(internal.Instructions[instrNum]["image"])
	} else if widgetCode > 2000 && widgetCode < 3000 {
		instrNum := widgetCode - 2000 - 1
		internal.ExternalLaunch(internal.Instructions[instrNum]["audio"])
	} else if widgetCode > 3000 && widgetCode < 4000 {
		instrNum := widgetCode - 3000 - 1
		internal.ExternalLaunch(internal.Instructions[instrNum]["video"])
	} else if widgetCode > 4000 && widgetCode < 5000 {
		// bring up update instruction dialog
		instrNum := widgetCode - 4000 - 1
		internal.ToUpdateInstrNum = instrNum
		if internal.Instructions[instrNum]["image"] != "" {
			internal.DrawViewAddImage(window, internal.CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAddImageMouseCallback)
			window.SetKeyCallback(internal.VaikeyCallback)
		} else if internal.Instructions[instrNum]["audio"] != "" {
			internal.DrawViewAIS(window, internal.CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAISMouseCallback)
			window.SetKeyCallback(internal.VaiskeyCallback)
		} else if internal.Instructions[instrNum]["video"] != "" {
			internal.DrawViewAddVideo(window, internal.CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAddVideoMouseCallback)
			window.SetKeyCallback(internal.VavkeyCallback)
		}
	} else if widgetCode > 5000 {
		// delete from instructions slice
		instrNum := widgetCode - 5000 - 1
		internal.Instructions = slices.Delete(internal.Instructions, instrNum, instrNum+1)

		internal.ObjCoords = make(map[int]g143.RectSpecs)
		internal.DrawWorkView(window)
	}

}
