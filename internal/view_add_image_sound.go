package internal

import (
	"image"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawViewAIS(window *glfw.Window, currentFrame image.Image) {
	VaisObjCoords = make(map[int]g143.RectSpecs)
	VaisInputsStore = make(map[string]string)

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
	dialogHeight := 480

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	ggCtx.Fill()

	// Add Image
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Add Image + Sound Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	// Add Btn
	ggCtx.SetHexColor("#56845A")
	addStrW, _ := ggCtx.MeasureString("Add")
	addBtnOriginX := dialogWidth - int(addStrW) - 50 + dialogOriginX
	ggCtx.DrawRoundedRectangle(float64(addBtnOriginX), float64(dialogOriginY)+20, addStrW+20, 30, 10)
	ggCtx.Fill()
	addBtnRS := g143.RectSpecs{OriginX: addBtnOriginX, OriginY: dialogOriginY + 20, Width: int(addStrW) + 20, Height: 30}
	VaisObjCoords[VAIS_AddBtn] = addBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Add", float64(addBtnOriginX)+10, float64(dialogOriginY)+20+20)

	// Close Btn
	ggCtx.SetHexColor("#B75F5F")
	closeStrW, _ := ggCtx.MeasureString("Close")
	closeBtnX := dialogOriginX + dialogWidth - 50 - int(addStrW) - 30 - int(closeStrW)
	ggCtx.DrawRoundedRectangle(float64(closeBtnX), float64(dialogOriginY)+20, closeStrW+20, 30, 10)
	ggCtx.Fill()
	closeBtnRS := g143.RectSpecs{OriginX: closeBtnX, OriginY: dialogOriginY + 20, Width: int(closeStrW) + 20, Height: 30}
	VaisObjCoords[VAIS_CloseBtn] = closeBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Close", float64(closeBtnX)+10, float64(dialogOriginY)+20+20)

	// add image form
	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX)+40, float64(dialogOriginY)+20+50, float64(dialogWidth)-80, 240, 10)
	ggCtx.Fill()
	selectImgRS := g143.RectSpecs{Width: dialogWidth - 80, Height: 240, OriginX: dialogOriginX + 40, OriginY: dialogOriginY + 20 + 50}
	VaisObjCoords[VAIS_SelectImg] = selectImgRS

	aicStr := "[click to pick an image]"
	ggCtx.SetHexColor("#444")
	aicStrW, _ := ggCtx.MeasureString(aicStr)
	aicStrOriginX := selectImgRS.OriginX + (selectImgRS.Width-int(aicStrW))/2
	ggCtx.DrawString(aicStr, float64(aicStrOriginX), float64(selectImgRS.OriginY)+20+20)

	// audio file input
	ggCtx.SetHexColor("#eee")
	sflInputY := selectImgRS.OriginY + selectImgRS.Height + 10
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX)+40, float64(sflInputY), float64(dialogWidth)-80, 30, 10)
	ggCtx.Fill()
	sflBtnRS := g143.NRectSpecs(dialogOriginX+40, sflInputY, dialogWidth-80, 30)
	VaisObjCoords[VAIS_SelectAudio] = sflBtnRS

	pAFL := "[click to pick audio]"
	pAFLW, _ := ggCtx.MeasureString(pAFL)
	pAFLX := sflBtnRS.OriginX + (sflBtnRS.Width-int(pAFLW))/2
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString(pAFL, float64(pAFLX), float64(sflInputY)+FontSize)

	// audio begin
	aBL := "audio begin (mm:ss)"
	aBLW, _ := ggCtx.MeasureString(aBL)
	aBLY := sflBtnRS.OriginY + sflBtnRS.Height + 10 + 20
	ggCtx.DrawString(aBL, float64(dialogOriginX)+40, float64(aBLY))

	ggCtx.SetHexColor("#eee")
	aBInputX := aBLW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(aBInputX, float64(aBLY)-FontSize, 100, 30, 10)
	ggCtx.Fill()
	aBInputRS := g143.NRectSpecs(int(aBInputX), aBLY-FontSize, 100, 30)
	VaisObjCoords[VAIS_AudioBeginInput] = aBInputRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", aBInputX+10, float64(aBLY))

	// audio end
	aEL := "audio end (mm:ss)"
	aELW, _ := ggCtx.MeasureString(aEL)
	aELY := aBInputRS.OriginY + aBInputRS.Height + 10 + 20
	ggCtx.DrawString(aEL, float64(dialogOriginX)+40, float64(aELY))

	ggCtx.SetHexColor("#eee")
	aEInputX := aELW + 50 + float64(dialogOriginX) + 40
	ggCtx.DrawRoundedRectangle(aEInputX, float64(aELY)-FontSize, 100, 30, 10)
	ggCtx.Fill()
	aEInputRS := g143.NRectSpecs(int(aEInputX), aELY-FontSize, 100, 30)
	VaisObjCoords[VAIS_AudioEndInput] = aEInputRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", aEInputX+10, float64(aELY))

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}
