package internal

import (
	"fmt"
	"strconv"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawBeginView(window *glfw.Window) {
	ProjObjCoords = make(map[int]g143.RectSpecs)
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 30)
	if err != nil {
		panic(err)
	}

	// first column
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("New Video", 30, 30+30)

	// project name input
	ggCtx.LoadFontFace(fontPath, 20)
	ggCtx.SetHexColor("#444")
	ggCtx.DrawRectangle(30, 90, 420, 40)
	ggCtx.Fill()
	pniRS := g143.NRectSpecs(30, 90, 420, 40)
	ProjObjCoords[PROJ_NameInput] = pniRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(33, 93, 420-6, 40-6)
	ggCtx.Fill()

	npStr := "New Video"
	npStrW, _ := ggCtx.MeasureString(npStr)
	npBtnW := npStrW + 40
	ggCtx.SetHexColor("#B3AE97")
	ggCtx.DrawRoundedRectangle(200, 140, npBtnW, 50, 10)
	ggCtx.Fill()
	ProjObjCoords[PROJ_NewProject] = g143.NRectSpecs(200, 140, int(npBtnW), 50)

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(npStr, 200+20, 150+20)

	// second column
	ggCtx.SetHexColor("#999")
	ggCtx.DrawRoundedRectangle(500, 30, 4, 700, 2)
	ggCtx.Fill()

	ggCtx.LoadFontFace(fontPath, 40)
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Continue Videos", 550, 30+30)

	ggCtx.SetHexColor("#444")
	ggCtx.LoadFontFace(fontPath, 20)

	projectFiles := GetProjectFiles()
	currentX := 550
	currentY := 60
	for i, pf := range projectFiles {
		pfStrW, _ := ggCtx.MeasureString(pf.Name)

		ggCtx.SetHexColor("#5F699F")
		ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30, pfStrW+20, FontSize+10, 10)
		ggCtx.Fill()

		pfRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30,
			Width: int(pfStrW) + 20, Height: FontSize + 10}
		ProjObjCoords[1000+(i+1)] = pfRS

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawString(pf.Name, float64(currentX)+10, float64(currentY)+FontSize+30)

		newX := currentX + int(pfStrW) + 30
		if newX > (wWidth - int(pfStrW)) {
			currentY += 40
			currentX = 550
		} else {
			currentX += int(pfStrW) + 20 + 10
		}

	}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}

