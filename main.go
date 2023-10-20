package main

import (
	"image"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps       = 10
	fontSize  = 20
	AddImgBtn = 11
	AddVidBtn = 12
	OpenWDBtn = 13
	RenderBtn = 14
)

var objCoords map[int]g143.RectSpecs
var currentWindowFrame image.Image
var instructions []map[string]string

// var tmpFrame image.Image
var inChannel chan bool
var clearAfterRender bool

func main() {
	_, err := GetRootPath()
	if err != nil {
		panic(err)
	}

	runtime.LockOSThread()

	objCoords = make(map[int]g143.RectSpecs)
	instructions = make([]map[string]string, 0)
	inChannel = make(chan bool)

	window := g143.NewWindow(1200, 800, "videos349: a simple video editor", false)
	allDraws(window)

	go func() {
		for {
			<-inChannel
			render(instructions)
			clearAfterRender = true
		}
	}()

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
	// respond to the keyboard
	// window.SetKeyCallback(keyCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		if clearAfterRender {
			// clear the UI and redraw
			instructions = make([]map[string]string, 0)
			allDraws(window)
			drawEndRenderView(window, currentWindowFrame)
			time.Sleep(5 * time.Second)
			allDraws(window)
			// register the ViewMain mouse callback
			window.SetMouseButtonCallback(mouseBtnCallback)
			clearAfterRender = false
		}

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}

}

