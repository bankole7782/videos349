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
	OldFrame     image.Image
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

func (ctx *Ctx) drawInput(inputId, originX, originY, inputWidth int, placeholder string) g143.Rect {
	height := 40
	ctx.ggCtx.SetHexColor(fontColor)
	ctx.ggCtx.DrawRectangle(float64(originX), float64(originY), float64(inputWidth), float64(height))
	ctx.ggCtx.Fill()

	ctx.ggCtx.SetHexColor("#fff")
	ctx.ggCtx.DrawRectangle(float64(originX)+3, float64(originY)+3, float64(inputWidth)-6, float64(height)-6)
	ctx.ggCtx.Fill()

	entryRect := g143.Rect{Width: inputWidth, Height: height, OriginX: originX, OriginY: originY}
	(*ctx.ObjCoords)[inputId] = entryRect

	ctx.ggCtx.SetHexColor("#aaa")
	ctx.ggCtx.DrawString(placeholder, float64(originX+15), float64(originY)+FontSize+5)
	return entryRect
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
