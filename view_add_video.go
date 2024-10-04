package main

import (
	"image"
	"strings"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func DrawViewAddVideo(window *glfw.Window, currentFrame image.Image) {
	VavObjCoords = make(map[int]g143.Rect)
	VavInputsStore = make(map[string]string)

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
	dialogHeight := 280

	dialogOriginX := (wWidth - dialogWidth) / 2
	dialogOriginY := (wHeight - dialogHeight) / 2

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRoundedRectangle(float64(dialogOriginX), float64(dialogOriginY), float64(dialogWidth),
		float64(dialogHeight), 20)
	ggCtx.Fill()

	// Add Video Header
	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("Add Video Configuration", float64(dialogOriginX)+20, float64(dialogOriginY)+20+20)

	// Add Btn
	ggCtx.SetHexColor("#56845A")
	addStrW, _ := ggCtx.MeasureString("Add")
	addBtnOriginX := dialogWidth - int(addStrW) - 50 + dialogOriginX
	ggCtx.DrawRectangle(float64(addBtnOriginX), float64(dialogOriginY)+20, addStrW+20, 30)
	ggCtx.Fill()
	addBtnRS := g143.Rect{OriginX: addBtnOriginX, OriginY: dialogOriginY + 20, Width: int(addStrW) + 20, Height: 30}
	VavObjCoords[VAV_AddBtn] = addBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Add", float64(addBtnOriginX)+10, float64(dialogOriginY)+20+20)

	// Close Btn
	ggCtx.SetHexColor("#B75F5F")
	closeStrW, _ := ggCtx.MeasureString("Close")
	closeBtnX := dialogOriginX + dialogWidth - 50 - int(addStrW) - 30 - int(closeStrW)
	ggCtx.DrawRectangle(float64(closeBtnX), float64(dialogOriginY)+20, closeStrW+20, 30)
	ggCtx.Fill()
	closeBtnRS := g143.Rect{OriginX: closeBtnX, OriginY: dialogOriginY + 20, Width: int(closeStrW) + 20, Height: 30}
	VavObjCoords[VAV_CloseBtn] = closeBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString("Close", float64(closeBtnX)+10, float64(dialogOriginY)+20+20)
	// end of top bar

	// pick video label and btn
	pvInputX := dialogOriginX + 40
	pvInputW := dialogWidth - 80
	pvInputY := dialogOriginY + FontSize + 70

	ggCtx.SetHexColor("#eee")
	ggCtx.DrawRectangle(float64(pvInputX), float64(pvInputY)-FontSize, float64(pvInputW), 30)
	ggCtx.Fill()
	pfRS := g143.Rect{OriginX: int(pvInputX), OriginY: pvInputY - FontSize, Width: pvInputW, Height: 30}
	VavObjCoords[VAV_PickVideo] = pfRS

	ggCtx.SetHexColor("#444")
	pVFL := "[click to pick video file]"
	pVFLW, _ := ggCtx.MeasureString(pVFL)
	pVFLX := pvInputX + (pfRS.Width-int(pVFLW))/2
	ggCtx.DrawString(pVFL, float64(pVFLX), float64(pvInputY))

	// begin str label and input
	beginStr := "begin (mm:ss)"
	beginStrW, _ := ggCtx.MeasureString(beginStr)
	beginStrX := dialogOriginX + 40
	beginStrY := pvInputY + FontSize + 30
	ggCtx.DrawString(beginStr, float64(beginStrX), float64(beginStrY))

	ggCtx.SetHexColor("#eee")
	beginInputX := beginStrW + 30 + float64(dialogOriginX) + 40
	ggCtx.DrawRectangle(beginInputX, float64(beginStrY)-FontSize, 100, 30)
	ggCtx.Fill()
	biRS := g143.Rect{OriginX: int(beginInputX), OriginY: beginStrY - FontSize, Width: 100, Height: 30}
	VavObjCoords[VAV_BeginInput] = biRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", beginInputX+10, float64(beginStrY))

	// end str label and input
	endStr := "end (mm:ss)"
	endStrW, _ := ggCtx.MeasureString(endStr)
	endStrX := dialogOriginX + 40
	endStrY := beginStrY + FontSize + 30
	ggCtx.DrawString(endStr, float64(endStrX), float64(endStrY))

	ggCtx.SetHexColor("#eee")
	endInputX := endStrW + 30 + float64(dialogOriginX) + 40
	ggCtx.DrawRectangle(endInputX, float64(endStrY)-FontSize, 100, 30)
	ggCtx.Fill()
	eiRS := g143.Rect{OriginX: int(endInputX), OriginY: endStrY - FontSize, Width: 100, Height: 30}
	VavObjCoords[VAV_EndInput] = eiRS

	ggCtx.SetHexColor("#444")
	ggCtx.DrawString("0:00", endInputX+10, float64(endStrY))

	// speedUp checkbox
	ggCtx.SetHexColor("#444")
	suL := "speed up video"
	suLW, _ := ggCtx.MeasureString(suL)
	sulX := dialogOriginX + 40
	sulY := endStrY + FontSize + 30
	ggCtx.DrawString(suL, float64(sulX), float64(sulY))

	suCX := suLW + 30 + float64(dialogOriginX) + 40
	ggCtx.DrawRectangle(suCX, float64(sulY)-FontSize, 30, 30)
	ggCtx.Fill()
	suRS := g143.NewRect(int(suCX), sulY-FontSize, 30, 30)
	VavObjCoords[VAV_SpeedUpCheckbox] = suRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(suCX+2, float64(suRS.OriginY)+2, float64(suRS.Width)-4,
		float64(suRS.Height)-4)
	ggCtx.Fill()

	// blackAndWhite checkbox
	ggCtx.SetHexColor("#444")
	bwL := "black and white video"
	bwLW, _ := ggCtx.MeasureString(bwL)
	bwLX := suRS.OriginX + suRS.Width + 40
	ggCtx.DrawString(bwL, float64(bwLX), float64(sulY))

	bwCX := bwLX + int(bwLW) + 40
	ggCtx.DrawRectangle(float64(bwCX), float64(sulY)-FontSize, 30, 30)
	ggCtx.Fill()
	bwRS := g143.NewRect(bwCX, sulY-FontSize, 30, 30)
	VavObjCoords[VAV_BlackAndWhiteCheckbox] = bwRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(bwCX)+2, float64(bwRS.OriginY)+2, float64(bwRS.Width)-4,
		float64(bwRS.Height)-4)
	ggCtx.Fill()

	if IsUpdateDialog {
		// write picked video filename
		filename := Instructions[ToUpdateInstrNum]["video"]
		rootPath, _ := GetRootPath()
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(pfRS.OriginX), float64(pfRS.OriginY),
			float64(pfRS.Width), float64(pfRS.Height), 10)
		ggCtx.Fill()

		displayFilename := strings.ReplaceAll(filename, rootPath, "")
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(displayFilename, float64(pfRS.OriginX+10), float64(pfRS.OriginY+20))

		// write video begin
		beginStr := Instructions[ToUpdateInstrNum]["begin"]
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(biRS.OriginX), float64(biRS.OriginY), float64(biRS.Width),
			float64(biRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(beginStr, float64(biRS.OriginX+10), float64(biRS.OriginY)+FontSize)

		// write video end
		ggCtx.SetHexColor("#eee")
		ggCtx.DrawRoundedRectangle(float64(eiRS.OriginX), float64(eiRS.OriginY), float64(eiRS.Width),
			float64(eiRS.Height), 10)
		ggCtx.Fill()

		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(EndInputEnteredTxt, float64(eiRS.OriginX+10), float64(eiRS.OriginY)+FontSize)

		// draw speedup checkbox
		if strings.ToLower(Instructions[ToUpdateInstrNum]["speedup"]) == "true" {
			ggCtx.SetHexColor("#444")
			ggCtx.DrawRectangle(float64(suRS.OriginX)+4, float64(suRS.OriginY)+4, float64(suRS.Width)-8,
				float64(suRS.Height)-8)
			ggCtx.Fill()
		}

		if strings.ToLower(Instructions[ToUpdateInstrNum]["blackwhite"]) == "true" {
			ggCtx.SetHexColor("#444")
			ggCtx.DrawRectangle(float64(bwRS.OriginX)+4, float64(bwRS.OriginY)+4, float64(bwRS.Width)-8,
				float64(bwRS.Height)-8)
			ggCtx.Fill()
		}

	}
	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	CurrentWindowFrame = ggCtx.Image()
}
