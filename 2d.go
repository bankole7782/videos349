package main

import (
	"image"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
)

type Ctx struct {
	WindowWidth  int
	WindowHeight int
	ggCtx        *gg.Context
	ObjCoords    *map[int]g143.Rect
}

func New2dCtx(wWidth, wHeight int, objCoords *map[int]g143.Rect) Ctx {
	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	ctx := Ctx{WindowWidth: wWidth, WindowHeight: wHeight, ggCtx: ggCtx,
		ObjCoords: objCoords}
	return ctx
}

func Continue2dCtx(img image.Image, objCoords *map[int]g143.Rect) Ctx {
	ggCtx := gg.NewContextForImage(img)

	// load font
	fontPath := GetDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	ctx := Ctx{WindowWidth: img.Bounds().Dx(), WindowHeight: img.Bounds().Dy(), ggCtx: ggCtx,
		ObjCoords: objCoords}
	return ctx
}

func (ctx *Ctx) drawButtonA(btnId, originX, originY int, text, textColor, bgColor string) g143.Rect {
	// draw bounding rect
	textW, textH := ctx.ggCtx.MeasureString(text)
	width, height := textW+20, textH+15
	ctx.ggCtx.SetHexColor(bgColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(width), float64(height))
	ctx.ggCtx.Fill()

	// draw text
	ctx.ggCtx.SetHexColor(textColor)
	ctx.ggCtx.DrawString(text, float64(originX)+10, float64(originY)+FontSize)

	// save dimensions
	btnARect := g143.NewRect(originX, originY, int(width), int(height))
	(*ctx.ObjCoords)[btnId] = btnARect
	return btnARect
}

func (ctx *Ctx) drawButtonB(btnId, originX, originY int, text, textColor, bgColor, circleColor string) g143.Rect {
	// draw bounding rect
	textW, textH := ctx.ggCtx.MeasureString(text)
	width, height := textW+80, textH+30
	ctx.ggCtx.SetHexColor(bgColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(width), float64(height))
	ctx.ggCtx.Fill()

	// draw text
	ctx.ggCtx.SetHexColor(textColor)
	ctx.ggCtx.DrawString(text, float64(originX)+20, float64(originY)+FontSize+10)

	// draw circle
	ctx.ggCtx.SetHexColor(circleColor)
	ctx.ggCtx.DrawCircle(float64(originX)+width-30, float64(originY)+(height/2), 10)
	ctx.ggCtx.Fill()

	// save dimensions
	btnARect := g143.NewRect(originX, originY, int(width), int(height))
	(*ctx.ObjCoords)[btnId] = btnARect
	return btnARect
}

func (ctx *Ctx) drawButtonC(btnId, originX, originY int, bgColor string) g143.Rect {
	// draw bounding rect
	width, height := FontSize, FontSize
	ctx.ggCtx.SetHexColor(bgColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY)+2, float64(width), float64(height))
	ctx.ggCtx.Fill()

	// save dimensions
	btnARect := g143.NewRect(originX, originY, int(width), int(height))
	(*ctx.ObjCoords)[btnId] = btnARect
	return btnARect
}

func (ctx *Ctx) drawInput(inputId, originX, originY, inputWidth int, placeholder string, isDefault bool) g143.Rect {
	height := 30
	ctx.ggCtx.SetHexColor(fontColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(inputWidth), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+2, float64(originY)+2, float64(inputWidth)-4, float64(height)-4)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: inputWidth, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	if isDefault {
		ctx.ggCtx.SetHexColor("#444")
		ctx.ggCtx.DrawString(placeholder, float64(originX+15), float64(originY)+FontSize)
	} else {
		ctx.ggCtx.SetHexColor("#aaa")
		ctx.ggCtx.DrawString(placeholder, float64(originX+15), float64(originY)+FontSize)
	}
	return entryRect
}

func (ctx *Ctx) drawFileInput(inputId, originX, originY, inputWidth int, placeholder string) g143.Rect {
	ctx.ggCtx.SetHexColor("#eee")
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(inputWidth), 30)
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#444")
	placeholderW, _ := ctx.ggCtx.MeasureString(placeholder)
	placeholderX := originX + (inputWidth-int(placeholderW))/2
	ctx.ggCtx.DrawString(placeholder, float64(placeholderX), float64(originY)+FontSize)

	fIRect := g143.NewRect(originX, originY, inputWidth, 30)
	(*ctx.ObjCoords)[inputId] = fIRect
	return fIRect
}

func (ctx *Ctx) drawCheckbox(inputId, originX, originY int, isSelected bool) g143.Rect {
	width := 30
	height := 30
	ctx.ggCtx.SetHexColor(fontColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(width), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+2, float64(originY)+2, float64(width)-4, float64(height)-4)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: width, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	if isSelected {
		ctx.ggCtx.SetHexColor("#444")
		ctx.ggCtx.DrawRectangle(float64(originX)+4, float64(originY)+4, float64(width)-8, float64(height)-8)
		ctx.ggCtx.Fill()
	}
	return entryRect
}

func (ctx *Ctx) windowRect() g143.Rect {
	return g143.NewRect(0, 0, ctx.WindowWidth, ctx.WindowHeight)
}

func nextHorizontalCoords(aRect g143.Rect, margin int) (int, int) {
	nextOriginX := aRect.OriginX + aRect.Width + margin
	nextOriginY := aRect.OriginY
	return nextOriginX, nextOriginY
}

func nextVerticalCoords(aRect g143.Rect, margin int) (int, int) {
	nextOriginX := margin
	nextOriginY := aRect.OriginY + aRect.Height + margin
	return nextOriginX, nextOriginY
}
