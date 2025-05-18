package main

import (
	"fmt"
	"strconv"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawBeginView(window *glfw.Window) {
	ProjObjCoords = make(map[int]g143.Rect)
	wWidth, wHeight := window.GetSize()

	theCtx := New2dCtx(wWidth, wHeight, &ProjObjCoords)

	fontPath := GetDefaultFontPath()
	theCtx.ggCtx.LoadFontFace(fontPath, 30)

	theCtx.ggCtx.SetHexColor(fontColor)
	theCtx.ggCtx.DrawString("New Project", 20, 10+30)

	theCtx.ggCtx.LoadFontFace(fontPath, 20)
	pnIRect := theCtx.drawInput(PROJ_NameInput, 20, 60, 420, "enter project name", false)
	pnBtnX, pnBtnY := nextHorizontalCoords(pnIRect, 30)
	nPRS := theCtx.drawButtonA(PROJ_NewProject, pnBtnX, pnBtnY, "New Project", fontColor, "#B3AE97")
	oWDBX, _ := nextHorizontalCoords(nPRS, 40)
	theCtx.drawButtonB(OpenWDBtn, oWDBX, 10, "Open Working Directory", "#fff", "#56845A", "#56845A")

	// second row border
	_, borderY := nextVerticalCoords(pnIRect, 10)
	theCtx.ggCtx.SetHexColor("#999")
	theCtx.ggCtx.DrawRoundedRectangle(10, float64(borderY), float64(wWidth)-20, 2, 2)
	theCtx.ggCtx.Fill()

	theCtx.ggCtx.LoadFontFace(fontPath, 30)
	theCtx.ggCtx.SetHexColor(fontColor)
	theCtx.ggCtx.DrawString("Continue Projects", 20, float64(borderY)+12+30)
	theCtx.ggCtx.LoadFontFace(fontPath, 20)

	projectFiles := GetProjectFiles()
	currentX := 40
	currentY := borderY + 22 + 30 + 10
	for i, pf := range projectFiles {

		btnId := 1000 + (i + 1)
		pfRect := theCtx.drawButtonA(btnId, currentX, currentY, pf.Name, "#fff", "#5F699F")

		newX := currentX + pfRect.Width + 10
		if newX > (wWidth - pfRect.Width) {
			currentY += 50
			currentX = 40
		} else {
			currentX += pfRect.Width + 10
		}

	}

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, theCtx.ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = theCtx.ggCtx.Image()
}

func DrawWorkView(window *glfw.Window, page int) {
	CurrentPage = page

	window.SetTitle(fmt.Sprintf("Project: %s ---- %s", ProjectName, ProgTitle))

	ObjCoords = make(map[int]g143.Rect)

	wWidth, wHeight := window.GetSize()
	theCtx := New2dCtx(wWidth, wHeight, &ObjCoords)

	// draw top buttons
	aIBRect := theCtx.drawButtonB(AddImgBtn, 50, 10, "Add Image", "#fff", "#5C909C", "#286775")
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

var scrollEventCount = 0

func FirstUIScrollCallback(window *glfw.Window, xoff, yoff float64) {

	if scrollEventCount != 5 {
		scrollEventCount += 1
		return
	}

	scrollEventCount = 0

	if xoff == 0 && yoff == -1 && CurrentPage != TotalPages() {
		ObjCoords = make(map[int]g143.Rect)
		DrawWorkView(window, CurrentPage+1)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	} else if xoff == 0 && yoff == 1 && CurrentPage != 1 {
		ObjCoords = make(map[int]g143.Rect)
		DrawWorkView(window, CurrentPage-1)
		window.SetCursorPosCallback(getHoverCB(ObjCoords))
	}

}
