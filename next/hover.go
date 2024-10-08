package main

import (
	"image"
	"runtime"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func getHoverCB(state map[int]g143.Rect) glfw.CursorPosCallback {
	return func(window *glfw.Window, xpos, ypos float64) {
		if runtime.GOOS == "linux" {
			// linux fires too many events
			cursorEventsCount += 1
			if cursorEventsCount != 10 {
				return
			} else {
				cursorEventsCount = 0
			}
		}

		wWidth, wHeight := window.GetSize()

		var widgetRS g143.Rect
		var widgetCode int

		xPosInt := int(xpos)
		yPosInt := int(ypos)
		for code, RS := range state {
			if g143.InRect(RS, xPosInt, yPosInt) {
				widgetRS = RS
				widgetCode = code
				break
			}
		}

		if widgetCode == 0 {
			// send the last drawn frame to glfw window
			windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
			g143.DrawImage(wWidth, wHeight, CurrentWindowFrame, windowRS)
			window.SwapBuffers()
			return
		}

		rectA := image.Rect(widgetRS.OriginX, widgetRS.OriginY,
			widgetRS.OriginX+widgetRS.Width,
			widgetRS.OriginY+widgetRS.Height)

		pieceOfCurrentFrame := imaging.Crop(CurrentWindowFrame, rectA)
		invertedPiece := imaging.AdjustBrightness(pieceOfCurrentFrame, -20)

		ggCtx := gg.NewContextForImage(CurrentWindowFrame)
		ggCtx.DrawImage(invertedPiece, widgetRS.OriginX, widgetRS.OriginY)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()
	}
}
