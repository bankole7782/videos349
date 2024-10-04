package main

import (
	"image"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawViewAddImage(window *glfw.Window, currentFrame image.Image) {
	VaiObjCoords = make(map[int]g143.RectSpecs)
	VaiInputsStore = make(map[string]string)

	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background image
	img := imaging.AdjustBrightness(currentFrame, -40)
	ggCtx.DrawImage(img, 0, 0)

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	// dialog rectangle
	dialogWidth := 600
	dialogHeight := 400

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	ggCtx.Fill()

	// Add Image
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Add Image Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	// Add Btn
	ggCtx.SetHexColor("#56845A")
	addStrW, _ := ggCtx.MeasureString("Add")
	addBtnOriginX := dialogWidth - int(addStrW) - 50 + dialogOriginX
	ggCtx.DrawRectangle(float64(addBtnOriginX), float64(dialogOriginY)+20, addStrW+20, 30)
	ggCtx.Fill()
	addBtnRS := g143.RectSpecs{OriginX: addBtnOriginX, OriginY: dialogOriginY + 20, Width: int(addStrW) + 20, Height: 30}
	VaiObjCoords[VAI_AddBtn] = addBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Add", float64(addBtnOriginX)+10, float64(dialogOriginY)+20+20)

	// Close Btn
	ggCtx.SetHexColor("#B75F5F")
	closeStrW, _ := ggCtx.MeasureString("Close")
	closeBtnX := dialogOriginX + dialogWidth - 50 - int(addStrW) - 30 - int(closeStrW)
	ggCtx.DrawRectangle(float64(closeBtnX), float64(dialogOriginY)+20, closeStrW+20, 30)
	ggCtx.Fill()
	closeBtnRS := g143.RectSpecs{OriginX: closeBtnX, OriginY: dialogOriginY + 20, Width: int(closeStrW) + 20, Height: 30}
	VaiObjCoords[VAI_CloseBtn] = closeBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Close", float64(closeBtnX)+10, float64(dialogOriginY)+20+20)

	// add image form
	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRectangle(float64(dialogOriginX)+40, float64(dialogOriginY)+20+50, float64(dialogWidth)-80, 240)
	ggCtx.Fill()
	selectImgRS := g143.RectSpecs{Width: dialogWidth - 80, Height: 240, OriginX: dialogOriginX + 40, OriginY: dialogOriginY + 20 + 50}
	VaiObjCoords[VAI_SelectImg] = selectImgRS

	aicStr := "[click to pick an image]"
	ggCtx.SetHexColor("#444")
	aicStrW, _ := ggCtx.MeasureString(aicStr)
	aicStrOriginX := selectImgRS.OriginX + (selectImgRS.Width-int(aicStrW))/2
	ggCtx.DrawString(aicStr, float64(aicStrOriginX), float64(selectImgRS.OriginY)+20+20)

	if IsUpdateDialog {
		// show picked image
		filename := Instructions[ToUpdateInstrNum]["image"]
		img, _ := imaging.Open(filename)
		img = imaging.Fit(img, selectImgRS.Width-20, selectImgRS.Height-20, imaging.Lanczos)
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(selectImgRS.OriginX), float64(selectImgRS.OriginY), float64(selectImgRS.Width),
			float64(selectImgRS.Height), 10)
		ggCtx.Fill()
		ggCtx.DrawImage(img, selectImgRS.OriginX+10, selectImgRS.OriginY+10)
	}

	durLabel := "duration (in seconds)"
	durLabelW, _ := ggCtx.MeasureString(durLabel)
	durLabelY := selectImgRS.OriginY + selectImgRS.Height + 10 + 20
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(durLabel, float64(dialogOriginX)+40, float64(durLabelY))

	ggCtx.SetHexColor("#eee")
	durInputX := durLabelW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRectangle(durInputX, float64(durLabelY)-FontSize, 100, 30)
	ggCtx.Fill()
	durInputRS := g143.RectSpecs{OriginX: int(durInputX), OriginY: durLabelY - 20, Width: 100, Height: 30}
	VaiObjCoords[VAI_DurInput] = durInputRS

	// default duration
	ggCtx.SetHexColor("#444")
	if IsUpdateDialog {
		durationStr := Instructions[ToUpdateInstrNum]["duration"]
		ggCtx.DrawString(durationStr, durInputX+20, float64(durLabelY))
	} else {
		ggCtx.DrawString("5", durInputX+20, float64(durLabelY))
	}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}
