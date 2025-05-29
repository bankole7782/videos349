package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func drawItemsView(window *glfw.Window, page int) {
	CurrentPage = page

	window.SetTitle(fmt.Sprintf("Project: %s ---- %s", ProjectName, ProgTitle))

	ObjCoords = make(map[int]g143.Rect)

	wWidth, wHeight := window.GetSize()
	theCtx := New2dCtx(wWidth, wHeight, &ObjCoords)

	// draw top buttons
	bBRect := theCtx.drawButtonB(BackBtn, 50, 10, "Back", "#fff", "#845B5B", "#845B5B")
	aIBX, aIBY := nextHorizontalCoords(bBRect, 40)
	aIBRect := theCtx.drawButtonB(AddImgBtn, aIBX, aIBY, "Add Image", "#fff", "#5C909C", "#286775")
	aISX, aISY := nextHorizontalCoords(aIBRect, 10)
	aIBSRect := theCtx.drawButtonB(AddImgSoundBtn, aISX, aISY, "Add Image + Audio", "#fff", "#5C909C", "#286775")
	aVBX, aVBY := nextHorizontalCoords(aIBSRect, 10)
	aVBRect := theCtx.drawButtonB(AddVidBtn, aVBX, aVBY, "Add Video", "#fff", "#81577F", "#633260")
	rBX, rBY := nextHorizontalCoords(aVBRect, 10)
	theCtx.drawButtonB(RenderBtn, rBX, rBY, "Render", "#fff", "#B19644", "#DECC6E")

	// draw end of topbar demarcation
	_, demarcY := nextVerticalCoords(aVBRect, 10)
	theCtx.ggCtx.SetHexColor("#aaa")
	theCtx.ggCtx.DrawRectangle(10, float64(demarcY), float64(wWidth)-20, 3)
	theCtx.ggCtx.Fill()

	// show instructions
	currentY := demarcY + 10
	currentX := 10
	shortInstrs := GetPageInstructions(page)
	for j, instr := range shortInstrs {
		// for i, instr := range Instructions {
		i := (PageSize * (page - 1)) + j

		// inbetween buttons
		iAIBtnId := 6000 + (i + 1)
		iAIBtnRect := theCtx.drawButtonC(iAIBtnId, currentX, currentY+20, "#5C909C")
		_, iAISBtnY := nextVerticalCoords(iAIBtnRect, 10)
		iAISBtnId := 7000 + (i + 1)
		iAISBtnRect := theCtx.drawButtonC(iAISBtnId, currentX, iAISBtnY, "#5C909C")
		iAVBtnId := 8000 + (i + 1)
		_, iAVBtnY := nextVerticalCoords(iAISBtnRect, 10)
		theCtx.drawButtonC(iAVBtnId, currentX, iAVBtnY, "#81577F")

		currentX += 40

		kStr := strconv.Itoa(i+1) + "  [" + instr["kind"] + "]"
		kStrW, _ := theCtx.ggCtx.MeasureString(kStr)

		theCtx.ggCtx.SetHexColor("#444")
		theCtx.ggCtx.DrawString(kStr, float64(currentX), float64(currentY)+FontSize)

		eBtnId := 4000 + (i + 1)
		editBtnX := currentX + int(kStrW) + 50
		eDBRect := theCtx.drawButtonC(eBtnId, editBtnX, currentY, "#5A8A5E")
		delBtnX, _ := nextHorizontalCoords(eDBRect, 10)
		delBtnId := 5000 + (i + 1)
		theCtx.drawButtonC(delBtnId, delBtnX, currentY, "#A84E4E")

		vBtnW := 0
		if instr["kind"] == "image" {
			viaStr := "View Image Asset #" + strconv.Itoa(i+1)
			vBtnId := 1000 + (i + 1)
			vBtnRect := theCtx.drawButtonA(vBtnId, currentX, currentY+30, viaStr, "#fff", "#5F699F")
			vBtnW = vBtnRect.Width
			_, durStrY := nextVerticalCoords(vBtnRect, 5)
			// duration
			var durStr string
			if _, ok := instr["audio"]; ok {
				durStr = "begin: " + instr["audio_begin"] + " | end: " + instr["audio_end"]
			} else {
				durStr = "duration: " + instr["duration"]
			}

			theCtx.ggCtx.SetHexColor("#444")
			theCtx.ggCtx.DrawString(durStr, float64(currentX), float64(durStrY)+FontSize)

			// view audio asset
			if _, ok := instr["audio"]; ok && instr["audio"] != "" {
				vaaBtnId := 2000 + (i + 1)
				vaaStr := "View Audio Asset #" + strconv.Itoa(i+1)
				vaaY := durStrY + FontSize + 10
				theCtx.drawButtonA(vaaBtnId, currentX, vaaY, vaaStr, "#fff", "#74A299")
			}

		} else if instr["kind"] == "video" {
			viaStr := "View Video Asset #" + strconv.Itoa(i+1)
			vVBtnId := 3000 + (i + 1)
			vVBtnRect := theCtx.drawButtonA(vVBtnId, currentX, currentY+30, viaStr, "#fff", "#5F699F")
			vBtnW = vVBtnRect.Width

			// duration
			durStr := "begin: " + instr["begin"] + " | end: " + instr["end"]
			theCtx.ggCtx.SetHexColor("#444")
			theCtx.ggCtx.DrawString(durStr, float64(currentX), float64(currentY)+FontSize+30+15+FontSize)
		}

		newX := currentX + vBtnW + 10
		if newX > (wWidth - vBtnW) {
			currentY += 160
			currentX = 20
		} else {
			currentX += vBtnW + 20
		}
	}

	// draw our site below
	theCtx.ggCtx.SetHexColor("#444")
	msg := fmt.Sprintf("VideoLength: %s  Total Pages: %d  Current Page: %d", TotalVideoLength(), TotalPages(), CurrentPage)
	fromAddrWidth, fromAddrHeight := theCtx.ggCtx.MeasureString(msg)
	fromAddrOriginX := (wWidth - int(fromAddrWidth)) / 2
	theCtx.ggCtx.DrawString(msg, float64(fromAddrOriginX), float64(wHeight-int(fromAddrHeight)))

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

func iVMouseBtnCB(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range ObjCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			// break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case BackBtn:
		// save work
		jsonBytes, _ := json.Marshal(Instructions)
		rootPath, _ := GetRootPath()
		outPath := filepath.Join(rootPath, ProjectName)
		os.WriteFile(outPath, jsonBytes, 0777)

		// clear some variables
		Instructions = make([]map[string]string, 0)
		ProjectName = ""
		window.SetTitle(ProgTitle)

		// redraw
		drawFirstView(window)
		window.SetMouseButtonCallback(fVMouseCB)
		window.SetKeyCallback(fVKeyCB)
		window.SetCursorPosCallback(getHoverCB(ProjObjCoords))
		window.SetScrollCallback(nil)

	case AddImgBtn:
		// tmpFrame = CurrentWindowFrame
		drawViewAddImage(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAIMouseCB)
		window.SetKeyCallback(vAIkeyCB)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VAIObjCoords))

	case AddImgSoundBtn:
		drawViewAIS(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAISMouseCB)
		window.SetKeyCallback(vAISKeyCB)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VAISObjCoords))

	case AddVidBtn:
		drawViewAddVideo(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAVMouseCB)
		window.SetKeyCallback(vAVKeyCB)
		window.SetScrollCallback(nil)
		window.SetCursorPosCallback(getHoverCB(VAVObjCoords))

	case RenderBtn:
		if len(Instructions) == 0 {
			return
		}
		SavedWorkViewFrame = CurrentWindowFrame
		DrawRenderView(window, CurrentWindowFrame, 0.0)
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
		IsUpdateDialog = true

		if Instructions[instrNum]["kind"] == "image" && Instructions[instrNum]["audio"] != "" {

			drawViewAIS(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(vAISMouseCB)
			window.SetKeyCallback(vAISKeyCB)
			window.SetCursorPosCallback(getHoverCB(VAISObjCoords))

		} else if Instructions[instrNum]["kind"] == "image" {
			drawViewAddImage(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(vAIMouseCB)
			window.SetKeyCallback(vAIkeyCB)
			window.SetCursorPosCallback(getHoverCB(VAIObjCoords))

		} else if Instructions[instrNum]["video"] != "" {
			drawViewAddVideo(window, CurrentWindowFrame)
			window.SetMouseButtonCallback(vAVMouseCB)
			window.SetKeyCallback(vAVKeyCB)
			window.SetCursorPosCallback(getHoverCB(VAVObjCoords))

		}
	} else if widgetCode > 5000 && widgetCode < 6000 {
		// delete from instructions slice
		instrNum := widgetCode - 5000 - 1
		Instructions = slices.Delete(Instructions, instrNum, instrNum+1)

		ObjCoords = make(map[int]g143.Rect)
		drawItemsView(window, CurrentPage)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	} else if widgetCode > 6000 && widgetCode < 7000 {
		instrNum := widgetCode - 6000 - 1
		IsInsertBeforeDialog = true
		ToInsertBefore = instrNum

		drawViewAddImage(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAIMouseCB)
		window.SetKeyCallback(vAIkeyCB)
		window.SetCursorPosCallback(getHoverCB(VAIObjCoords))
	} else if widgetCode > 7000 && widgetCode < 8000 {
		instrNum := widgetCode - 7000 - 1
		IsInsertBeforeDialog = true
		ToInsertBefore = instrNum

		drawViewAIS(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAISMouseCB)
		window.SetKeyCallback(vAISKeyCB)
		window.SetCursorPosCallback(getHoverCB(VAISObjCoords))
	} else if widgetCode > 8000 && widgetCode < 9000 {
		instrNum := widgetCode - 8000 - 1
		IsInsertBeforeDialog = true
		ToInsertBefore = instrNum

		drawViewAddVideo(window, CurrentWindowFrame)
		window.SetMouseButtonCallback(vAVMouseCB)
		window.SetKeyCallback(vAVKeyCB)
		window.SetCursorPosCallback(getHoverCB(VAVObjCoords))
	}

}

func iVScrollBtnCB(window *glfw.Window, xoff, yoff float64) {

	if scrollEventCount != 5 {
		scrollEventCount += 1
		return
	}

	scrollEventCount = 0

	if xoff == 0 && yoff == -1 && CurrentPage != TotalPages() {
		ObjCoords = make(map[int]g143.Rect)
		drawItemsView(window, CurrentPage+1)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	} else if xoff == 0 && yoff == 1 && CurrentPage != 1 {
		ObjCoords = make(map[int]g143.Rect)
		drawItemsView(window, CurrentPage-1)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	}

}