// func AllDraws(window *glfw.Window) {
func DrawWorkView(window *glfw.Window, page int) {
	CurrentPage = page

	window.SetTitle(fmt.Sprintf("Project: %s ---- %s", ProjectName, ProgTitle))

	ObjCoords = make(map[int]g143.RectSpecs)

	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// // intro text
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	// add image button
	addImgStr := "Add Image"
	addImgStrWidth, addImgStrHeight := ggCtx.MeasureString(addImgStr)
	addImgBtnWidth := addImgStrWidth + 80
	addImgBtnHeight := addImgStrHeight + 30
	ggCtx.SetHexColor("#5C909C")
	ggCtx.DrawRoundedRectangle(10, 10, addImgBtnWidth, addImgBtnHeight, addImgBtnHeight/2)
	ggCtx.Fill()

	addImgBtnRS := g143.RectSpecs{Width: int(addImgBtnWidth), Height: int(addImgBtnHeight),
		OriginX: 10, OriginY: 10}
	ObjCoords[AddImgBtn] = addImgBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(addImgStr, 30, 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#286775")
	ggCtx.DrawCircle(10+addImgBtnWidth-40, 10+addImgBtnHeight/2, 10)
	ggCtx.Fill()

	// Add Image + Audio Button
	aisStr := "Add Image + Audio"
	aisStrWidth, aisStrHeight := ggCtx.MeasureString(aisStr)
	aisBtnWidth := aisStrWidth + 80
	aisBtnHeight := aisStrHeight + 30
	ggCtx.SetHexColor("#5C909C")
	aisBtnOriginX := float64(addImgBtnRS.Width+addImgBtnRS.OriginX) + 10 // gutter
	ggCtx.DrawRoundedRectangle(aisBtnOriginX, 10, aisBtnWidth, aisBtnHeight, aisBtnHeight/2)
	ggCtx.Fill()

	aisBtnRS := g143.RectSpecs{Width: int(aisBtnWidth), Height: int(aisBtnHeight),
		OriginX: int(aisBtnOriginX), OriginY: 10}
	ObjCoords[AddImgSoundBtn] = aisBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(aisStr, 30+float64(aisBtnRS.OriginX), 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#286775")
	ggCtx.DrawCircle(float64(aisBtnRS.OriginX)+aisBtnWidth-30, 10+aisBtnHeight/2, 10)
	ggCtx.Fill()

	// Add Video Button
	addVidStr := "Add Video"
	addVidStrWidth, addVidStrHeight := ggCtx.MeasureString(addVidStr)
	addVidBtnWidth := addVidStrWidth + 80
	addVidBtnHeight := addVidStrHeight + 30
	ggCtx.SetHexColor("#81577F")
	addVidBtnOriginX := float64(aisBtnRS.Width+aisBtnRS.OriginX) + 10 // gutter
	ggCtx.DrawRoundedRectangle(addVidBtnOriginX, 10, addVidBtnWidth, addVidBtnHeight, addVidBtnHeight/2)
	ggCtx.Fill()

	addVidBtnRS := g143.RectSpecs{Width: int(addVidBtnWidth), Height: int(addVidBtnHeight),
		OriginX: int(addVidBtnOriginX), OriginY: 10}
	ObjCoords[AddVidBtn] = addVidBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(addVidStr, 30+float64(addVidBtnRS.OriginX), 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#633260")
	ggCtx.DrawCircle(float64(addVidBtnRS.OriginX)+addVidBtnWidth-30, 10+addVidBtnHeight/2, 10)
	ggCtx.Fill()

	// Open Working Directory button
	owdStr := "Open Working Directory"
	owdStrWidth, owdStrHeight := ggCtx.MeasureString(owdStr)
	openWDBtnWidth := owdStrWidth + 60
	openWDBtnHeight := owdStrHeight + 30
	ggCtx.SetHexColor("#56845A")
	openWDBtnOriginX := float64(addVidBtnRS.OriginX+addVidBtnRS.Width) + 40
	ggCtx.DrawRoundedRectangle(openWDBtnOriginX, 10, openWDBtnWidth, openWDBtnHeight, openWDBtnHeight/2)
	ggCtx.Fill()

	openWDBtnRS := g143.RectSpecs{Width: int(openWDBtnWidth), Height: int(openWDBtnHeight),
		OriginX: int(openWDBtnOriginX), OriginY: 10}
	ObjCoords[OpenWDBtn] = openWDBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(owdStr, 30+float64(openWDBtnRS.OriginX), 10+owdStrHeight+15)

	// Render button
	rbStr := "Render Video"
	rbStrW, rbStrH := ggCtx.MeasureString(rbStr)
	renderBtnW := rbStrW + 60
	renderBtnH := rbStrH + 30
	ggCtx.SetHexColor("#B19644")
	renderBtnX := openWDBtnRS.OriginX + openWDBtnRS.Width + 20
	ggCtx.DrawRoundedRectangle(float64(renderBtnX), 10, renderBtnW, renderBtnH, renderBtnH/2)
	ggCtx.Fill()

	rbRS := g143.RectSpecs{OriginX: renderBtnX, OriginY: 10, Width: int(renderBtnW),
		Height: int(renderBtnH)}
	ObjCoords[RenderBtn] = rbRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(rbStr, float64(rbRS.OriginX)+30, 10+rbStrH+15)
	// draw end of topbar demarcation
	ggCtx.SetHexColor("#999")
	ggCtx.DrawRectangle(10, float64(openWDBtnRS.OriginY+openWDBtnRS.Height+10), float64(wWidth)-20, 3)
	ggCtx.Fill()

	currentY := openWDBtnRS.OriginY + openWDBtnRS.Height + 10 + 10
	currentX := 20
	// show instructions
	shortInstrs := GetPageInstructions(page)
	for j, instr := range shortInstrs {
		// for i, instr := range Instructions {
		i := (PageSize * (page - 1)) + j
		kStr := strconv.Itoa(i+1) + "  [" + instr["kind"] + "]"
		kStrW, _ := ggCtx.MeasureString(kStr)

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(kStr, float64(currentX), float64(currentY)+FontSize)

		ggCtx.SetHexColor("#5A8A5E")
		ggCtx.DrawRoundedRectangle(float64(currentX)+kStrW+50, float64(currentY), FontSize, FontSize, FontSize/2)
		ggCtx.Fill()
		editRS := g143.NRectSpecs(currentX+int(kStrW)+50, currentY, FontSize, FontSize)
		ObjCoords[4000+(i+1)] = editRS

		ggCtx.SetHexColor("#A84E4E")
		ggCtx.DrawRoundedRectangle(float64(currentX)+kStrW+50+30, float64(currentY), FontSize, FontSize, FontSize/2)
		ggCtx.Fill()
		delRS := g143.NRectSpecs(currentX+int(kStrW)+50+30, currentY, FontSize, FontSize)
		ObjCoords[5000+(i+1)] = delRS

		viaStr := "View Image Asset #" + strconv.Itoa(i+1)
		viaStrW, _ := ggCtx.MeasureString(viaStr)

		if instr["kind"] == "image" {
			// view image asset
			viaStr = "View Image Asset #" + strconv.Itoa(i+1)
			viaStrW, _ := ggCtx.MeasureString(viaStr)

			ggCtx.SetHexColor("#5F699F")
			ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30, viaStrW+20, FontSize+10, 10)
			ggCtx.Fill()

			viabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30,
				Width: int(viaStrW) + 20, Height: FontSize + 10}
			ObjCoords[1000+(i+1)] = viabRS

			ggCtx.SetHexColor("#fff")
			ggCtx.DrawString(viaStr, float64(currentX)+10, float64(currentY)+FontSize+30)

			// duration
			var durStr string
			if _, ok := instr["audio"]; ok {
				durStr = "begin: " + instr["audio_begin"] + " | end: " + instr["audio_end"]
			} else {
				durStr = "duration: " + instr["duration"]
			}

			ggCtx.SetHexColor("#444")
			ggCtx.DrawString(durStr, float64(currentX), float64(currentY)+FontSize+30+15+FontSize)

			// view audio asset
			if _, ok := instr["audio"]; ok && instr["audio"] != "" {
				vaaStr := "View Audio Asset #" + strconv.Itoa(i+1)
				vaaStrW, _ := ggCtx.MeasureString(vaaStr)
				ggCtx.SetHexColor("#74A299")

				ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30+65, vaaStrW+20, FontSize+10, 10)
				ggCtx.Fill()
				vaabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30 + 65,
					Width: int(vaaStrW) + 20, Height: FontSize + 10}
				ObjCoords[2000+(i+1)] = vaabRS

				ggCtx.SetHexColor("#fff")
				ggCtx.DrawString(vaaStr, float64(currentX)+10, float64(currentY)+FontSize+30+65)
			}

		} else if instr["kind"] == "video" {
			// view video asset
			viaStr = "View Video Asset #" + strconv.Itoa(i+1)
			viaStrW, _ := ggCtx.MeasureString(viaStr)

			ggCtx.SetHexColor("#5F699F")
			ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30, viaStrW+20, FontSize+10, 10)
			ggCtx.Fill()
			vvabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30,
				Width: int(viaStrW) + 20, Height: FontSize + 10}
			ObjCoords[3000+(i+1)] = vvabRS

			ggCtx.SetHexColor("#fff")
			ggCtx.DrawString(viaStr, float64(currentX)+10, float64(currentY)+FontSize+30)

			// duration
			durStr := "begin: " + instr["begin"] + " | end: " + instr["end"]
			ggCtx.SetHexColor("#444")
			ggCtx.DrawString(durStr, float64(currentX), float64(currentY)+FontSize+30+15+FontSize)
		}

		newX := currentX + int(viaStrW) + 20
		if newX > (wWidth - int(viaStrW)) {
			currentY += 160
			currentX = 20
		} else {
			currentX += int(viaStrW) + 20 + 10
		}
	}

	// draw our site below
	ggCtx.SetHexColor("#9C5858")
	msg := fmt.Sprintf("VideoLength: %ds  Total Pages: %d  Current Page: %d", 0, TotalPages(), CurrentPage)
	fromAddrWidth, fromAddrHeight := ggCtx.MeasureString(msg)
	fromAddrOriginX := (wWidth - int(fromAddrWidth)) / 2
	ggCtx.DrawString(msg, float64(fromAddrOriginX), float64(wHeight-int(fromAddrHeight)))

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}

var scrollEventCount = 0

func FirstUIScrollCallback(window *glfw.Window, xoff, yoff float64) {

	if scrollEventCount != 5 {
		scrollEventCount += 1
		return
	}

	scrollEventCount = 0

	if xoff == 0 && yoff == -1 && CurrentPage != TotalPages() {
		ObjCoords = make(map[int]g143.RectSpecs)
		DrawWorkView(window, CurrentPage+1)
	} else if xoff == 0 && yoff == 1 && CurrentPage != 1 {
		ObjCoords = make(map[int]g143.RectSpecs)
		DrawWorkView(window, CurrentPage-1)
	}

}
