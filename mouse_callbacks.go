package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"

	g143 "github.com/bankole7782/graphics143"
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

	for code, RS := range ProjObjCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	rootPath, _ := GetRootPath()

	switch widgetCode {
	case PROJ_NewProject:
		if NameInputEnteredTxt == "" {
			return
		}

		// create file
		ProjectName = NameInputEnteredTxt + ".v3p"
		outPath := filepath.Join(rootPath, ProjectName)
		os.WriteFile(outPath, []byte(""), 0777)

		// move to work view
		DrawWorkView(window, 1)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		// quick hover effect
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	}

	if widgetCode > 1000 && widgetCode < 2000 {
		num := widgetCode - 1000 - 1
		projectFile := GetProjectFiles()[num]

		ProjectName = projectFile.Name

		// load instructions
		obj := make([]map[string]string, 0)
		rootPath, _ := GetRootPath()
		inPath := filepath.Join(rootPath, ProjectName)
		rawBytes, _ := os.ReadFile(inPath)
		json.Unmarshal(rawBytes, &obj)

		Instructions = append(Instructions, obj...)

		// move to work view
		DrawWorkView(window, 1)
		window.SetMouseButtonCallback(workViewMouseBtnCallback)
		window.SetKeyCallback(nil)
		window.SetScrollCallback(FirstUIScrollCallback)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
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

	for code, RS := range ObjCoords {
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
	case AddImgBtn:
		// tmpFrame = CurrentWindowFrame
		DrawViewAddImage(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAddImageMouseCallback)
		window.SetKeyCallback(VaikeyCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VaiObjCoords))

	case AddImgSoundBtn:
		DrawViewAIS(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAISMouseCallback)
		window.SetKeyCallback(VaiskeyCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VaisObjCoords))

	case AddVidBtn:
		DrawViewAddVideo(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(viewAddVideoMouseCallback)
		window.SetKeyCallback(VavkeyCallback)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VavObjCoords))

	case OpenWDBtn:
		rootPath, _ := GetRootPath()
		ExternalLaunch(rootPath)

	case RenderBtn:
		if len(Instructions) == 0 {
			return
		}
		DrawRenderView(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(nil)
		window.SetKeyCallback(nil)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(nil)
		InChannel <- true
	}

	// for generated buttons
	if widgetCode > 1000 && widgetCode < 2000 {
		instrNum := widgetCode - 1000 - 1
		ExternalLaunch(Instructions[instrNum]["image"])
	} else if widgetCode > 2000 && widgetCode < 3000 {
		instrNum := widgetCode - 2000 - 1
		ExternalLaunch(Instructions[instrNum]["audio"])
	} else if widgetCode > 3000 && widgetCode < 4000 {
		instrNum := widgetCode - 3000 - 1
		ExternalLaunch(Instructions[instrNum]["video"])
	} else if widgetCode > 4000 && widgetCode < 5000 {
		// bring up update instruction dialog
		instrNum := widgetCode - 4000 - 1
		ToUpdateInstrNum = instrNum
		if Instructions[instrNum]["image"] != "" {
			DrawViewAddImage(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAddImageMouseCallback)
			window.SetKeyCallback(VaikeyCallback)
			window.SetCursorPosCallback(getHoverCB(VaiObjCoords))

		} else if Instructions[instrNum]["audio"] != "" {
			DrawViewAIS(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAISMouseCallback)
			window.SetKeyCallback(VaiskeyCallback)
			window.SetCursorPosCallback(getHoverCB(VaisObjCoords))

		} else if Instructions[instrNum]["video"] != "" {
			DrawViewAddVideo(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(viewAddVideoMouseCallback)
			window.SetKeyCallback(VavkeyCallback)
			window.SetCursorPosCallback(getHoverCB(VavObjCoords))

		}
	} else if widgetCode > 5000 {
		// delete from instructions slice
		instrNum := widgetCode - 5000 - 1
		Instructions = slices.Delete(Instructions, instrNum, instrNum+1)

		ObjCoords = make(map[int]g143.RectSpecs)
		DrawWorkView(window, CurrentPage)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	}

}
