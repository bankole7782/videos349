package main

import (
	"image"
	"slices"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func drawViewAddImage(window *glfw.Window, currentFrame image.Image) {
	VAIObjCoords = make(map[int]g143.Rect)
	VAIInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()
	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	theCtx := Continue2dCtx(img, &VAIObjCoords)

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 200

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	theCtx.ggCtx.SetHexColor("#fff")
	theCtx.ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	theCtx.ggCtx.Fill()

	// Add Image
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString("Add Image Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	addBtnOriginX := dialogWidth + dialogOriginX - 160
	addBtnRect := theCtx.drawButtonA(VAI_AddBtn, addBtnOriginX, dialogOriginY+20, "Add", "#fff", "#56845A")
	closeBtnX, _ := nextHorizontalCoords(addBtnRect, 10)
	theCtx.drawButtonA(VAI_CloseBtn, closeBtnX, addBtnRect.OriginY, "Close", "#fff", "#B75F5F")

	placeholder := "[click to pick an image]"
	if IsUpdateDialog {
		filename := Instructions[ToUpdateInstrNum]["image"]
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		placeholder = displayFilename
	}
	pHRect := theCtx.drawFileInput(VAI_SelectImg, dialogOriginX+20, dialogOriginY+40+30, dialogWidth-40, placeholder)
	_, durLabelY := nextVerticalCoords(pHRect, 30)
	durLabel := "duration (in seconds)"
	durLabelW, _ := theCtx.ggCtx.MeasureString(durLabel)
	theCtx.ggCtx.SetHexColor("#444")
	theCtx.ggCtx.DrawString(durLabel, float64(dialogOriginX)+20, float64(durLabelY))

	theCtx.drawInput(VAI_DurInput, dialogOriginX+int(durLabelW)+40, durLabelY-FontSize, 80, "5", true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

func vAIMouseCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range VAIObjCoords {
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

		drawItemsView(window, CurrentPage)
		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetScrollCallback(iVScrollBtnCB)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	case VAI_SelectImg:
		filename := PickImageFile()
		if filename == "" {
			return
		}
		VAIInputsStore["image"] = filename
		rootPath, _ := GetRootPath()
		displayFilename := strings.ReplaceAll(filename, rootPath, "")

		theCtx := Continue2dCtx(CurrentWindowFrame, &VAIObjCoords)
		theCtx.drawFileInput(VAI_SelectImg, widgetRS.OriginX, widgetRS.OriginY, widgetRS.Width, displayFilename)

		// send the frame to glfw window
		g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), theCtx.windowRect())
		window.SwapBuffers()

		// save the frame
		CurrentWindowFrame = theCtx.ggCtx.Image()

	case VAI_AddBtn:

		if IsUpdateDialog {
			oldInstr := Instructions[ToUpdateInstrNum]
			if filename, ok := VAIInputsStore["image"]; ok {
				oldInstr["image"] = filename
			}
			if VAI_DurationEnteredTxt != "" {
				oldInstr["duration"] = VAI_DurationEnteredTxt
				VAI_DurationEnteredTxt = ""
			}

			Instructions[ToUpdateInstrNum] = oldInstr
			IsUpdateDialog = false

		} else {

			if VAIInputsStore["image"] == "" {
				return
			}

			if VAI_DurationEnteredTxt == "" {
				VAIInputsStore["duration"] = "5"
			} else {
				VAIInputsStore["duration"] = VAI_DurationEnteredTxt
				VAI_DurationEnteredTxt = ""
			}

			if IsInsertBeforeDialog {
				item := map[string]string{
					"kind":     "image",
					"image":    VAIInputsStore["image"],
					"duration": VAIInputsStore["duration"],
				}
				Instructions = slices.Insert(Instructions, ToInsertBefore, item)
				IsInsertBeforeDialog = false
			} else {
				Instructions = append(Instructions, map[string]string{
					"kind":     "image",
					"image":    VAIInputsStore["image"],
					"duration": VAIInputsStore["duration"],
				})

			}

		}

		drawItemsView(window, TotalPages())

		// register the ViewMain mouse callback
		window.SetMouseButtonCallback(iVMouseBtnCB)
		// unregister the keyCallback
		window.SetKeyCallback(nil)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))

	}

}

func vAIkeyCB(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
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

	dIRS := VAIObjCoords[VAI_DurInput]
	theCtx := Continue2dCtx(CurrentWindowFrame, &VAIObjCoords)
	theCtx.drawInput(VAI_DurInput, dIRS.OriginX, dIRS.OriginY, dIRS.Width, VAI_DurationEnteredTxt, true)

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}