func getDefaultFontPath() string {
	fontPath := filepath.Join(os.TempDir(), "v349_font.ttf")
	os.WriteFile(fontPath, DefaultFont, 0777)
	return fontPath
}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// // intro text
	fontPath := getDefaultFontPath()
	err := ggCtx.LoadFontFace(fontPath, 20)
	if err != nil {
		panic(err)
	}

	// add image button
	addImgStr := "Add Image"
	addImgStrWidth, addImgStrHeight := ggCtx.MeasureString(addImgStr)
	addImgBtnWidth := addImgStrWidth + 80
	addImgBtnHeight := addImgStrHeight + 30
	ggCtx.SetHexColor("#B75F5F")
	ggCtx.DrawRoundedRectangle(10, 10, addImgBtnWidth, addImgBtnHeight, addImgBtnHeight/2)
	ggCtx.Fill()

	addImgBtnRS := g143.RectSpecs{Width: int(addImgBtnWidth), Height: int(addImgBtnHeight),
		OriginX: 10, OriginY: 10}
	objCoords[AddImgBtn] = addImgBtnRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(addImgStr, 30, 10+addImgStrHeight+15)

	ggCtx.SetHexColor("#633232")
	ggCtx.DrawCircle(10+addImgBtnWidth-40, 10+addImgBtnHeight/2, 10)
	ggCtx.Fill()

	// Add Video Button
	addVidStr := "Add Video"
	addVidStrWidth, addVidStrHeight := ggCtx.MeasureString(addVidStr)
	addVidBtnWidth := addVidStrWidth + 80
	addVidBtnHeight := addVidStrHeight + 30
	ggCtx.SetHexColor("#81577F")
	addVidBtnOriginX := float64(addImgBtnRS.Width+addImgBtnRS.OriginX) + 10 // gutter
	ggCtx.DrawRoundedRectangle(addVidBtnOriginX, 10, addVidBtnWidth, addVidBtnHeight, addVidBtnHeight/2)
	ggCtx.Fill()

	addVidBtnRS := g143.RectSpecs{Width: int(addVidBtnWidth), Height: int(addVidBtnHeight),
		OriginX: int(addVidBtnOriginX), OriginY: 10}
	objCoords[AddVidBtn] = addVidBtnRS

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
	objCoords[OpenWDBtn] = openWDBtnRS

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
	objCoords[RenderBtn] = rbRS

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawString(rbStr, float64(rbRS.OriginX)+30, 10+rbStrH+15)
	// draw end of topbar demarcation
	ggCtx.SetHexColor("#999")
	ggCtx.DrawRectangle(10, float64(openWDBtnRS.OriginY+openWDBtnRS.Height+10), float64(wWidth)-20, 3)
	ggCtx.Fill()

	currentY := openWDBtnRS.OriginY + openWDBtnRS.Height + 10 + 10
	currentX := 20
	// show instructions
	for i, instr := range instructions {
		ggCtx.SetHexColor("#444")
		ggCtx.DrawString(strconv.Itoa(i+1)+"  ["+instr["kind"]+"]", float64(currentX), float64(currentY)+fontSize)

		viaStr := "View Image Asset #" + strconv.Itoa(i+1)
		viaStrW, _ := ggCtx.MeasureString(viaStr)

		if instr["kind"] == "image" {
			// view image asset
			viaStr := "View Image Asset #" + strconv.Itoa(i+1)
			viaStrW, _ := ggCtx.MeasureString(viaStr)

			ggCtx.SetHexColor("#5F699F")
			ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30, viaStrW+20, fontSize+10, 10)
			ggCtx.Fill()

			viabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30,
				Width: int(viaStrW) + 20, Height: fontSize + 10}
			objCoords[1000+(i+1)] = viabRS

			ggCtx.SetHexColor("#fff")
			ggCtx.DrawString(viaStr, float64(currentX)+10, float64(currentY)+fontSize+30)

			// duration
			durStr := "duration: " + instr["duration"]
			ggCtx.SetHexColor("#444")
			ggCtx.DrawString(durStr, float64(currentX), float64(currentY)+fontSize+30+15+fontSize)

			// view audio asset optional
			if instr["audio_optional"] != "" {
				vaaStr := "View Audio Asset #" + strconv.Itoa(i+1)
				vaaStrW, _ := ggCtx.MeasureString(vaaStr)
				ggCtx.SetHexColor("#74A299")

				ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30+65, vaaStrW+20, fontSize+10, 10)
				ggCtx.Fill()
				vaabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30 + 65,
					Width: int(vaaStrW) + 20, Height: fontSize + 10}
				objCoords[2000+(i+1)] = vaabRS

				ggCtx.SetHexColor("#fff")
				ggCtx.DrawString(vaaStr, float64(currentX)+10, float64(currentY)+fontSize+30+65)
			}

		} else if instr["kind"] == "video" {
			// view video asset
			viaStr := "View Video Asset #" + strconv.Itoa(i+1)
			viaStrW, _ := ggCtx.MeasureString(viaStr)

			ggCtx.SetHexColor("#5F699F")
			ggCtx.DrawRoundedRectangle(float64(currentX), float64(currentY)+30, viaStrW+20, fontSize+10, 10)
			ggCtx.Fill()
			vvabRS := g143.RectSpecs{OriginX: currentX, OriginY: currentY + 30,
				Width: int(viaStrW) + 20, Height: fontSize + 10}
			objCoords[3000+(i+1)] = vvabRS

			ggCtx.SetHexColor("#fff")
			ggCtx.DrawString(viaStr, float64(currentX)+10, float64(currentY)+fontSize+30)

			// duration
			durStr := "begin: " + instr["begin"] + " | end: " + instr["end"]
			ggCtx.SetHexColor("#444")
			ggCtx.DrawString(durStr, float64(currentX), float64(currentY)+fontSize+30+15+fontSize)

		}

		newX := currentX + int(viaStrW) + 20
		if newX > (wWidth - int(viaStrW)) {
			currentY += 160
			currentX = 20
		} else {
			currentX += int(viaStrW) + 20 + 10
		}
	}
	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.RectSpecs
	var widgetCode int

	for code, RS := range objCoords {
		if g143.InRectSpecs(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {
	case AddImgBtn:
		// tmpFrame = currentWindowFrame
		drawViewAddImage(window, currentWindowFrame)
		window.SetMouseButtonCallback(viewAddImageMouseCallback)
		window.SetKeyCallback(vaikeyCallback)

	case AddVidBtn:
		drawViewAddVideo(window, currentWindowFrame)
		window.SetMouseButtonCallback(viewAddVideoMouseCallback)
		window.SetKeyCallback(vavkeyCallback)

	case OpenWDBtn:
		rootPath, _ := GetRootPath()
		externalLaunch(rootPath)

	case RenderBtn:
		drawRenderView(window, currentWindowFrame)
		window.SetMouseButtonCallback(nil)
		window.SetKeyCallback(nil)
		inChannel <- true
	}

	// for generated buttons
	if widgetCode > 1000 && widgetCode < 2000 {
		instrNum := widgetCode - 1000
		externalLaunch(instructions[instrNum-1]["image"])
	} else if widgetCode > 2000 && widgetCode < 3000 {
		instrNum := widgetCode - 2000
		externalLaunch(instructions[instrNum-1]["audio_optional"])
	} else if widgetCode > 3000 {
		instrNum := widgetCode - 3000
		externalLaunch(instructions[instrNum-1]["video"])
	}

}
