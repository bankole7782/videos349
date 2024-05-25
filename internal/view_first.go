package internal

import (
	"strconv"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func AllDraws(window *glfw.Window) {
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
	for i, instr := range Instructions {
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(strconv.Itoa(i+1)+"  ["+instr["kind"]+"]", float64(currentX), float64(currentY)+FontSize)

		viaStr := "View Image Asset #" + strconv.Itoa(i+1)
		viaStrW, _ := ggCtx.MeasureString(viaStr)

		if instr["kind"] == "image" {
			// view image asset
			viaStr := "View Image Asset #" + strconv.Itoa(i+1)
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

			// view audio asset optional
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
			viaStr := "View Video Asset #" + strconv.Itoa(i+1)
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

			// view audio asset optional
			if instr["audio_optional"] != "" {
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
	fromAddr := "sae.ng"
	fromAddrWidth, fromAddrHeight := ggCtx.MeasureString(fromAddr)
	fromAddrOriginX := (wWidth - int(fromAddrWidth)) / 2
	ggCtx.DrawString(fromAddr, float64(fromAddrOriginX), float64(wHeight-int(fromAddrHeight)))
	fars := g143.RectSpecs{OriginX: fromAddrOriginX, OriginY: wHeight - 40,
		Width: int(fromAddrWidth), Height: 40}
	ObjCoords[OurSite] = fars

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}
